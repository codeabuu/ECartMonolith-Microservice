package products

import (
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/infrastructure/products"
)

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() {

}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, cmd.PriceCents)

	s.repo.Save
}
