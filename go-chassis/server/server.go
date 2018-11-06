package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/core/server"
	rf "github.com/go-chassis/go-chassis/server/restful"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/

type RestFulHello struct {
}

func (r *RestFulHello) BodySize(b *rf.Context) {
	req := b.ReadRequest()
	n, err := io.Copy(ioutil.Discard, req.Body)
	if err != nil {
		b.WriteHeader(http.StatusInternalServerError)
		b.Write([]byte(err.Error()))
		return
	}
	b.WriteHeader(http.StatusOK)
	b.Write([]byte(fmt.Sprintf("size:%d", n)))
}

func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodPost, Path: "/bodysize", ResourceFuncName: "BodySize"},
	}
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//
	// })
	chassis.RegisterSchema("rest", &RestFulHello{}, server.WithSchemaID("RestHelloService"))
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
