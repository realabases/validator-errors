package validatorerrors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type RuleFunc func(e validator.FieldError) string

type ValidatorErrors struct {
	rules map[string]RuleFunc
}

func New() *ValidatorErrors {
	return &ValidatorErrors{
		rules: map[string]RuleFunc{},
	}
}

func (v *ValidatorErrors) AddRule(tag string, fn RuleFunc) {
	v.rules[tag] = fn
}

func (v *ValidatorErrors) RemoveRule(tag string) {
	delete(v.rules, tag)
}

func (v *ValidatorErrors) AddDefaultRule(tag string) {
	switch tag {
	case "required":
		v.rules[tag] = func(e validator.FieldError) string {
			return fmt.Sprintf("%s is required", e.Field())
		}
	case "min":
		v.rules[tag] = func(e validator.FieldError) string {
			return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
		}
	case "max":
		v.rules[tag] = func(e validator.FieldError) string {
			return fmt.Sprintf("%s can't be more than %s characters", e.Field(), e.Param())
		}
	case "email":
		v.rules[tag] = func(e validator.FieldError) string {
			return fmt.Sprintf("%s must be a valid email", e.Field())
		}
	default:
		v.rules[tag] = func(e validator.FieldError) string {
			return fmt.Sprintf("%s is invalid", e.Field())
		}
	}
}

func (v *ValidatorErrors) FormatValidationErrors(err error) map[string]string {
	errors := map[string]string{}

	for _, e := range err.(validator.ValidationErrors) {
		if msgFunc, ok := v.rules[e.Tag()]; ok {
			errors[e.Field()] = msgFunc(e)
		} else {
			errors[e.Field()] = "invalid value"
		}
	}

	return errors
}
