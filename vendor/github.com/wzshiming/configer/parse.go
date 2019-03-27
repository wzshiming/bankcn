package configer

import (
	"errors"
	"path/filepath"
	"strings"
)

var (
	parseMap = map[string]func(data []byte, v interface{}) error{
		"json": jsonUnmarshal,
		"xml":  xmlUnmarshal,
		"yaml": yamlUnmarshal,
		"yml":  yamlUnmarshal,
		"toml": tomlUnmarshal,
		"ini":  iniUnmarshal,
		"hcl":  hclUnmarshal,
	}
	parseList = []func(data []byte, v interface{}) error{
		jsonUnmarshal,
		hclUnmarshal,
		xmlUnmarshal,
		tomlUnmarshal,
		yamlUnmarshal,
		iniUnmarshal,
	}
)

func Parse(data []byte, cfgpath string, val interface{}) (err error) {
	ext := strings.TrimPrefix(filepath.Ext(cfgpath), ".")
	pf := parseMap[ext]
	if pf != nil {
		return pf(data, val)
	}
	for _, v := range parseList {
		err = v(data, val)
		if err != nil {
			continue
		}
		return
	}
	return errors.New("Failed to decode config")
}
