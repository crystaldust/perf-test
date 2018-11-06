package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.POST("/").To(handler))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func handler(req *restful.Request, resp *restful.Response) {
	n, err := io.Copy(ioutil.Discard, req.Request.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(fmt.Sprintf("size:%d", n)))
}
