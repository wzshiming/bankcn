# configer

配置文件加载，支持INI，XML，YAML，JSON，HCL，TOML，Shell环境

- [English](https://github.com/wzshiming/configer/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/configer/blob/master/README_cn.md)

## 安装

``` bash
go get -u -v github.com/wzshiming/configer
```

## 使用

[API 文档](http://godoc.org/github.com/wzshiming/configer)

[示例](https://github.com/wzshiming/configer/blob/master/examples/main.go)

优先级: env > default > conf

``` golang
package main

import (
	configer "github.com/wzshiming/configer"
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	examples1()
}

type BB struct {
	Hello   string `configer:"world"`            // 取默认值 "world"
	Shell   string `configer:",env" env:"SHELL"` // 从 env 环境变量里取
	EnvNone string `configer:",env" env:"NONE"`  // 空的 env
}

type TT struct {
	LoadFilePath string `configer:"./examples1.json,env"`      // 加载文件的路径
	BB           BB     `configer:",load" load:"LoadFilePath"` // 加载路径字段
}

func examples1() {
	b := BB{}

	configer.Load(&b)
	ffmt.Puts(b)
	/*
		{
		 Hello:   "world"
		 Shell:   "/bin/bash"
		 EnvNone: ""
		}
	*/

	configer.Load(&b, "./examples1.json")
	ffmt.Puts(b)
	/*
		{
		 Hello:   "json"
		 Shell:   "/bin/bash"
		 EnvNone: "env none"
		}
	*/

	t := TT{}
	configer.Load(&t)
	ffmt.Puts(t)
	/*
		{
		 LoadFilePath: "./examples1.json"
		 BB:           {
		  Hello:   "json"
		  Shell:   "/bin/bash"
		  EnvNone: "env none"
		 }
		}
	*/
}

```

文件: examples1.json:

``` json
{
    "Hello": "json",
    "Shell": "Priority default < env",
    "EnvNone": "env none"
}
```

## MIT许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/configer/blob/master/LICENSE)。
