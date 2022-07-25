package validator

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/password"
)

var (
	emailRules = []validation.Rule{
		validation.Required,
		is.Email,
	}
	passwordRules = []validation.Rule{
		validation.Required,
		validation.Length(5, 250),
	}
)

type UserValidation interface {
	AuthForceLogin(ctx context.Context, req *models.ForceLoginRequest) error
	AuthLogin(ctx context.Context, req *models.LoginRequest) error
	AuthRestore(ctx context.Context, req *models.RestoreRequest) error
	AuthRestoreCheck(ctx context.Context, req *models.TokenRequest) error
	AuthRestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) error
	AuthChangePassword(ctx context.Context, req *models.ChangePasswordRequest) error
	AuthRegistration(ctx context.Context, req *models.RegistrationRequest) error
	AuthRegistrationConfirm(ctx context.Context, req *models.TokenRequest) error

	UserCreate(ctx context.Context, req *models.CreateRequest) error
	UserProfileUpdate(ctx context.Context, req *models.UserProfileUpdateInput) error
}

func (v *validationImpl) UserProfileUpdate(ctx context.Context, req *models.UserProfileUpdateInput) error {
	errs := Errors{}

	if e := IsValid(
		req.Name,
		validation.Required,
		validation.Length(3, 250),
	); len(e) > 0 {
		errs.AddErrors("name", e...)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) UserCreate(ctx context.Context, req *models.CreateRequest) error {
	errs := Errors{}

	if req.Network == nil || (req.Network != nil && *req.Network != "instagram") {
		if e := IsValid(req.Email, emailRules...); len(e) > 0 {
			errs.AddErrors("email", e...)
		}
	}

	if e := IsValid(
		req.Name,
		validation.Required,
		validation.Length(3, 250),
	); len(e) > 0 {
		errs.AddErrors("name", e...)
	}
	if req.Password != nil {
		if e := IsValid(req.Password, passwordRules...); len(e) > 0 {
			errs.AddErrors("password", e...)
		}
	}

	if len(errs) == 0 {
		user, err := v.store.UserFindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}

		if user != nil {
			errs.AddErrors("email", "user already exists")
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthForceLogin(ctx context.Context, req *models.ForceLoginRequest) error {
	errs := Errors{}

	if e := IsValid(req.ID, validation.Required); len(e) > 0 {
		errs.AddErrors("id", e...)
	} else {
		u, err := v.store.UserFindByID(ctx, req.ID)
		if err != nil {
			return err
		}
		if u == nil {
			errs.AddErrors("id", "user not found")
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthRegistrationConfirm(ctx context.Context, req *models.TokenRequest) error {
	errs := Errors{}

	if e := IsValid(req.Token, validation.Required); len(e) > 0 {
		errs.AddErrors("token", e...)
	} else {
		u, err := v.store.UserFindByToken(ctx, req.Token)
		if err != nil {
			return err
		}
		if u == nil {
			errs.AddErrors("token", "invalid token")
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthRegistration(ctx context.Context, req *models.RegistrationRequest) error {
	errs := Errors{}

	if e := IsValid(req.Name, validation.Required); len(e) > 0 {
		errs.AddErrors("name", e...)
	}
	if e := IsValid(req.Email, emailRules...); len(e) > 0 {
		errs.AddErrors("email", e...)
	} else {
		u, err := v.store.UserFindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if u != nil {
			errs.AddErrors("email", "user already exists")
		}
	}

	if e := IsValid(req.Password, passwordRules...); len(e) > 0 {
		errs.AddErrors("password", e...)
	}

	if e := IsValid(req.PasswordConfirm, passwordRules...); len(e) > 0 {
		errs.AddErrors("passwordConfirm", e...)
	} else if req.Password != req.PasswordConfirm {
		errs.AddErrors("passwordConfirm", "password do not match")
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthChangePassword(ctx context.Context, req *models.ChangePasswordRequest) error {
	errs := Errors{}

	usr, err := v.store.UserFindByID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if usr == nil {
		return fmt.Errorf("usr not found")
	}

	if e := IsValid(req.PasswordCurrent, passwordRules...); len(e) > 0 {
		errs.AddErrors("passwordCurrent", e...)
	} else {
		if usr.Password != nil && !password.CheckPasswordHash([]byte(req.Password), []byte(*usr.Password)) {
			errs.AddErrors("passwordCurrent", "invalid password")
		}
	}

	if e := IsValid(req.Password, passwordRules...); len(e) > 0 {
		errs.AddErrors("password", e...)
	}

	if e := IsValid(req.PasswordConfirm, passwordRules...); len(e) > 0 {
		errs.AddErrors("passwordConfirm", e...)
	} else if req.Password != req.PasswordConfirm {
		errs.AddErrors("passwordConfirm", "password do not match")
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthRestoreCheck(ctx context.Context, req *models.TokenRequest) error {
	errs := Errors{}

	if e := IsValid(req.Token, validation.Required); len(e) > 0 {
		errs.AddErrors("token", e...)
	}

	if len(errs) == 0 {
		u, err := v.store.UserFindByToken(ctx, req.Token)
		if err != nil {
			return err
		}
		if u == nil {
			errs.AddErrors("token", "invalid token")
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthLogin(ctx context.Context, req *models.LoginRequest) error {
	errs := Errors{}

	if e := IsValid(req.Email, emailRules...); len(e) > 0 {
		errs.AddErrors("email", e...)
	}
	if e := IsValid(req.Password, passwordRules...); len(e) > 0 {
		errs.AddErrors("password", e...)
	}

	if len(errs) == 0 {
		usr, err := v.store.UserFindByEmailActive(ctx, req.Email)
		if err != nil {
			return err
		}

		if usr == nil {
			errs.AddErrors("email", "user not found")
		} else {
			if usr.Password == nil {
				errs.AddErrors("password", "invalid password")
			} else {
				hash := []byte(*usr.Password)
				if ok := password.CheckPasswordHash([]byte(req.Password), hash); !ok {
					errs.AddErrors("password", "invalid password")
				}
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthRestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) error {
	errs := Errors{}

	if e := IsValid(req.Token, validation.Required); len(e) > 0 {
		errs.AddErrors("token", e...)
	} else {
		usr, err := v.store.UserFindByToken(ctx, req.Token)
		if err != nil {
			return err
		}

		if usr == nil {
			errs.AddErrors("token", "invalid token")
		}
	}

	if e := IsValid(req.Password, passwordRules...); len(e) > 0 {
		errs.AddErrors("password", e...)
	}

	if e := IsValid(req.PasswordConfirm, passwordRules...); len(e) > 0 {
		errs.AddErrors("passwordConfirm", e...)
	} else if req.Password != req.PasswordConfirm {
		errs.AddErrors("passwordConfirm", "password do not match")
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (v *validationImpl) AuthRestore(ctx context.Context, req *models.RestoreRequest) error {
	errs := Errors{}

	if e := IsValid(req.Email, emailRules...); len(e) > 0 {
		errs.AddErrors("email", e...)
	} else {
		usr, err := v.store.UserFindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}

		if usr == nil {
			errs.AddErrors("email", "user not found")
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
