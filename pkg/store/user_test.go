package store

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserListAndPaginateEmpty(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := NewSQLStore(db)

	users, count, err := r.UserListAndPaginate(ctx, 10, 0)
	require.NoError(t, err)
	require.Equal(t, count, 0)
	require.Len(t, users, 0)
}

func TestUserRepositoryUserCreate(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := NewSQLStore(db)
	u := &models.User{
		Name:     "test",
		Email:    "user@demo.com",
		IsActive: true,
	}
	require.NoError(t, r.UserCreate(ctx, u))
	require.NotEmpty(t, u.ID)
	require.NotEmpty(t, u.CreatedAt)
	require.NotEmpty(t, u.UpdatedAt)
	require.Nil(t, u.Password)
}

func TestUserStore(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	r := NewSQLStore(db)

	u := &models.User{
		Name:     "test",
		Email:    "user@demo.com",
		IsActive: true,
	}
	require.NoError(t, r.UserCreate(ctx, u))
	require.NotEmpty(t, u.CreatedAt)
	require.NotEmpty(t, u.UpdatedAt)
	require.Nil(t, u.Password)

	token := "123"
	u.Token = &token
	require.NoError(t, r.UserUpdate(ctx, u, "token"))

	uID, err := r.UserFindByID(ctx, u.ID)
	require.NoError(t, err)
	require.Equal(t, uID.ID, u.ID)

	uEmail, err := r.UserFindByEmail(ctx, u.Email)
	require.NoError(t, err)
	require.Equal(t, uEmail.ID, u.ID)

	uToken, err := r.UserFindByToken(ctx, *u.Token)
	require.NoError(t, err)
	require.Equal(t, uToken.ID, u.ID)

	users, count, err := r.UserListAndPaginate(ctx, 10, 0)
	require.NoError(t, err)
	require.Equal(t, count, 1)
	require.Len(t, users, 1)
}
