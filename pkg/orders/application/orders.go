package application

type productsService interface {
}

type paymentsService interface {
}

type OrdersService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommand struct {
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s OrdersService) orderByID(id orders.ID) (orders.Order, error) {

}
