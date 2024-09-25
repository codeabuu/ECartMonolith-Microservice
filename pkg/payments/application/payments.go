package application

import (
	"log"
	"time"

	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
)

type ordersService interface {
	MarkOrderAsPaid(orderID string) error
}

type PaymentsService struct {
	ordersService ordersService
}

func NewPaymentsService(ordersService ordersService) PaymentsService {
	return PaymentsService{ordersService}
}

func (s PaymentsService) InitializeOrderPayment(orderID string, price price.Price) error {
	log.Printf("initialising payment for order %s", orderID)

	go func() {
		time.Sleep(time.Millisecond * 500)
		if err := s.PostOrderPayment(orderID); err != nil {
			log.Printf("cant post order payment: %s", err)
		}
	}()
	return nil
}

func (s PaymentsService) PostOrderPayment(orderID string) error {
	log.Printf("payment for order %s done, marking orer as paid", orderID)

	return s.ordersService.MarkOrderAsPaid(orderID)
}
