package orders

import (
	"errors"

	"github.com/codeabuu/ECartMonolith-Microservice/pkg/common/price"
)

type ProductID string

var ErrEmptyProduct = errors.New("empty product id")

type Product struct {
	id    ProductID
	name  string
	price price.Price
}

func (p Product) ID() ProductID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error) {
	if len(id) == 0 {
		return Product{}, ErrEmptyProduct
	}
	return Product{id, name, price}, nil
}
