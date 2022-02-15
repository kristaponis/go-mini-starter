package contexts

import (
	"context"

	"github.com/kristaponis/go-mini-starter/views"
)

// userPrivateKey is unexported type to be used as context key.
type privateString string

const (
	userKey privateString = "user"
)

// WithUser sets context key, it takes context and ViewUser type.
func WithUser(ctx context.Context, user *views.ViewUser) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser checks passed context value, makes type assertion to check
// if the context value is of certain type, and the returns views.ViewUser.
// This sets type of ViewUser, with less models.User fields.
func GetUser(ctx context.Context) *views.ViewUser {
	if value := ctx.Value(userKey); value != nil {
		if user, ok := value.(*views.ViewUser); ok {
			return user
		}
	}
	return nil
}
