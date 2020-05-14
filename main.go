package main

import (
	"net/http"

	"./internal/router"
	"./internal/storages/memstore"
)

func main() {
	r := router.New(memstore.New())
	http.ListenAndServe(":8080", r.RootHandler())
}
