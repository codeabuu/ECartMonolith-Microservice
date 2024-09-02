package products

import "errors"

var ErrNotFound = errors.New("Product not found")

type Repository interface {
	Save(*Product) error
	ByID(ID) (*Product, error)
}
