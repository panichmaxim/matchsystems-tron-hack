package models

type RestoreRequest struct {
	Email string `json:"email"`
}

type RestoreConfirmRequest struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForceLoginRequest struct {
	ID int64 `json:"id"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type RegistrationRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Name            string `json:"name"`
}

type Jwt struct {
	// ID user id for frontend only (readonly)
	ID int64 `json:"id"`
	// Permissions user permissions for frontend only (readonly)
	Permissions []string `json:"permissions"`
	// AccessToken
	AccessToken string `json:"accessToken"`
	// RefreshToken for renew AccessToken
	RefreshToken string `json:"refreshToken"`
}

type ChangePasswordRequest struct {
	UserID          int64  `json:"userId"`
	PasswordCurrent string `json:"passwordCurrent"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type CreateRequest struct {
	Name     string
	Email    string
	Password *string
	IsActive bool
	Network  *string
}

type UserProfileUpdateInput struct {
	Name   *string
	Avatar *string
}

type UserListRequest struct {
	Page     int
	PageSize int
}
