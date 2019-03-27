package configer

import (
	"encoding/json"
	"encoding/xml"

	toml "github.com/BurntSushi/toml"
	hcl "github.com/hashicorp/hcl"
	ini "gopkg.in/ini.v1"
	yaml "gopkg.in/yaml.v2"
)

var (
	jsonUnmarshal = json.Unmarshal
	xmlUnmarshal  = xml.Unmarshal
	yamlUnmarshal = yaml.Unmarshal
	tomlUnmarshal = toml.Unmarshal
	hclUnmarshal  = hcl.Unmarshal
	iniUnmarshal  = func(d []byte, v interface{}) error {
		f, err := ini.InsensitiveLoad(d)
		if err != nil {
			return err
		}
		return f.MapTo(v)
	}
)
