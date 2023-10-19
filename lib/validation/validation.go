package validation

import (
	"fmt"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var en ut.Translator

func init() {
	english := en_US.New()
	trans := ut.New(english, english)
	en, _ = trans.GetTranslator("en")

	validate = validator.New()

	validate.RegisterTranslation("required", en,
		func(ut ut.Translator) error {
			fmt.Println(ut)
			return ut.Add("required", "{0} is required", false)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	validate.RegisterTranslation("email", en, func(ut ut.Translator) error {
		return ut.Add("email", "Invalid email address for {0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}

func GetValidator() *validator.Validate {
	return validate
}

func ValidateStruct(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		errs := Translate(err)
		return errs[0]
	}
	return nil
}

func Translate(err error) (errs []error) {
	return translateError(err, en)
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
