package products

import (
	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
	products "github.com/codeabuu/ECartMonolith-Microservice/pkg/shop/domain"
	"github.com/pkg/errors"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductsService {
	return ProductsService{repo, readModel}
}

func (s ProductsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {
	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	if err != nil {
		return errors.Wrap(err, "Invalid prod price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	if err != nil {
		return errors.Wrap(err, "Cant create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "cannot save product")
	}
	return nil
}
