package validation

import (
	"github.com/go-playground/validator/v10"
	"library2/internal/domain/model"
)

var validate *validator.Validate

func init() {

	validate = validator.New()
}

func Validate(b model.Book) error {
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
