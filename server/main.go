package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":9999"
	caFile := "../cert/ca/ca.pem"
	certFile, keyFile := "../cert/server/server.pem", "../cert/server/server-key.pem"
	// certFile, keyFile := "../cert/peer/peer.pem", "../cert/peer/peer-key.pem"

	pem, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pem)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", time.Now().Unix())
	})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}
