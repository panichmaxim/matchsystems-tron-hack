// Code generated by gowrap. DO NOT EDIT.
// template: ../../resources/opentelemetry
// gowrap: http://github.com/hexdigest/gowrap

package validator

//go:generate gowrap gen -p gitlab.com/rubin-dev/api/pkg/validator -i Validation -t ../../resources/opentelemetry -o otlp.go -l ""

import (
	"context"

	"gitlab.com/rubin-dev/api/pkg/models"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ValidationWithTracing implements Validation interface instrumented with opentracing spans
type ValidationWithTracing struct {
	Validation
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewValidationWithTracing returns ValidationWithTracing
func NewValidationWithTracing(base Validation, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) ValidationWithTracing {
	d := ValidationWithTracing{
		Validation: base,
		_instance:  instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// Address implements Validation
func (_d ValidationWithTracing) Address(ctx context.Context, value string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.Address")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"value": value}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.Address(ctx, value)
}

// AddressCount implements Validation
func (_d ValidationWithTracing) AddressCount(ctx context.Context, value string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AddressCount")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"value": value}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AddressCount(ctx, value)
}

// AuthChangePassword implements Validation
func (_d ValidationWithTracing) AuthChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthChangePassword")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthChangePassword(ctx, req)
}

// AuthForceLogin implements Validation
func (_d ValidationWithTracing) AuthForceLogin(ctx context.Context, req *models.ForceLoginRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthForceLogin")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthForceLogin(ctx, req)
}

// AuthLogin implements Validation
func (_d ValidationWithTracing) AuthLogin(ctx context.Context, req *models.LoginRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthLogin")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthLogin(ctx, req)
}

// AuthRegistration implements Validation
func (_d ValidationWithTracing) AuthRegistration(ctx context.Context, req *models.RegistrationRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthRegistration")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthRegistration(ctx, req)
}

// AuthRegistrationConfirm implements Validation
func (_d ValidationWithTracing) AuthRegistrationConfirm(ctx context.Context, req *models.TokenRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthRegistrationConfirm")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthRegistrationConfirm(ctx, req)
}

// AuthRestore implements Validation
func (_d ValidationWithTracing) AuthRestore(ctx context.Context, req *models.RestoreRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthRestore")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthRestore(ctx, req)
}

// AuthRestoreCheck implements Validation
func (_d ValidationWithTracing) AuthRestoreCheck(ctx context.Context, req *models.TokenRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthRestoreCheck")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthRestoreCheck(ctx, req)
}

// AuthRestoreConfirm implements Validation
func (_d ValidationWithTracing) AuthRestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.AuthRestoreConfirm")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.AuthRestoreConfirm(ctx, req)
}

// BtcFindAddressByHash implements Validation
func (_d ValidationWithTracing) BtcFindAddressByHash(ctx context.Context, address string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindAddressByHash")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"address": address}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindAddressByHash(ctx, address)
}

// BtcFindAllInputAndOutputByTransaction implements Validation
func (_d ValidationWithTracing) BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindAllInputAndOutputByTransaction")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"txid":     txid,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindAllInputAndOutputByTransaction(ctx, txid, page, pageSize)
}

// BtcFindBlockByHash implements Validation
func (_d ValidationWithTracing) BtcFindBlockByHash(ctx context.Context, hash string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindBlockByHash")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"hash": hash}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindBlockByHash(ctx, hash)
}

// BtcFindBlockByNumber implements Validation
func (_d ValidationWithTracing) BtcFindBlockByNumber(ctx context.Context, height int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindBlockByNumber")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"height": height}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindBlockByNumber(ctx, height)
}

// BtcFindBlockByTransaction implements Validation
func (_d ValidationWithTracing) BtcFindBlockByTransaction(ctx context.Context, txid string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindBlockByTransaction")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"txid": txid}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindBlockByTransaction(ctx, txid)
}

// BtcFindContactByAddress implements Validation
func (_d ValidationWithTracing) BtcFindContactByAddress(ctx context.Context, address string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindContactByAddress")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"address": address}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindContactByAddress(ctx, address)
}

// BtcFindIncomingTransactions implements Validation
func (_d ValidationWithTracing) BtcFindIncomingTransactions(ctx context.Context, txid string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindIncomingTransactions")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"txid":     txid,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindIncomingTransactions(ctx, txid, page, pageSize)
}

// BtcFindMentionsForAddress implements Validation
func (_d ValidationWithTracing) BtcFindMentionsForAddress(ctx context.Context, address string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindMentionsForAddress")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"address":  address,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindMentionsForAddress(ctx, address, page, pageSize)
}

// BtcFindOutcomingTransactions implements Validation
func (_d ValidationWithTracing) BtcFindOutcomingTransactions(ctx context.Context, txid string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindOutcomingTransactions")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"txid":     txid,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindOutcomingTransactions(ctx, txid, page, pageSize)
}

// BtcFindRiskScore implements Validation
func (_d ValidationWithTracing) BtcFindRiskScore(ctx context.Context, address string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindRiskScore")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"address": address}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindRiskScore(ctx, address)
}

// BtcFindTransactionByHash implements Validation
func (_d ValidationWithTracing) BtcFindTransactionByHash(ctx context.Context, address string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindTransactionByHash")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"address": address}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindTransactionByHash(ctx, address)
}

// BtcFindTransactionsByAddress implements Validation
func (_d ValidationWithTracing) BtcFindTransactionsByAddress(ctx context.Context, address string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindTransactionsByAddress")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"address":  address,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindTransactionsByAddress(ctx, address, page, pageSize)
}

// BtcFindTransactionsInBlock implements Validation
func (_d ValidationWithTracing) BtcFindTransactionsInBlock(ctx context.Context, height int, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindTransactionsInBlock")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"height":   height,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindTransactionsInBlock(ctx, height, page, pageSize)
}

// BtcFindTransactionsInBlockByHash implements Validation
func (_d ValidationWithTracing) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindTransactionsInBlockByHash")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"hash":     hash,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindTransactionsInBlockByHash(ctx, hash, page, pageSize)
}

// BtcFindWalletAddresses implements Validation
func (_d ValidationWithTracing) BtcFindWalletAddresses(ctx context.Context, wid string, page int, pageSize int) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindWalletAddresses")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"wid":      wid,
				"page":     page,
				"pageSize": pageSize}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindWalletAddresses(ctx, wid, page, pageSize)
}

// BtcFindWalletByWid implements Validation
func (_d ValidationWithTracing) BtcFindWalletByWid(ctx context.Context, wid string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindWalletByWid")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"wid": wid}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindWalletByWid(ctx, wid)
}

// BtcFindWalletForAddress implements Validation
func (_d ValidationWithTracing) BtcFindWalletForAddress(ctx context.Context, address string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.BtcFindWalletForAddress")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"address": address}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.BtcFindWalletForAddress(ctx, address)
}

// Search implements Validation
func (_d ValidationWithTracing) Search(ctx context.Context, query string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.Search")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"query": query}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.Search(ctx, query)
}

// SearchCount implements Validation
func (_d ValidationWithTracing) SearchCount(ctx context.Context, query string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.SearchCount")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":   ctx,
				"query": query}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.SearchCount(ctx, query)
}

// UserCreate implements Validation
func (_d ValidationWithTracing) UserCreate(ctx context.Context, req *models.CreateRequest) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.UserCreate")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.UserCreate(ctx, req)
}

// UserProfileUpdate implements Validation
func (_d ValidationWithTracing) UserProfileUpdate(ctx context.Context, req *models.UserProfileUpdateInput) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "Validation.UserProfileUpdate")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"req": req}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Validation.UserProfileUpdate(ctx, req)
}
