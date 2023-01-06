package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/zakariawahyu/go-fiber-mongo-clean/exception"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
)

func ValidateProduct(product model.ProductRequest) {
	err := validation.ValidateStruct(&product,
		validation.Field(&product.Name, validation.Required),
		validation.Field(&product.Price, validation.Required),
		validation.Field(&product.Quantity, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationErr{
			Message: err.Error(),
		})
	}
}
