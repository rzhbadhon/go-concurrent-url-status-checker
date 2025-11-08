package cmd

import "net/http"

func Server() {
	mux := http.NewServeMux()

	http.ListenAndServe(":7080", mux)


}