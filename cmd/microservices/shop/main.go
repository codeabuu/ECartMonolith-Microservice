package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("starting shop ms...")

	ctx := cmd.Context()

	r := createShopMicroservice()

	server := &http.Server{Addr: os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-ctx.Done()
	log.Println("Closing shop microservice")
	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createShopMicroservice() *chi.Mux {
	shopProductRepo := shop_infra_product.NewMemoryRepository()

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)

	return r
}
