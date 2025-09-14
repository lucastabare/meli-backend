package main

import (
	"log"
	"net/http"
	"meli/internal/config"
	"meli/internal/repository/jsonstore"
	"meli/internal/service"
	th "meli/internal/transport/http"
)

func main() {
	cfg := config.Load()
	store := jsonstore.NewStore(cfg.DataDir)

	productRepo := jsonstore.NewProductRepo(store)
	sellerRepo  := jsonstore.NewSellerRepo(store)
	paymentRepo := jsonstore.NewPaymentRepo(store)

	productSvc := service.NewProductService(productRepo)
	sellerSvc  := service.NewSellerService(sellerRepo)
	paymentSvc := service.NewPaymentService(paymentRepo)

	productHandler := th.NewProductHandler(productSvc)
	sellerHandler  := th.NewSellerHandler(sellerSvc)
	paymentHandler := th.NewPaymentHandler(paymentSvc)

	router := th.NewRouterFull(productHandler, sellerHandler, paymentHandler)

	log.Printf("API listening on %s", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, router))
}
