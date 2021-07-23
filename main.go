package main

import (
	"fileWatcher/orders/delivery"
	"fileWatcher/orders/usecase"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetFlags(log.Lshortfile)

	oUcase := usecase.New("/Sync")
	oHandler := delivery.New(oUcase)
	r := chi.NewRouter()

	r.Post("/order/new", oHandler.NewOrder)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	log.Println(http.ListenAndServe(":8080", r))
}
