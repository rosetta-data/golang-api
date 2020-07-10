package products

import (
	"net/http"
	"products-api/app/constants"
	"products-api/app/utils"
)

// DeleteProduct command
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p := newDeleteProductRequestParser(r)
	p.parse()

	if !p.IsValid() {
		utils.WriteBadRequest(w, p.ValidationErrors)
		return
	}

	if err := c.Repository.DeleteProduct(r.Context(), p.command); err != nil {
		switch err {
		case utils.ErrNotFound:
			utils.WriteNotFound(w, nil)
		default:
			utils.WriteInternalError(w, nil)
		}
		return
	}

	utils.WriteNoContent(w)
}

func newDeleteProductRequestParser(r *http.Request) *deleteProductRequestParser {
	return &deleteProductRequestParser{utils.RequestValidator{Request: r}, nil}
}

type deleteProductRequestParser struct {
	utils.RequestValidator
	command *DeleteProductCommand
}

func (v *deleteProductRequestParser) parse() {
	validationErrors := map[string]string{}

	id, err := utils.GetRequestVarInt(v.Request, constants.FieldID)
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidProductID
	}

	if 0 < len(validationErrors) {
		v.ValidationErrors = validationErrors
	} else {
		v.command = &DeleteProductCommand{id}
	}
}
