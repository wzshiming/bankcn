// Code generated; DO NOT EDIT.
// file ./cmd/bankcn/router_gen.go

package main

import (
	json "encoding/json"
	http "net/http"

	mux "github.com/gorilla/mux"
	githubComWzshimingBankcn "github.com/wzshiming/bankcn"
	githubComWzshimingBankcnServiceBankcn "github.com/wzshiming/bankcn/service/bankcn"
	ui "github.com/wzshiming/openapi/ui"
	redoc "github.com/wzshiming/openapi/ui/redoc"
	swaggereditor "github.com/wzshiming/openapi/ui/swaggereditor"
	swaggerui "github.com/wzshiming/openapi/ui/swaggerui"
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

	_routeBankcn := router.PathPrefix("/bankcn").Subrouter()
	if len(fs) != 0 {
		_routeBankcn.Use(fs...)
	}

	// Registered routing GET /bankcn/banks
	var __operationGetBankcnBanks http.Handler
	__operationGetBankcnBanks = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBanks(_bankcnService, w, r)
	})
	_routeBankcn.Methods("GET").Path("/banks").Handler(__operationGetBankcnBanks)

	// Registered routing GET /bankcn/banks/{bank}.jpg
	var __operationGetBankcnBanksBankJpg http.Handler
	__operationGetBankcnBanksBankJpg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBanksBankJpg(_bankcnService, w, r)
	})
	_routeBankcn.Methods("GET").Path("/banks/{bank}.jpg").Handler(__operationGetBankcnBanksBankJpg)

	// Registered routing GET /bankcn/bank_id/{bank_id}
	var __operationGetBankcnBankIDBankID http.Handler
	__operationGetBankcnBankIDBankID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBankIDBankID(_bankcnService, w, r)
	})
	_routeBankcn.Methods("GET").Path("/bank_id/{bank_id}").Handler(__operationGetBankcnBankIDBankID)

	// Registered routing GET /bankcn/banks/{bank}/{area_id}
	var __operationGetBankcnBanksBankAreaID http.Handler
	__operationGetBankcnBanksBankAreaID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetBankcnBanksBankAreaID(_bankcnService, w, r)
	})
	_routeBankcn.Methods("GET").Path("/banks/{bank}/{area_id}").Handler(__operationGetBankcnBanksBankAreaID)

	return router
}

// _requestPathAreaID Parsing the path for of area_id
func _requestPathAreaID(w http.ResponseWriter, r *http.Request) (_areaID string, err error) {

	var _rawAreaID = mux.Vars(r)["area_id"]
	_areaID = string(_rawAreaID)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _requestPathBankID Parsing the path for of bank_id
func _requestPathBankID(w http.ResponseWriter, r *http.Request) (_bankID string, err error) {

	var _rawBankID = mux.Vars(r)["bank_id"]
	_bankID = string(_rawBankID)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _requestPathBank Parsing the path for of bank
func _requestPathBank(w http.ResponseWriter, r *http.Request) (_bank string, err error) {

	var _rawBank = mux.Vars(r)["bank"]
	_bank = string(_rawBank)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _requestQueryFilter Parsing the query for of filter
func _requestQueryFilter(w http.ResponseWriter, r *http.Request) (_filter string, err error) {

	var _rawFilter = r.URL.Query()["filter"]

	if len(_rawFilter) == 0 {
		return
	}
	_filter = string(_rawFilter[0])
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _operationGetBankcnBankIDBankID Is the route of Verify
func _operationGetBankcnBankIDBankID(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {
	var _bankID string
	var _valid *githubComWzshimingBankcn.Valid
	var _err error

	// Parsing bank_id.
	_bankID, _err = _requestPathBankID(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Verify.
	_valid, _err = s.Verify(_bankID)

	// Response code 200 OK for valid.
	if _valid != nil {
		var __valid []byte
		__valid, _err = json.Marshal(_valid)
		if _err != nil {
			http.Error(w, _err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__valid)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	var __valid []byte
	__valid, _err = json.Marshal(_valid)
	if _err != nil {
		http.Error(w, _err.Error(), 500)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__valid)

	return
}

// _operationGetBankcnBanksBankAreaID Is the route of Get
func _operationGetBankcnBanksBankAreaID(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {
	var _bank string
	var _areaID string
	var _filter string
	var _banks []*githubComWzshimingBankcn.Bank
	var _err error

	// Parsing bank.
	_bank, _err = _requestPathBank(w, r)
	if _err != nil {
		return
	}

	// Parsing area_id.
	_areaID, _err = _requestPathAreaID(w, r)
	if _err != nil {
		return
	}

	// Parsing filter.
	_filter, _err = _requestQueryFilter(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Get.
	_banks, _err = s.Get(_bank, _areaID, _filter)

	// Response code 200 OK for banks.
	if _banks != nil {
		var __banks []byte
		__banks, _err = json.Marshal(_banks)
		if _err != nil {
			http.Error(w, _err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__banks)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	var __banks []byte
	__banks, _err = json.Marshal(_banks)
	if _err != nil {
		http.Error(w, _err.Error(), 500)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__banks)

	return
}

// _operationGetBankcnBanksBankJpg Is the route of BanksJPG
func _operationGetBankcnBanksBankJpg(s *githubComWzshimingBankcnServiceBankcn.BankcnService, w http.ResponseWriter, r *http.Request) {
	var _bank string
	var _file []byte
	var _err_1 error

	// Parsing bank.
	_bank, _err_1 = _requestPathBank(w, r)
	if _err_1 != nil {
		return
	}

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.BanksJPG.
	_file, _err_1 = s.BanksJPG(_bank)

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
	if _err_1 != nil {
		http.Error(w, _err_1.Error(), 404)
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
	var _bankItems []*githubComWzshimingBankcn.BankItem
	var _err error

	// Call github.com/wzshiming/bankcn/service/bankcn BankcnService.Banks.
	_bankItems, _err = s.Banks()

	// Response code 200 OK for bankItems.
	if _bankItems != nil {
		var __bankItems []byte
		__bankItems, _err = json.Marshal(_bankItems)
		if _err != nil {
			http.Error(w, _err.Error(), 500)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__bankItems)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	var __bankItems []byte
	__bankItems, _err = json.Marshal(_bankItems)
	if _err != nil {
		http.Error(w, _err.Error(), 500)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__bankItems)

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
  /bankcn/bank_id/{bank_id}:
    get:
      tags:
      - BankcnService
      summary: Verify  获取银行卡号的 所属银行 以及 类型
      description: |-
        Verify  获取银行卡号的 所属银行 以及 类型
        #route:"GET /bank_id/{bank_id}"#
      parameters:
      - $ref: '#/components/parameters/bank_id_path'
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
  /bankcn/banks:
    get:
      tags:
      - BankcnService
      summary: Banks  获取全部的银行类型的列表
      description: |-
        Banks  获取全部的银行类型的列表
        #route:"GET /banks"#
      responses:
        "200":
          description: Response code is 200
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/BankItem'
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
      summary: Banks  获取某银行图标
      description: |-
        Banks  获取某银行图标
        #route:"GET /banks/{bank}.jpg"#
      parameters:
      - $ref: '#/components/parameters/bank_path'
      responses:
        "200":
          description: |2-

            #content:"image/jpeg"#
          content:
            image/jpeg:
              schema:
                type: string
                format: binary
        "404":
          description: |2-

            #code:"404"#
          content:
            text/plain:
              schema:
                type: string
                format: error
  /bankcn/banks/{bank}/{area_id}:
    get:
      tags:
      - BankcnService
      summary: Get  获取某地区的银行数据
      description: |-
        Get  获取某地区的银行数据
        #route:"GET /banks/{bank}/{area_id}"#
      parameters:
      - $ref: '#/components/parameters/bank_path'
      - $ref: '#/components/parameters/area_id_path'
      - $ref: '#/components/parameters/filter_query'
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
    BankItem:
      required:
      - bank
      - bank_name
      type: object
      properties:
        bank:
          type: string
          description: 银行代号
        bank_name:
          type: string
          description: 银行名
    Valid:
      required:
      - bank_id_type
      - bank
      type: object
      properties:
        bank:
          type: string
          description: 银行卡所属银行
        bank_id_type:
          type: string
          description: 银行卡 类型
  responses:
    bankItems_body:
      description: Response code is 200
      content:
        application/json:
          schema:
            type: array
            items:
              $ref: '#/components/schemas/BankItem'
    banks_body:
      description: Response code is 200
      content:
        application/json:
          schema:
            type: array
            items:
              $ref: '#/components/schemas/Bank'
    err_body:
      description: Response code is 400
      content:
        text/plain:
          schema:
            type: string
            format: error
    err_body.1:
      description: |2-

        #code:"404"#
      content:
        text/plain:
          schema:
            type: string
            format: error
    file_body:
      description: |2-

        #content:"image/jpeg"#
      content:
        image/jpeg:
          schema:
            type: string
            format: binary
    valid_body:
      description: Response code is 200
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Valid'
  parameters:
    area_id_path:
      name: area_id
      in: path
      description: '#name:"area_id"#'
      required: true
      schema:
        type: string
    bank_id_path:
      name: bank_id
      in: path
      description: '#name:"bank_id"#'
      required: true
      schema:
        type: string
    bank_path:
      name: bank
      in: path
      description: '#name:"bank"#'
      required: true
      schema:
        type: string
    filter_query:
      name: filter
      in: query
      schema:
        type: string
tags:
- name: BankcnService
  description: "BankcnService \n#path:\"/bankcn/\"#"
`)
var OpenAPI4JSON = []byte(`{"openapi":"3.0.1","info":{"title":"OpenAPI Demo","description":"Automatically generated","contact":{"name":"wzshiming","url":"https://github.com/wzshiming/gen"},"version":"0.0.1"},"servers":[{"url":"/"},{"url":"{scheme}{host}{port}{path}","variables":{"host":{"enum":["localhost"],"default":"localhost"},"path":{"enum":["/"],"default":"/"},"port":{"enum":[""],"default":""},"scheme":{"enum":["http://","https://"],"default":"http://"}}}],"paths":{"/bankcn/bank_id/{bank_id}":{"get":{"tags":["BankcnService"],"summary":"Verify  获取银行卡号的 所属银行 以及 类型","description":"Verify  获取银行卡号的 所属银行 以及 类型\n#route:\"GET /bank_id/{bank_id}\"#","parameters":[{"$ref":"#/components/parameters/bank_id_path"}],"responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"$ref":"#/components/schemas/Valid"}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/banks":{"get":{"tags":["BankcnService"],"summary":"Banks  获取全部的银行类型的列表","description":"Banks  获取全部的银行类型的列表\n#route:\"GET /banks\"#","responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/BankItem"}}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/banks/{bank}.jpg":{"get":{"tags":["BankcnService"],"summary":"Banks  获取某银行图标","description":"Banks  获取某银行图标\n#route:\"GET /banks/{bank}.jpg\"#","parameters":[{"$ref":"#/components/parameters/bank_path"}],"responses":{"200":{"description":"\n#content:\"image/jpeg\"#","content":{"image/jpeg":{"schema":{"type":"string","format":"binary"}}}},"404":{"description":"\n#code:\"404\"#","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}},"/bankcn/banks/{bank}/{area_id}":{"get":{"tags":["BankcnService"],"summary":"Get  获取某地区的银行数据","description":"Get  获取某地区的银行数据\n#route:\"GET /banks/{bank}/{area_id}\"#","parameters":[{"$ref":"#/components/parameters/bank_path"},{"$ref":"#/components/parameters/area_id_path"},{"$ref":"#/components/parameters/filter_query"}],"responses":{"200":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Bank"}}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}}},"components":{"schemas":{"Bank":{"type":"object","properties":{"address":{"type":"string","description":"联系地址"},"area_id":{"type":"string","description":"所在区域代号"},"bank":{"type":"string","description":"银行标识符"},"bank_union_id":{"type":"string","description":"银行联行号"},"name":{"type":"string","description":"银行名字"},"phone":{"type":"string","description":"联系电话"}}},"BankItem":{"required":["bank","bank_name"],"type":"object","properties":{"bank":{"type":"string","description":"银行代号"},"bank_name":{"type":"string","description":"银行名"}}},"Valid":{"required":["bank_id_type","bank"],"type":"object","properties":{"bank":{"type":"string","description":"银行卡所属银行"},"bank_id_type":{"type":"string","description":"银行卡 类型"}}}},"responses":{"bankItems_body":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/BankItem"}}}}},"banks_body":{"description":"Response code is 200","content":{"application/json":{"schema":{"type":"array","items":{"$ref":"#/components/schemas/Bank"}}}}},"err_body":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}},"err_body.1":{"description":"\n#code:\"404\"#","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}},"file_body":{"description":"\n#content:\"image/jpeg\"#","content":{"image/jpeg":{"schema":{"type":"string","format":"binary"}}}},"valid_body":{"description":"Response code is 200","content":{"application/json":{"schema":{"$ref":"#/components/schemas/Valid"}}}}},"parameters":{"area_id_path":{"name":"area_id","in":"path","description":"#name:\"area_id\"#","required":true,"schema":{"type":"string"}},"bank_id_path":{"name":"bank_id","in":"path","description":"#name:\"bank_id\"#","required":true,"schema":{"type":"string"}},"bank_path":{"name":"bank","in":"path","description":"#name:\"bank\"#","required":true,"schema":{"type":"string"}},"filter_query":{"name":"filter","in":"query","schema":{"type":"string"}}}},"tags":[{"name":"BankcnService","description":"BankcnService \n#path:\"/bankcn/\"#"}]}`)

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
