package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/productsvc"
)

//go:generate weaver generate ./...
func main() {
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("restock_notif", opts)
	if err != nil {
		log.Fatal(err)
	}

	pdtsvc, err := weaver.Get[productsvc.T](root)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pdtsvc.IncrementStock(r.Context(), 1, 10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	http.Serve(lis, nil)
}
