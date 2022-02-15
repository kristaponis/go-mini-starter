package middlewares

import (
	"net/http"

	"github.com/kristaponis/go-mini-starter/contexts"
	"github.com/kristaponis/go-mini-starter/models"
	"github.com/kristaponis/go-mini-starter/views"
)

// CheckUser checks if the user is logged in by checking remember_token
// cookie and comparing it with the hashed version in the database.
// If the user is found, add this user to context, if the remember_token or user
// is not found, proceed as guest site visitor and access only public pages.
func CheckUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Get remember_token cookie from the request.
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// Lookup user in the database by remember token.
		user, err := models.NewUser().ByRememberToken(cookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// If the user is found, create usr struct to hold user values. usr is
		// used to pass user values to the context down the chain and not
		// the models.User object itself. This struct replaces models.User.
		usr := &views.ViewUser{
			Name:  user.Name,
			Email: user.Email,
		}

		// Pass the usr to the context.
		ctx := r.Context()
		ctx = contexts.WithUser(ctx, usr)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
