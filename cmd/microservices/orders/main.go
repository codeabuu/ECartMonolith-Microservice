package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/cmd"
	orders_app "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/application"
	orders_infra_orders "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/infrasructure/orders"
	orders_infra_payments "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/infrasructure/payments"
	order_infra_product "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/infrasructure/shop"
	orders_private_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/interfaces/private/http"
	orders_public_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/interfaces/public/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting orders Microservice..")

	ctx := cmd.Context()

	r, closeFn := createOrderMicroservice() //// Create the microservice components (router and services)
	defer closeFn()                         //Ensure cleanup (closing connections) when done

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-ctx.Done()
	log.Println("Closing orders microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createOrderMicroservice() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR")) //Wait until RabbitMQ is up
	shopHTTPClient := order_infra_product.NewHTTPClient(os.Getenv("SHOP_SHOP_SERVICE_ADDR"))

	ordersToPayQueue, err := orders_infra_payments.NewAMQPService(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		panic(err)
	}

	ordersRepo := orders_infra_orders.NewMemoryRepository()
	ordersService := orders_app.NewOrdersService(
		shopHTTPClient,
		ordersToPayQueue,
		ordersRepo,
	)

	r := cmd.CreateRouter()

	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)

	return r, func() {
		err := ordersToPayQueue.Close()
		if err != nil {
			log.Printf("cannot close orders queue: %s", err)
		}
	}
}
