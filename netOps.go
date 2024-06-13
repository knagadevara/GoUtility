package utl

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func CallApi(httpMethod, apiURL string) *http.Response {

	req, err := http.NewRequest(httpMethod, apiURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
