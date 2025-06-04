package main

import (
	"fmt"
	"net/http"
)

type AppServer struct {
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "index.html")
	fmt.Fprintln(w, "Hello World from web server")
}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
