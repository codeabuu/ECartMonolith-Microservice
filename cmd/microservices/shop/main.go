package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/cmd"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/shop"
	shop_app "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/application"
	shop_infra_product "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/infrastructure/products"
	shop_interfaces_private_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/interfaces/private/http"
	shop_interfaces_public_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/interfaces/public/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting shop microservice..")

	ctx := cmd.Context()

	r := createShopMicroservice()

	server := &http.Server{Addr: os.Getenv("SHOP_SHOP_SERVICE_BIND_ADDR"), Handler: r}

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
	shopProductsService := shop_app.NewProductsService(shopProductRepo, shopProductRepo)

	if err := shop.LoadShopFixtures(shopProductsService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)

	return r
}
