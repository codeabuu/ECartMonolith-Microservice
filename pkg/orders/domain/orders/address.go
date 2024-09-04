package orders

import "errors"

type Address struct {
	name     string
	street   string
	city     string
	postCode string
	country  string
}

func NewAddress(name string, street string, city string, postCode string, country string) (Address, error) {
	if len(name) == 0 {
		return Address{}, errors.New("Pls insert name")
	}
	if len(street) == 0 {
		return Address{}, errors.New("pls insert street")
	}
	if len(city) == 0 {
		return Address{}, errors.New("Pls insert city")
	}
	if len(postCode) == 0 {
		return Address{}, errors.New("pls insert postcode")
	}
	if len(country) == 0 {
		return Address{}, errors.New("pls insert post code")
	}
	return Address{name, street, city, postCode, country}, nil
}

func (a Address) Name() string {
	return a.name
}

func (a Address) Street() string {
	return a.street
}

func (a Address) City() string {
	return a.city
}

func (a Address) Postcode() string {
	return a.postCode
}

func (a Address) Country() string {
	return a.country
}
