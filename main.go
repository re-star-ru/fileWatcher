package main

import (
	"fileWatcher/orders/delivery"
	"fileWatcher/orders/usecase"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const host = ":8080"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	oUcase := usecase.New()
	oHandler := delivery.New(oUcase)
	r := chi.NewRouter()

	r.Post("/order/new", oHandler.NewOrder)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	log.Println("service started at", host)
	log.Println(http.ListenAndServe(host, r))
}
