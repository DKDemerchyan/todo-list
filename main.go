package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	webDir := "./web"
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", mux)
	if err != nil {
		panic(err)
	}
}
