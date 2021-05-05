package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// CreateProduct command
func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateProductRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.createProduct(r.Context(), command); err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteCreated(w, command.Product, nil)
}

func parseCreateProductRequest(r *http.Request) (*createProductCommand, error) {
	validationErrors := map[string]string{}

	product := &product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &createProductCommand{product}, nil
}
