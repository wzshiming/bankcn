// Code generated; DO NOT EDIT.
// file ./cmd/bankcn/router_gen.go

package main

import (
	json "encoding/json"
	mux "github.com/gorilla/mux"
	githubComWzshimingBankcn "github.com/wzshiming/bankcn"
	githubComWzshimingBankcnServiceBankcn "github.com/wzshiming/bankcn/service/bankcn"
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

	// BankcnService Define the method scope
	var _bankcnService githubComWzshimingBankcnServiceBankcn.BankcnService
	RouteBankcnService(router, &_bankcnService)

	router = RouteOpenAPI(router)

	return router
}

// RouteBankcnService is routing for BankcnService
func RouteBankcnService(router *mux.Router, _bankcnService *githubComWzshimingBankcnServiceBankcn.BankcnService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/bankcn").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing GET /bankcn/banks
	var __operationGetBankcnBanks http.Handler
	__operationGetBankcnBanks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBanks(_bankcnService, w, r)
	})
	subrouter.Methods("GET").Path("/banks").Handler(__operationGetBankcnBanks)

	// Registered routing GET /bankcn/{bank_id}
	var __operationGetBankcnBankID http.Handler
	__operationGetBankcnBankID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBankID(_bankcnService, w, r)
	})
	subrouter.Methods("GET").Path("/{bank_id}").Handler(__operationGetBankcnBankID)

	// Registered routing GET /bankcn/banks/{bank}.jpg
	var __operationGetBankcnBanksBankJpg http.Handler
	__operationGetBankcnBanksBankJpg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBanksBankJpg(_bankcnService, w, r)
	})
	subrouter.Methods("GET").Path("/banks/{bank}.jpg").Handler(__operationGetBankcnBanksBankJpg)

	// Registered routing GET /bankcn/{bank}/{area_id}
	var __operationGetBankcnBankAreaID http.Handler
	__operationGetBankcnBankAreaID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBankAreaID(_bankcnService, w, r)
	})
	subrouter.Methods("GET").Path("/{bank}/{area_id}").Handler(__operationGetBankcnBankAreaID)

	return router
}

// _requestPathAreaID Parsing the path for of area_id
func _requestPathAreaID(w http.ResponseWriter, r *http.Request) (_areaID string, err error) {

	var _raw_areaID = mux.Vars(r)["area_id"]
	_areaID = string(_raw_areaID)

	return
}

// _requestPathBank Parsing the path for of bank
func _requestPathBank(w http.ResponseWriter, r *http.Request) (_bank string, err error) {

	var _raw_bank = mux.Vars(r)["bank"]
	_bank = string(_raw_bank)

	return
}

// _requestPathBankID Parsing the path for of bank_id
func _requestPathBankID(w http.ResponseWriter, r *http.Request) (_bankID string, err error) {

	var _raw_bankID = mux.Vars(r)["bank_id"]
	_bankID = string(_raw_bankID)

	return
}

// _operationGetBankcnBankID Is the route of Verify
func _operationGetBankcnBankID(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _bankID string
	var _valid *githubComWzshimingBankcn.Valid

	// Parsing bank_id.
	_bankID, err = _requestPathBankID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Verify.
	_valid, err = s.Verify(_bankID)

	// Response code 200 OK for valid.
	if _valid != nil {
		var __valid []byte
		__valid, err = json.Marshal(_valid)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__valid)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __valid []byte
	__valid, err = json.Marshal(_valid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__valid)

	return
}

// _operationGetBankcnBankAreaID Is the route of Get
func _operationGetBankcnBankAreaID(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _bank string
	var _areaID string
	var _banks []*githubComWzshimingBankcn.Bank

	// Parsing bank.
	_bank, err = _requestPathBank(w, r)
	if err != nil {
		return
	}

	// Parsing area_id.
	_areaID, err = _requestPathAreaID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Get.
	_banks, err = s.Get(_bank, _areaID)

	// Response code 200 OK for banks.
	if _banks != nil {
		var __banks []byte
		__banks, err = json.Marshal(_banks)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__banks)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __banks []byte
	__banks, err = json.Marshal(_banks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__banks)

	return
}

// _operationGetBankcnBanksBankJpg Is the route of BanksJPG
func _operationGetBankcnBanksBankJpg(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _bank string
	var _file []byte

	// Parsing bank.
	_bank, err = _requestPathBank(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.BanksJPG.
	_file, err = s.BanksJPG(_bank)

	// Response code 200 OK for file.
	if _file != nil {
		var __file []byte
		__file = _file

		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(200)
		w.Write(__file)
		return
	}

	// Response code 404 Not Found for err.
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	var __file []byte
	__file = _file

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(200)
	w.Write(__file)

	return
}

// _operationGetBankcnBanks Is the route of Banks
func _operationGetBankcnBanks(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _banks_1 map[string]string

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Banks.
	_banks_1, err = s.Banks()

	// Response code 200 OK for banks.
	if _banks_1 != nil {
		var __banks_1 []byte
		__banks_1, err = json.Marshal(_banks_1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__banks_1)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __banks_1 []byte
	__banks_1, err = json.Marshal(_banks_1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__banks_1)

	return
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
paths:
  /bankcn/banks:
    get:
      tags:
      - BankcnService
      summary: 'Banks #route:"GET /banks"# 获取全部的银行'
      description: 'Banks #route:"GET /banks"# 获取全部的银行'
      responses:
        "200":
          description: Response code is 200
          content:
            application/json:
              schema:
                type: object
                additionalProperties:
                  type: string
        "400":
          description: Response code is 400
          content:
            text/plain:
              schema:
                type: string
                format: error
  /bankcn/banks/{bank}.jpg:
    get:
      tags:
      - BankcnService
      summary: 'Banks #route:"GET /banks/{bank}.jpg"# 获取某地区的银行图标'
      description: 'Banks #route:"GET /banks/{bank}.jpg"# 获取某地区的银行图标'
      parameters:
      - $ref: '#/components/parameters/bank'
      responses:
        "200":
          description: '#content:"image/jpeg"#'
          content:
            image/jpeg: {}
        "404":
          description: '#code:"404"#'
          content:
            text/plain:
              schema:
                type: string
                format: error
  /bankcn/{bank_id}:
    get:
      tags:
      - BankcnService
      summary: 'Verify #route:"GET /{bank_id}"# 获取银行卡属性'
      description: 'Verify #route:"GET /{bank_id}"# 获取银行卡属性'
      parameters:
      - $ref: '#/components/parameters/bank_id'
      responses:
        "200":
          description: Response code is 200
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Valid'
        "400":
          description: Response code is 400
          content:
            text/plain:
              schema:
                type: string
                format: error
  /bankcn/{bank}/{area_id}:
    get:
      tags:
      - BankcnService
      summary: 'Get #route:"GET /{bank}/{area_id}"# 获取某地区的银行数据'
      description: 'Get #route:"GET /{bank}/{area_id}"# 获取某地区的银行数据'
      parameters:
      - $ref: '#/components/parameters/bank'
      - $ref: '#/components/parameters/area_id'
      responses:
        "200":
          description: Response code is 200
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Bank'
        "400":
          description: Response code is 400
          content:
            text/plain:
              schema:
                type: string
                format: error
components:
  schemas:
    Bank:
      type: object
      properties:
        address:
          type: string
          description: 联系地址
        area_id:
          type: string
          description: 所在区域代号
        bank:
          type: string
          description: 银行标识符
        bank_union_id:
          type: string
          description: 银行联行号
        name:
          type: string
          description: 银行名字
        phone:
          type: string
          description: 联系电话
    Valid:
      type: object
      properties:
        bank:
          type: string
          description: 银行卡所属银行
        bank_id_type:
          type: string
          description: 银行卡 类型
  responses:
    banks:
      description: Response code is 200
      content:
        application/json:
          schema:
            type: array
            items:
              $ref: '#/components/schemas/Bank'
    banks.1:
      description: Response code is 200
      content:
        application/json:
          schema:
            type: object
            additionalProperties:
              type: string
    err:
      description: Response code is 400
      content:
        text/plain:
          schema:
            type: string
            format: error
    err.1:
      description: '#code:"404"#'
      content:
        text/plain:
          schema:
            type: string
            format: error
    file:
      description: '#content:"image/jpeg"#'
      content:
        image/jpeg: {}
    valid:
      description: Response code is 200
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Valid'
  parameters:
    area_id:
      name: area_id
      in: path
      description: '#name:"area_id"#'
      required: true
      schema:
        type: string
    bank:
      name: bank
      in: path
      description: '#name:"bank"#'
      required: true
      schema:
        type: string
    bank_id:
      name: bank_id
      in: path
      description: '#name:"bank_id"#'
      required: true
      schema:
        type: string
tags:
- name: BankcnService
  description: 'BankcnService #path:"/bankcn/"#'
`)
var OpenAPI4JSON = []byte(`{"openapi":"3.0.1","info":{"title":"OpenAPI Demo","description":"Automatically generated","contact":{"name":"wzshiming","url":"https://github.com/wzshiming/gen"},"version":"0.0.1"},"servers":[{"url":"/"},{"url":"{scheme}{host}{port}{path}","variables":{"host":{"enum":["localhost"],"default":"localhost"},"path":{"enum":["/"],"default":"/"},"port":{"enum":[""],"default":""},"scheme":{"enum":["http://","https://"],"default":"http://"}}}],"paths":{"/bankcn/banks":{"get":{"tags":["BankcnService"],"summary":"Banks #route:\"GET /banks\"# 获取全部的银行","description":"Banks #route:\"GET /banks\"# 获取全部的银行","responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"object","additionalProperties":{"type":"string"}}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/banks/{bank}.jpg":{"get":{"tags":["BankcnService"],"summary":"Banks #route:\"GET /banks/{bank}.jpg\"# 获取某地区的银行图标","description":"Banks #route:\"GET /banks/{bank}.jpg\"# 获取某地区的银行图标","parameters":[{"$ref":"#/components/parameters/bank"}],"responses":{"200":{"description":"#content:\"image/jpeg\"#","content":{"image/jpeg":{}}},"404":{"description":"#code:\"404\"#","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/{bank_id}":{"get":{"tags":["BankcnService"],"summary":"Verify #route:\"GET /{bank_id}\"# 获取银行卡属性","description":"Verify #route:\"GET /{bank_id}\"# 获取银行卡属性","parameters":[{"$ref":"#/components/parameters/bank_id"}],"responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"$ref":"#/components/schemas/Valid"}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/{bank}/{area_id}":{"get":{"tags":["BankcnService"],"summary":"Get #route:\"GET /{bank}/{area_id}\"# 获取某地区的银行数据","description":"Get #route:\"GET /{bank}/{area_id}\"# 获取某地区的银行数据","parameters":[{"$ref":"#/components/parameters/bank"},{"$ref":"#/components/parameters/area_id"}],"responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Bank"}}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}}},"components":{"schemas":{"Bank":{"type":"object","properties":{"address":{"type":"string","description":"联系地址"},"area_id":{"type":"string","description":"所在区域代号"},"bank":{"type":"string","description":"银行标识符"},"bank_union_id":{"type":"string","description":"银行联行号"},"name":{"type":"string","description":"银行名字"},"phone":{"type":"string","description":"联系电话"}}},"Valid":{"type":"object","properties":{"bank":{"type":"string","description":"银行卡所属银行"},"bank_id_type":{"type":"string","description":"银行卡 类型"}}}},"responses":{"banks":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Bank"}}}}},"banks.1":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"object","additionalProperties":{"type":"string"}}}}},"err":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}},"err.1":{"description":"#code:\"404\"#","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}},"file":{"description":"#content:\"image/jpeg\"#","content":{"image/jpeg":{}}},"valid":{"description":"Response code is 200","content":{"application/json":{"schema":{"$ref":"#/components/schemas/Valid"}}}}},"parameters":{"area_id":{"name":"area_id","in":"path","description":"#name:\"area_id\"#","required":true,"schema":{"type":"string"}},"bank":{"name":"bank","in":"path","description":"#name:\"bank\"#","required":true,"schema":{"type":"string"}},"bank_id":{"name":"bank_id","in":"path","description":"#name:\"bank_id\"#","required":true,"schema":{"type":"string"}}}},"tags":[{"name":"BankcnService","description":"BankcnService #path:\"/bankcn/\"#"}]}`)

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
