# configer

Configuration loader that support INI, XML, YAML, JSON, HCL, TOML, Shell Environment

- [English](https://github.com/wzshiming/configer/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/configer/blob/master/README_cn.md)

## Install

``` bash
go get -u -v github.com/wzshiming/configer
```

## Usage

[API Documentation](http://godoc.org/github.com/wzshiming/configer)

[Examples](https://github.com/wzshiming/configer/blob/master/examples/main.go)

Priority: env > default > conf

``` golang
package main

import (
	configer "gopkg.in/configer.v1"
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	examples1()
}

type BB struct {
	Hello   string `configer:"world"`            // Take the default value "world"
	Shell   string `configer:",env" env:"SHELL"` // Take the value of env
	EnvNone string `configer:",env" env:"NONE"`  // An empty env
}

type TT struct {
	LoadFilePath string `configer:"./examples1.json,env"`      // Loaded file path
	BB           BB     `configer:",load" load:"LoadFilePath"` // Load path field
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

With examples1.json:

``` json
{
    "Hello": "json",
    "Shell": "Priority default < env",
    "EnvNone": "env none"
}
```

## MIT License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/configer/blob/master/LICENSE) for the full license text.
