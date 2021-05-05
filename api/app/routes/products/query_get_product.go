package products

import (
	"net/http"
	"products-api/app/constants"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

// GetProduct query
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetProductRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	product, err := c.repository.getProduct(r.Context(), query)

	if err != nil {
		switch err {
		case goutils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, product, nil)
}

func parseGetProductRequest(r *http.Request) (*getProductQuery, error) {
	validationErrors := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[constants.FieldID])
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &getProductQuery{id}, nil
}
