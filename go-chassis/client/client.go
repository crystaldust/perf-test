package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
)

var jsonBytes []byte
var imageBytes []byte

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	var err error

	sampleFolder := os.Getenv("SAMPLE_FOLDER")
	if sampleFolder == "" {
		sampleFolder = "../../testdata"
	}
	sampleJsonPath := fmt.Sprintf("%s/sample.json", sampleFolder)
	sampleImagePath := fmt.Sprintf("%s/sample.png", sampleFolder)

	jsonBytes, err = ioutil.ReadFile(sampleJsonPath)
	if err != nil {
		panic("Failed to read sample json" + err.Error())
	}
	imageBytes, err = ioutil.ReadFile(sampleImagePath)
	if err != nil {
		panic("Failed to read image json")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		doJsonRequest(w)
	})
	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		doImageRequest(w)
	})

	http.ListenAndServe(":8002", mux)
}

func doJsonRequest(writer http.ResponseWriter) {
	req, err := rest.NewRequest("POST", "cse://RESTServer/bodysize", jsonBytes)
	if err != nil {
		lager.Logger.Error("new request failed.")
		return
	}

	resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
	if err != nil {
		lager.Logger.Error("do request failed.")
		return
	}
	defer resp.Body.Close()

	respBytes := httputil.ReadBody(resp)
	writer.Write(respBytes)
	// lager.Logger.Info("REST Server sayhello[GET]: " + string(respBytes))
}

func doImageRequest(writer http.ResponseWriter) {
	req, err := rest.NewRequest("POST", "cse://RESTServer/bodysize", imageBytes)
	if err != nil {
		lager.Logger.Error("new request failed.")
		return
	}

	resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
	if err != nil {
		lager.Logger.Error("do request failed.")
		return
	}
	defer resp.Body.Close()

	respBytes := httputil.ReadBody(resp)
	writer.Write(respBytes)
	// lager.Logger.Info("REST Server sayhello[GET]: " + string(respBytes))
}
