package validator

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var _ BtcNeoValidation = (*validationImpl)(nil)

type BtcNeoValidation interface {
	BtcFindContactByAddress(ctx context.Context, address string) error
	BtcFindAddressByHash(ctx context.Context, address string) error
	BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) error
	BtcFindWalletForAddress(ctx context.Context, address string) error
	BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) error
	BtcFindRiskScore(ctx context.Context, address string) error
	BtcFindTransactionByHash(ctx context.Context, address string) error
	BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) error
	BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) error
	BtcFindBlockByTransaction(ctx context.Context, txid string) error
	BtcFindBlockByNumber(ctx context.Context, height int) error
	BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) error
	BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) error
	BtcFindBlockByHash(ctx context.Context, hash string) error
	BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) error
	BtcFindWalletByWid(ctx context.Context, wid string) error
	BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) error
}

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

func (v *validationImpl) BtcFindContactByAddress(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) BtcFindAddressByHash(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindWalletForAddress(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindRiskScore(ctx context.Context, address string) error {
	return v.btcAddressValidate(ctx, address)
}

func (v *validationImpl) BtcFindTransactionByHash(ctx context.Context, hash string) error {
	errs := Errors{}

	if e := IsValid(hash, btcHashRules...); len(e) > 0 {
		errs.AddErrors("hash", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindBlockByTransaction(ctx context.Context, txid string) error {
	errs := Errors{}

	if e := IsValid(txid, btcTransactionRules...); len(e) > 0 {
		errs.AddErrors("txid", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) BtcFindBlockByNumber(ctx context.Context, height int) error {
	errs := Errors{}

	if e := IsValid(height, validation.Required); len(e) > 0 {
		errs.AddErrors("height", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) error {
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

func (v *validationImpl) BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindBlockByHash(ctx context.Context, hash string) error {
	errs := Errors{}

	if e := IsValid(hash, btcHashRules...); len(e) > 0 {
		errs.AddErrors("hash", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) error {
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

func (v *validationImpl) BtcFindWalletByWid(ctx context.Context, wid string) error {
	errs := Errors{}

	if e := IsValid(wid, validation.Required); len(e) > 0 {
		errs.AddErrors("wid", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) error {
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
