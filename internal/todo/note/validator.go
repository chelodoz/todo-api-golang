package note

import "github.com/go-playground/validator"

type CustomValidator interface {
	IsValid() bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(CustomValidator)
	return value.IsValid()
}
