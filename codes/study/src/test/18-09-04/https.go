package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cert := "E:\\data\\https\\server.cert"
	key := "E:\\data\\https\\server.key"
	ca := "E:\\data\\https\\ca-middle.crt"

	http.DefaultServeMux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("client cert %+v\n", request.TLS.PeerCertificates)

		for _, cert := range request.TLS.PeerCertificates {
			log.Printf("%+v", cert.DNSNames)

		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("test over"))
	})

	caData, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
		TLSConfig: &tls.Config{
			RootCAs:    pool,
			ClientAuth: tls.RequestClientCert,
		},
	}

	log.Fatal(server.ListenAndServeTLS(cert, key))
}
