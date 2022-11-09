package handlers

import (
	"log"
	"net/http"

	"github.com/kristaponis/go-mini-starter/helpers"
	"github.com/kristaponis/go-mini-starter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SignInWithCookie sets a session cookie for the user.
func SignInWithCookie(w http.ResponseWriter, user *models.User, key bson.D) error {
	// If user.Remember is empty string, create new remember token,
	// then hash remember token.
	if user.Remember == "" {
		token, err := helpers.RememberToken(64)
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = helpers.HMACHashString(user.Remember)

	// Set field to be updated in the database and update them.
	fields := bson.D{
		{Key: "$set", Value: bson.D{{Key: "remember_hash", Value: user.RememberHash}}},
	}
	if err := models.NewUser().UpdateFields(key, fields); err != nil {
		log.Println(err)
		http.Error(w, "Error signing user", http.StatusInternalServerError)
		return err
	}

	// Set cookie with user.Remember value.
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		Path:     "/",
		HttpOnly: true, // JavaScript can't access cookie.
	}
	http.SetCookie(w, &cookie)

	return nil
}
