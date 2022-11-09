package handlers

import (
	"net/http"

	"github.com/kristaponis/go-mini-starter/contexts"
	"github.com/kristaponis/go-mini-starter/views"
)

type StaticHandler struct {
	Home     *views.View
	Contacts *views.View
}

// NewStaticHandler initializes static templates. This creates template cache
// by parsing templates in memory.
func NewStaticHandler() *StaticHandler {
	return &StaticHandler{
		Home:     views.NewView("views/templates/home.html"),
		Contacts: views.NewView("views/templates/contacts.html"),
	}
}

// HomePage is main root page. It checks if the user is set,
// if it is, passes the user data to template. If the user is not set,
// it passes nil for the user and serves the page.
func (sh *StaticHandler) HomePage() http.HandlerFunc {
	// Do some stuff here one-time per-handler initialization.
	return func(w http.ResponseWriter, r *http.Request) {
		user := contexts.GetUser(r.Context())
		viewData := views.SetViewData(user, "", nil)
		sh.Home.Render(w, r, "base", viewData)
	}
}

// ContactsPage serves static /contacts page. It checks if the user is set,
// if it is, passes the user data to template. If the user is not set,
// it passes nil for the user and serves the page.
func (sh *StaticHandler) ContactsPage(w http.ResponseWriter, r *http.Request) {
	user := contexts.GetUser(r.Context())
	viewData := views.SetViewData(user, "", nil)
	sh.Contacts.Render(w, r, "base", viewData)
}

// Favicon handles serve favicon icon.
func Favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

// NotFound handles 404 error.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("<h3>404 Error - Not Found</h3>"))
}

// MethodNotAllowed handles 405 error, when wrong method is passed via request.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("<h3>405 Error - Method Not Allowed</h3>"))
}
