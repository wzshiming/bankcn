package configer

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const tag = "configer"

func ProcessTags(config interface{}) error {
	return processTags(reflect.ValueOf(config))
}

// 处理任意反射映射到结构体
func processTags(configValue reflect.Value) error {
	configValue = reflect.Indirect(configValue)
	switch configValue.Kind() {
	case reflect.Ptr:
		return processTags(configValue)
	case reflect.Slice, reflect.Array:
		for i := 0; i < configValue.Len(); i++ {
			err := processTags(configValue.Index(i))
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, v := range configValue.MapKeys() {
			err := processTags(configValue.MapIndex(v))
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		return processTagsStruct(configValue)
	}
	return nil
}

// 处理结构体
// `configer:"-"`  		 不处理
// `configer:",env"`	 没有默认值 从env 里取值
// `configer:"test"` 	 默认值 test
// `configer:"test,env"` 默认值 test 如果env 有值得话覆盖掉原本的值
func processTagsStruct(configValue reflect.Value) error {
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			fieldStruct = configType.Field(i)
			field       = configValue.Field(i)
			name        = fieldStruct.Name
		)

		if !field.CanAddr() || !field.CanInterface() {
			continue
		}

		tagvalue := fieldStruct.Tag.Get(tag)
		if tagvalue == "-" {
			continue
		}

		ts := strings.Split(tagvalue, ",")
		if len(ts) > 1 {
			for _, v := range ts[1:] {
				switch v {
				case "env":
					envname := fieldStruct.Tag.Get("env")
					if envname == "" {
						envname = name
					}

					// 从 环境变量里取值
					if value := os.Getenv(envname); value != "" {
						if err := yamlUnmarshal([]byte(value), field.Addr().Interface()); err != nil {
							return err
						}
					}
				case "load":
					loadf := fieldStruct.Tag.Get("load")
					if loadf == "" {
						continue
					}
					val := configValue.FieldByName(loadf)

					for field.Kind() == reflect.Ptr {
						if field.IsNil() {
							field.Set(reflect.New(field.Type().Elem()))
						}
						field = field.Elem()
					}
					if err := Load(field.Addr().Interface(), fmt.Sprint(val.Interface())); err != nil {
						return err
					}
				}
			}
		}

		if value := ts[0]; value != "" {
			// 判断是否为空
			if isBlank := reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()); isBlank {
				if err := yamlUnmarshal([]byte(value), field.Addr().Interface()); err != nil {
					return err
				}
			}
		}

		if err := processTags(field); err != nil {
			return err
		}
	}
	return nil
}
