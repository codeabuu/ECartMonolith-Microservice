package http

import (
	"net/http"

	common_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/common/http"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/application"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/orders/domain/orders"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func AddRoutes(router *chi.Mux, service application.OrdersService, repo orders.Repository) {
	resource := ordersResource{repo, service}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}

type ordersResource struct {
	repo    orders.Repository
	service application.OrdersService
}

func (o ordersResource) PostPaid(w http.ResponseWriter, r *http.Request) {
	cmd := application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(chi.URLParam(r, "id")),
	}
	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
