package middlewares

import (
	"net/http"

	"github.com/kristaponis/go-mini-starter/contexts"
)

// UserLogged checks if the user is logged in, and if it is, restrict access
// to page, where this middleware is applied. For ex. if the user is
// logged in, don't show /signup and /login pages.
func UserLogged(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contexts.GetUser(r.Context())
		if user != nil {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("<h3>404 Error - Not Found</h3>"))
			return
		}
		next(w, r)
	})
}
