package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	enl := en.New()
	unt := ut.New(enl, enl)
	transl, _ = unt.GetTranslator("en")

	en_translation.RegisterDefaultTranslations(Validate, transl)
}

func ValidateErr(err error) *rest.RestErr {
	var json_err *json.UnmarshalTypeError
	var validation_err validator.ValidationErrors

	if errors.As(err, &json_err) {
		return rest.NewBadRequestErr("invalid field type", nil)
	} else if errors.As(err, &validation_err) {
		causes := []rest.Cause{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := rest.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			causes = append(causes, cause)
		}

		return rest.NewBadRequestErr("some field are invalid", causes)
	} else {
		return rest.NewBadRequestErr("error trying to convert fields", nil)
	}
}

func ValidateVar(name string, err error) *rest.RestErr {
	var json_err *json.UnmarshalTypeError
	var validation_err validator.ValidationErrors

	if errors.As(err, &json_err) {
		return rest.NewBadRequestErr("invalid field type", nil)
	} else if errors.As(err, &validation_err) {
		causes := []rest.Cause{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := rest.Cause{
				Message: fmt.Sprintf("%s%s", name, e.Translate(transl)),
				Field:   name,
			}

			causes = append(causes, cause)
		}

		return rest.NewBadRequestErr("some field are invalid", causes)
	} else {
		return rest.NewBadRequestErr("error trying to convert fields", nil)
	}
}
