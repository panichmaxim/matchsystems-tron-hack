package validator

import "github.com/go-ozzo/ozzo-validation/v4"

func IsValid(value interface{}, rules ...validation.Rule) (errs []string) {
	for _, rule := range rules {
		if err := validation.Validate(value, rule); err != nil {
			errs = append(errs, err.Error())
		}
	}

	return errs
}
