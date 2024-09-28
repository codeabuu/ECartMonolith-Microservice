package payments

import (
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/domain/orders"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/payments/interfaces/intraprocess"
)

type IntraprocessService struct {
	orders chan<- intraprocess.OrderToProcess
}

func NewIntraprocessService(ordersChannel chan<- intraprocess.OrderToProcess) IntraprocessService {
	return IntraprocessService{ordersChannel}
}

func (i IntraprocessService) InitializedOrderPayment(id orders.ID, price price.Price) error {
	i.orders <- intraprocess.OrderToProcess{
		ID:    string(id),
		Price: price,
	}
	return nil
}
