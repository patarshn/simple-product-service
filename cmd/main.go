package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Init Repository...")
	repo := repository.NewProductRepository(nil)

	fmt.Println("Init Repository...")
	productService := service.NewProductService(repo)

	fmt.Println("Init Handler...")
	productHandler := handler.NewProductHandler(productService)

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}).Methods("GET")

	r.HandleFunc("/products", productHandler.GetAllProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProductByID).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
