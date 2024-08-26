package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	log.Println("Starting payments ms...")

	defer log.Println("closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservice()

	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.paymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	paymentsService := payments_app.NewPaymentsService{
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	}
	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("ampq://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)

	if err != nil {
		panic(err)
	}
	return paymentsInterface
}
