package http

import (
	"net/http"

	common_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/common/http"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
	products_domain "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type productsResource struct {
	repo products_domain.Repository
}

type ProductView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       PriceView `json:"price"`
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo}
	router.Get("/products/{id}", resource.Get)
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")))

	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		priceViewFromPrice(product.Price()),
	})
}
