package middlewares

import (
	"net/http"

	"github.com/kristaponis/go-mini-starter/contexts"
)

// RequireUser middleware checks if user is logged in
// to access protected page(s). If the user is not logged in,
// redirect user to login page.
func RequireUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contexts.GetUser(r.Context())
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	})
}
