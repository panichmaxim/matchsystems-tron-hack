package validator

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type NeoValidation interface {
	FindContactByAddress(ctx context.Context, address string) error
	FindAddressByHash(ctx context.Context, address string) error
	FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) error
	FindWalletForAddress(ctx context.Context, address string) error
	FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) error
	FindRiskScore(ctx context.Context, address string) error
	FindTransactionByHash(ctx context.Context, address string) error
	FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) error
	FindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) error
	FindBlockByTransaction(ctx context.Context, txid string) error
	FindBlockByNumber(ctx context.Context, height int) error
	FindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) error
	FindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) error
	FindBlockByHash(ctx context.Context, hash string) error
	FindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) error
	FindWalletByWid(ctx context.Context, wid string) error
	FindWalletAddresses(ctx context.Context, wid string, page, pageSize int) error
}

var _ NeoValidation = (*validationImpl)(nil)

var btcAddressRules = []validation.Rule{
	validation.Required,
	validation.Match(btcAddressRegex),
}

var btcTransactionRules = []validation.Rule{
	validation.Required,
	validation.Match(btcTransactionRegex),
}

var btcBlockTransactionRules = []validation.Rule{
	validation.Required,
	validation.Match(btcBlockTransactionRegex),
}

var btcHashRules = []validation.Rule{
	validation.Required,
	validation.Length(64, 64),
}

func (v *validationImpl) btcAddressValidate(ctx context.Context, address string) error {
	errs := Errors{}

	if e := IsValid(address, btcAddressRules...); len(e) > 0 {
		errs.AddErrors("address", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindContactByAddress(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) FindAddressByHash(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(address, btcAddressRules...); len(e) > 0 {
		errs.AddErrors("address", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindWalletForAddress(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(address, btcAddressRules...); len(e) > 0 {
		errs.AddErrors("address", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindRiskScore(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) FindTransactionByHash(ctx context.Context, hash string) error {
	errs := Errors{}

	if e := IsValid(hash, btcHashRules...); len(e) > 0 {
		errs.AddErrors("hash", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(txid, btcTransactionRules...); len(e) > 0 {
		errs.AddErrors("txid", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(txid, btcTransactionRules...); len(e) > 0 {
		errs.AddErrors("txid", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindBlockByTransaction(ctx context.Context, txid string) error {
	errs := Errors{}

	if e := IsValid(txid, btcTransactionRules...); len(e) > 0 {
		errs.AddErrors("txid", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindBlockByNumber(ctx context.Context, height int) error {
	errs := Errors{}

	if e := IsValid(height, validation.Required); len(e) > 0 {
		errs.AddErrors("height", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(height, validation.Required); len(e) > 0 {
		errs.AddErrors("height", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(txid, btcTransactionRules...); len(e) > 0 {
		errs.AddErrors("txid", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindBlockByHash(ctx context.Context, hash string) error {
	errs := Errors{}

	if e := IsValid(hash, btcHashRules...); len(e) > 0 {
		errs.AddErrors("hash", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(hash, btcHashRules...); len(e) > 0 {
		errs.AddErrors("hash", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindWalletByWid(ctx context.Context, wid string) error {
	errs := Errors{}

	if e := IsValid(wid, btcHashRules...); len(e) > 0 {
		errs.AddErrors("wid", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) FindWalletAddresses(ctx context.Context, wid string, page, pageSize int) error {
	errs := Errors{}

	if e := IsValid(wid, validation.Required); len(e) > 0 {
		errs.AddErrors("wid", e...)
	}

	if e := IsValid(page, validation.Required); len(e) > 0 {
		errs.AddErrors("page", e...)
	}

	if e := IsValid(pageSize, validation.Required); len(e) > 0 {
		errs.AddErrors("pageSize", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
