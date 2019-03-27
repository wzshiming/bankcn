package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wzshiming/bankcn/service/bankcn"
	"github.com/wzshiming/configer"
	ffmt "gopkg.in/ffmt.v1"
)

type conf struct {
	Port int `configer:"8080,env"`
}

func main() {

	var conf conf
	err := configer.Load(&conf)
	if err != nil {
		ffmt.Mark(err)
		return
	}

	ffmt.Puts(conf)
	mux0 := mux.NewRouter()

	RouteOpenAPI(mux0)

	{
		service, err := bankcn.NewBankcnService()
		if err != nil {
			return
		}
		RouteBankcnService(mux0, service)
	}

	mux := handlers.RecoveryHandler()(mux0)
	mux = handlers.CombinedLoggingHandler(os.Stdout, mux)
	p := fmt.Sprintf(":%d", conf.Port)

	fmt.Printf("Open http://127.0.0.1:%d/swagger/ with your browser.\n", conf.Port)
	err = http.ListenAndServe(p, mux)
	if err != nil {
		ffmt.Mark(err)
	}
	return
}
