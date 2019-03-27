// Code generated; DO NOT EDIT.
// file ./cmd/bankcn/router_gen.go

package main

import (
	mux "github.com/gorilla/mux"
	ui "github.com/wzshiming/openapi/ui"
	redoc "github.com/wzshiming/openapi/ui/redoc"
	swaggereditor "github.com/wzshiming/openapi/ui/swaggereditor"
	swaggerui "github.com/wzshiming/openapi/ui/swaggerui"
	http "net/http"
)

// Router is all routing for package
// generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	router = RouteOpenAPI(router)

	return router
}

var OpenAPI4YAML = []byte(`openapi: 3.0.1
info:
  title: OpenAPI Demo
  description: Automatically generated
  contact:
    name: wzshiming
    url: https://github.com/wzshiming/gen
  version: 0.0.1
servers:
- url: /
- url: '{scheme}{host}{port}{path}'
  variables:
    host:
      enum:
      - localhost
      default: localhost
    path:
      enum:
      - /
      default: /
    port:
      enum:
      - ""
      default: ""
    scheme:
      enum:
      - http://
      - https://
      default: http://
components: {}
`)
var OpenAPI4JSON = []byte(`{"openapi":"3.0.1","info":{"title":"OpenAPI Demo","description":"Automatically generated","contact":{"name":"wzshiming","url":"https://github.com/wzshiming/gen"},"version":"0.0.1"},"servers":[{"url":"/"},{"url":"{scheme}{host}{port}{path}","variables":{"host":{"enum":["localhost"],"default":"localhost"},"path":{"enum":["/"],"default":"/"},"port":{"enum":[""],"default":""},"scheme":{"enum":["http://","https://"],"default":"http://"}}}],"components":{}}`)

// RouteOpenAPI
func RouteOpenAPI(router *mux.Router) *mux.Router {
	openapi := map[string][]byte{
		"openapi.json": OpenAPI4JSON,
		"openapi.yml":  OpenAPI4YAML,
		"openapi.yaml": OpenAPI4YAML,
	}
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggereditor/").Handler(http.StripPrefix("/swaggereditor", ui.HandleWithFiles(openapi, swaggereditor.Asset)))
	router.PathPrefix("/redoc/").Handler(http.StripPrefix("/redoc", ui.HandleWithFiles(openapi, redoc.Asset)))
	return router
}
