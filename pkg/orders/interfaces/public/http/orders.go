package http

import (
	"net/http"

	common_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/common/http"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/application"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/domain/orders"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

type ordersResource struct {
	service    application.OrdersService
	repository orders.Repository
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrdersResponse struct {
	OrderID string
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	cmd := application.PlaceOrderCommand{
		OrderID:   orders.ID(uuid.NewV1().String()),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address),
	}
	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrdersResponse{
		OrderID: string(cmd.OrderID),
	})
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	render.Respond(w, r, OrderPaidView{string(order.ID()), order.Paid()})
}
