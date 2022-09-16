package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/transparentideas/gobank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(v)
	}
	return true
}
