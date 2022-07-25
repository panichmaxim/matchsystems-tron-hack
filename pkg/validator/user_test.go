package validator

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/store"
	"testing"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthLogin(ctx, &models.LoginRequest{Email: "foo@example.com", Password: "foobar"})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"user not found"}, errs["email"])
}

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.UserCreate(ctx, &models.CreateRequest{})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Len(t, errs, 2)
	require.Equal(t, []string{"cannot be blank"}, errs["email"])
	require.Equal(t, []string{"cannot be blank"}, errs["name"])
}

func TestLoginEmpty(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthLogin(ctx, &models.LoginRequest{})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["email"])
	require.Equal(t, []string{"cannot be blank"}, errs["password"])
}

func TestForceLogin(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthForceLogin(ctx, &models.ForceLoginRequest{ID: 0})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["id"])

	usr := &models.User{Name: "test", Email: "bar@example.com"}
	require.NoError(t, r.UserCreate(ctx, usr))
	err = v.AuthForceLogin(ctx, &models.ForceLoginRequest{ID: usr.ID})
	require.NoError(t, err)
}

func TestChangePassword(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthChangePassword(ctx, &models.ChangePasswordRequest{})
	require.NotNil(t, err)

	pwd := "1234567"
	user := &models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: &pwd,
	}
	err = r.UserCreate(ctx, user)
	require.NoError(t, err)

	err = v.AuthChangePassword(ctx, &models.ChangePasswordRequest{
		UserID:          user.ID,
		PasswordCurrent: "1234567",
	})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"invalid password"}, errs["passwordCurrent"])
	require.Equal(t, []string{"cannot be blank"}, errs["password"])
	require.Equal(t, []string{"cannot be blank"}, errs["passwordConfirm"])
}

func TestRestoreCheck(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthRestoreCheck(ctx, &models.TokenRequest{})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["token"])

	err = v.AuthRestoreCheck(ctx, &models.TokenRequest{Token: "123"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Len(t, errs, 1, "errors should be empty")
	require.Equal(t, []string{"invalid token"}, errs["token"])
}

func TestRegistration(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthRegistration(ctx, &models.RegistrationRequest{PasswordConfirm: "12345"})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["email"])
	require.Equal(t, []string{"cannot be blank"}, errs["password"])
	require.Equal(t, []string{"password do not match"}, errs["passwordConfirm"], "passwordConfirm should be a required field")

	err = v.AuthRegistration(ctx, &models.RegistrationRequest{Password: "123456", PasswordConfirm: "123456"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["email"])
	require.Empty(t, errs["password"])
	require.Empty(t, errs["passwordConfirm"], "passwordConfirm should be a required field")

	usr := &models.User{Name: "test", Email: "bar@example.com"}
	require.NoError(t, r.UserCreate(ctx, usr))
	err = v.AuthRegistration(ctx, &models.RegistrationRequest{Email: usr.Email})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"user already exists"}, errs["email"])
}

func TestRegistrationConfirm(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthRegistrationConfirm(ctx, &models.TokenRequest{})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["token"])

	err = v.AuthRegistrationConfirm(ctx, &models.TokenRequest{Token: "bar@example.com"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"invalid token"}, errs["token"])
}

func TestRestore(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthRestore(ctx, &models.RestoreRequest{Email: ""})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["email"])

	err = v.AuthRestore(ctx, &models.RestoreRequest{Email: "bar@example.com"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"user not found"}, errs["email"])
}

func TestRestoreConfirm(t *testing.T) {
	ctx := context.Background()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := store.NewSQLStore(db)
	v := NewValidation(r)

	err = v.AuthRestoreConfirm(ctx, &models.RestoreConfirmRequest{PasswordConfirm: "123456"})
	require.Error(t, err)
	errs, _ := err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["token"])
	require.Equal(t, []string{"cannot be blank"}, errs["password"])
	require.Equal(t, []string{"password do not match"}, errs["passwordConfirm"])

	err = v.AuthRestoreConfirm(ctx, &models.RestoreConfirmRequest{Password: "123456", PasswordConfirm: "123456"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"cannot be blank"}, errs["token"])
	require.Empty(t, errs["password"])
	require.Empty(t, errs["passwordConfirm"])

	err = v.AuthRestoreConfirm(ctx, &models.RestoreConfirmRequest{Token: "bar@example.com"})
	require.Error(t, err)
	errs, _ = err.(Errors)
	require.Equal(t, []string{"invalid token"}, errs["token"])
}
