package validators

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var ValidatorsMap = map[string]func(validator.FieldLevel) bool{
	"empty": Empty,
}

func Empty(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		return len(field.String()) == 0
	default:
		return false
	}
}
