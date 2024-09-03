package utilidades

import (
	"crypto/tls"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}
