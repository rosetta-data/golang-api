package products

import (
	"encoding/json"
	"net/http"
	"products-api/app/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

// UpdateProduct command
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	command, err := parseUpdateProductRequest(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.repository.updateProduct(r.Context(), command); err != nil {
		c.logError(err)
		switch err {
		case utils.ErrNotFound:
			goutils.WriteNotFound(w, nil)
		default:
			goutils.WriteInternalError(w, nil)
		}
		return
	}

	goutils.WriteOK(w, command.Product, nil)
}

func parseUpdateProductRequest(r *http.Request) (*updateProductCommand, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	product := &product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		errorMap[utils.FieldRequestBody] = utils.ErrCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	if id != product.ID {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &updateProductCommand{product}, nil
}
