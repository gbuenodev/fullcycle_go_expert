package validation

import (
	"encoding/json"
	"errors"

	resterr "github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/rest_err.go"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorEn "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	trasnl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		trasnl, _ = enTransl.GetTranslator("en")
		validatorEn.RegisterDefaultTranslations(value, trasnl)
	}
}

func ValidateErr(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidation) {
		errorCauses := []resterr.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, resterr.Causes{
				Message: e.Translate(trasnl),
				Field:   e.Field(),
			})
		}

		return resterr.NewBadRequestError("Some fields are invalid", errorCauses...)
	} else {
		return resterr.NewBadRequestError("Error trying to convert fields")
	}
}
