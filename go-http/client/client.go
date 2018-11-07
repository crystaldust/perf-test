package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/crystaldust/perf-test/util"
)

var targetUrl string
var jsonBytes []byte
var imageBytes []byte

func main() {
	targetUrl = os.Getenv("TARGET")
	if targetUrl == "" {
		targetUrl = "http://localhost:9000"
	}
	sampleFolder := os.Getenv("SAMPLE_FOLDER")
	if sampleFolder == "" {
		sampleFolder = "../../testdata"
	}
	sampleJsonPath := fmt.Sprintf("%s/sample.json", sampleFolder)
	sampleImagePath := fmt.Sprintf("%s/sample.png", sampleFolder)

	jsonBytes, _ = ioutil.ReadFile(sampleJsonPath)
	imageBytes, _ = ioutil.ReadFile(sampleImagePath)

	mux := http.NewServeMux()
	mux.HandleFunc("/json", sendJson)
	mux.HandleFunc("/image", sendImage)

	log.Fatal(http.ListenAndServe(":8000", mux))
}

func sendJson(w http.ResponseWriter, r *http.Request) {
	reader := bytes.NewReader(jsonBytes)
	sendReq(w, r, reader)
}

func sendImage(w http.ResponseWriter, r *http.Request) {
	reader := bytes.NewReader(imageBytes)
	sendReq(w, r, reader)
}

func sendReq(w http.ResponseWriter, r *http.Request, reader io.Reader) {
	req, err := http.NewRequest(http.MethodPost, targetUrl, reader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := err.Error()
		fmt.Println(errStr)
		w.Write([]byte(errStr))
		return
	}

	client := util.GetChassisHttpClient()
	// client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := err.Error()
		fmt.Println(errStr)
		w.Write([]byte(errStr))
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := err.Error()
		fmt.Println(errStr)
		w.Write([]byte(errStr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)

}
