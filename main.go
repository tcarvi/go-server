/*
 * Server for static websites
 */
package main

import (
	"log"
	"net/http"

	"github.com/tcarvi/go-server/handler"
)

func main() {
	var tHandler handler.THandler
	err := http.ListenAndServe(":8080", tHandler)
	// TODO HTTPS ACCESS: log.Fatal(srv.ListenAndServeTLS(certFile, keyFile string))
	if err != nil {
		log.Fatal(err)
	}

}
