package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/kristaponis/go-mini-starter/contexts"
	"github.com/kristaponis/go-mini-starter/helpers"
	"github.com/kristaponis/go-mini-starter/models"
	"github.com/kristaponis/go-mini-starter/views"
	"go.mongodb.org/mongo-driver/bson"
)

type UserHandler struct {
	SignupView    *views.View
	LoginView     *views.View
	DashboardView *views.View
}

// NewUserHandler initializes user templates. This creates template cache
// by parsing templates in memory.
func NewUserHandler() *UserHandler {
	return &UserHandler{
		SignupView:    views.NewView("views/templates/user/signup.html"),
		LoginView:     views.NewView("views/templates/user/login.html"),
		DashboardView: views.NewView("views/templates/user/dashboard.html"),
	}
}

// SignupUserForm renders a page with a signup form to create a new user.
// It also adds CSRF token to the form.
// GET /signup
func (uh *UserHandler) SignupUserForm(w http.ResponseWriter, r *http.Request) {
	uh.SignupView.Render(w, r, "base", nil)
}

// SignupUser processes the form when the new user creates account.
// It parses the signup form data, creates user in the database and
// after successful creation, signs in the user to dashboard.
// POST /signup
func (uh *UserHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data from the request. If there is an error, set error message
	// and render sign up form again. Log error to console.
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		viewData := views.SetViewData(nil, helpers.NewUserError(err).Message, nil)
		uh.SignupView.Render(w, r, "base", viewData)
		return
	}

	// Pass the form data to models.User fields.
	user := models.User{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
		Created:  time.Now(),
	}

	// Create new user. If there is an error(s), set alert message
	// and render sign up form again. err is of helpers.UserError type. For data
	// persistence of the form, address of the &usr is passed in the Data field of
	// views.SetViewData, not ViewUser field. If there is an error,
	// form data (name and email) will remain after rendering signup form again.
	if err := models.NewUser().Create(&user); err != nil {
		usr := views.ViewUser{
			Name:  user.Name,
			Email: user.Email,
		}
		viewData := views.SetViewData(nil, helpers.NewUserError(err).Message, &usr)
		uh.SignupView.Render(w, r, "base", viewData)
		return
	}

	// Sign in user with cookie and set remember token.
	key := bson.D{{Key: "email", Value: user.Email}}
	if err := SignInWithCookie(w, &user, key); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// After successful sign in redirect user to dashboard.
	http.Redirect(w, r, "/user/dashboard", http.StatusFound)
}

// LoginUserForm renders a page with a form to login existing user.
// GET /login
func (uh *UserHandler) LoginUserForm(w http.ResponseWriter, r *http.Request) {
	uh.LoginView.Render(w, r, "base", nil)
}

// LoginUser parses the login form data, verifies the email and password
// if they are correct, and signs in the user to dashboard.
// POST /login
func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data from the request. If there is an error, set error message
	// and render login form again. Log error to console.
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		viewData := views.SetViewData(nil, helpers.NewUserError(err).Message, nil)
		uh.LoginView.Render(w, r, "base", viewData)
		return
	}

	// Get email and password from the form values.
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	// Authenticate checks email and password of the provided email and password.
	// If authentication is successful, return the user from the database.
	// If there is an error, set error message and render login form again.
	user, err := models.NewUser().Authenticate(email, password)
	if err != nil {
		viewData := views.SetViewData(nil, helpers.NewUserError(err).Message, email)
		uh.LoginView.Render(w, r, "base", viewData)
		return
	}

	// Sign in user with cookie and set remember token. If there is an error,
	// set error message and render login form again.
	key := bson.D{{Key: "email", Value: user.Email}}
	if err := SignInWithCookie(w, user, key); err != nil {
		viewData := views.SetViewData(nil, helpers.NewUserError(err).Message, nil)
		uh.LoginView.Render(w, r, "base", viewData)
		return
	}

	// After successful authentication and sign in, redirect to the dashboard.
	http.Redirect(w, r, "/user/dashboard", http.StatusFound)
}

// LogoutUser deletes a user session cookie (remember_token)
// and updates the user with a new remember token.
// POST /logout
func (*UserHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	// Set new cookie with empty value.
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// Get the user from the context, create new remember token and
	// hash the value.
	user := contexts.GetUser(r.Context())
	token, _ := helpers.RememberToken(64)
	rememberHash := helpers.HMACHashString(token)

	// Update remember_hash value with the newly created remember hash,
	// to replace old remember_hash value in the database after logout.
	key := bson.D{{Key: "email", Value: user.Email}}
	fields := bson.D{{Key: "$set", Value: bson.D{{Key: "remember_hash", Value: rememberHash}}}}
	models.NewUser().UpdateFields(key, fields)

	// Redirect to home page.
	http.Redirect(w, r, "/", http.StatusFound)
}

// DeleteUser deletes a user from the database.
// POST /user/delete
func (*UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Set new cookie with empty value.
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// Get the user from the context
	user := contexts.GetUser(r.Context())

	if err := models.NewUser().Delete(user.Email); err != nil {
		log.Println("error deleting user")
	}

	// Redirect to home page.
	http.Redirect(w, r, "/", http.StatusFound)
}

// DashboardUser gets user from the context and pass it to
// template as viewData. This is user only protected page.
func (uh *UserHandler) DashboardUser(w http.ResponseWriter, r *http.Request) {
	user := contexts.GetUser(r.Context())
	viewData := views.SetViewData(user, "", nil)
	uh.DashboardView.Render(w, r, "base", viewData)
}
