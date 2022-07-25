package jwtoken

import (
	"context"
	"net/http"
	"strings"
)

const Header = "authorization"

const key = "authorization"

type contextKey struct{ fieldName string }

var AuthorizationStringCtxKey = &contextKey{key}

func WithAuthorizationContext(ctx context.Context) string {
	v := ctx.Value(AuthorizationStringCtxKey)
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

func ParseBearerToken(token string) string {
	if token == "" {
		return ""
	}

	bearerToken := strings.Split(token, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1]
	}

	return bearerToken[0]
}

func BearerAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(Header)
		ctx := context.WithValue(r.Context(), AuthorizationStringCtxKey, ParseBearerToken(authorizationHeader))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
