package main

import (
	"net/http"

	"github.com/jcbritobr/gorillaapi/router"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	router.SetupRouter(m)

	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}
