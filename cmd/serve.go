package cmd

import (
	"go-url-checker/rest"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", rest.Verify)

	http.ListenAndServe(":7080", mux)

}
