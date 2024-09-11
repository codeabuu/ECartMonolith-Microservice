package http

import (
	"net/http"

	common_http "github.com/codeabuu/ECartMonolith-Microservice/pkg/common/http"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
	products "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type productsResource struct {
	readModel productsReadModel
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
}

type productView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}

	render.Respond(w, r, view)
}
