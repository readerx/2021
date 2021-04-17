package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://192.168.1.25:30083/job/test/", nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.SetBasicAuth("admin", "admin")

	cert, err := tls.LoadX509KeyPair("e:/data/client.cert", "e:/data/client.key")
	if err != nil {
		log.Fatal(err)
	}

	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	repo, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	result, err := httputil.DumpResponse(repo, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("resp: [%s]\n", string(result))
}
