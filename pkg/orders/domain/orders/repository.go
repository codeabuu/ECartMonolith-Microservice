package orders

import "errors"

var ErrNotFound = errors.New("Order not found")

type Repository interface {
	Save(*Order) error
	ByID(ID) (*Order, error)
}
