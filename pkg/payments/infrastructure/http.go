package orders

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPClient struct {
	address string
}

func NewHTTPClient(address string) HTTPClient {
	return HTTPClient{address}
}

func (h HTTPClient) MarkOrderAsPaid(OrderID string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders/%s/paid", h.address, OrderID), nil)
	if err != nil {
		return errors.Wrap(err, "cant create request")

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "cant create request")
	}
	return resp.Body.Close()
}
