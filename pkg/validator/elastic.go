package validator

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ElasticValidation interface {
	Search(ctx context.Context, query string) error
	SearchCount(ctx context.Context, query string) error
	Address(ctx context.Context, value string) error
	AddressCount(ctx context.Context, value string) error
}

var _ ElasticValidation = (*validationImpl)(nil)

var elkQueryRules = []validation.Rule{
	validation.Required,
	validation.Length(2, 250),
}

func (v *validationImpl) Search(ctx context.Context, value string) error {
	errs := Errors{}

	if e := IsValid(value, elkQueryRules...); len(e) > 0 {
		errs.AddErrors("query", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) SearchCount(ctx context.Context, value string) error {
	errs := Errors{}

	if e := IsValid(value, elkQueryRules...); len(e) > 0 {
		errs.AddErrors("query", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) Address(ctx context.Context, value string) error {
	errs := Errors{}

	if e := IsValid(value, elkQueryRules...); len(e) > 0 {
		errs.AddErrors("query", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AddressCount(ctx context.Context, value string) error {
	errs := Errors{}

	if e := IsValid(value, elkQueryRules...); len(e) > 0 {
		errs.AddErrors("query", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
