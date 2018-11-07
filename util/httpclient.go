package util

import (
	"net"
	"net/http"
	"time"
)

var tp http.RoundTripper = &http.Transport{
	MaxIdleConns:        50,
	MaxIdleConnsPerHost: 50,
	DialContext: (&net.Dialer{
		KeepAlive: 60 * time.Second,
		Timeout:   60 * time.Second,
	}).DialContext,
}

// GetChassisHttpClient create the http client with the same
// options in go-chassis.
func GetChassisHttpClient() *http.Client {
	return &http.Client{
		Transport: tp,
	}
}
