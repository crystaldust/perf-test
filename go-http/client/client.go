package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var targetUrl string
var jsonBytes []byte
var imageBytes []byte

func main() {
	fmt.Printf("http_proxy: %s\n", os.Getenv("http_proxy"))

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

	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	listenAddr := ":" + port
	log.Fatal(http.ListenAndServe(listenAddr, mux))
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

	client := GetHttpClient()
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

var tp http.RoundTripper = &http.Transport{
	Proxy:               http.ProxyFromEnvironment,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	DialContext: (&net.Dialer{
		KeepAlive: 60 * time.Second,
		Timeout:   60 * time.Second,
	}).DialContext,
}

func GetHttpClient() *http.Client {
	return &http.Client{
		Transport: tp,
	}
}
