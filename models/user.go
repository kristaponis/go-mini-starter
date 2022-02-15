package models

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/kristaponis/go-mini-starter/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User represents the user structure in the database.
type User struct {
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	Password     string    `bson:"-"`
	PasswordHash string    `bson:"password_hash"`
	Remember     string    `bson:"-"`
	RememberHash string    `bson:"remember_hash"`
	Created      time.Time `bson:"created,omitempty"`
	Updated      time.Time `bson:"updated,omitempty"`
	Deleted      time.Time `bson:"deleted,omitempty"`
}

// NewUser initializes User type with its methods.
func NewUser() *User {
	return &User{}
}

// Create will validate user name, email and password. Then hash user password
// and create the user in the database.
func (*User) Create(user *User) error {
	// Normalize user name, email and password. Order: name, email, password.
	user.Name, user.Email, user.Password = helpers.NormalizeUserCreate(user.Name, user.Email, user.Password)

	// Validate user name, email and password.
	if err := helpers.ValidateUserCreate(user.Name, user.Email, user.Password); err != nil {
		return err
	}

	// Hash the password.
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password+os.Getenv("HASH_PEPPER")), bcrypt.DefaultCost)
	if err != nil {
		log.Println("models: error generating password hash")
		log.Println(err)
		return helpers.ErrGeneric
	}
	user.PasswordHash = string(hashed)
	user.Password = ""

	// Connect to the database.
	ctx := context.Background()
	usersColl := ConnectToDB().Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLL"))
	// Disconnect from the database.
	defer func() {
		if err = ConnectToDB().Disconnect(ctx); err != nil {
			log.Println("models: could not disconnect from the database")
			log.Println(err)
		}
	}()

	// Insert new user into the database.
	_, err = usersColl.InsertOne(ctx, user)
	if err != nil {
		log.Println("models: could not insert user into the database")
		log.Println(err)
		if mongo.IsDuplicateKeyError(err) {
			return helpers.ErrEmailDupKey
		}
		return helpers.ErrGeneric
	}

	return nil
}

// ByEmail will search the database for the user by provided email address:
// return user, nil - user found;
// return nil, ErrUserNotFound - user not found;
// return nil, ErrGeneric - user not found;
func (*User) ByEmail(e string) (*User, error) {
	var user User

	// Connect to the database.
	ctx := context.Background()
	usersColl := ConnectToDB().Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLL"))
	// Disconnect from the database.
	defer func() {
		if err := ConnectToDB().Disconnect(ctx); err != nil {
			log.Println("models: could not disconnect from the database")
			log.Println(err)
		}
	}()

	// Find user in the database.
	err := usersColl.FindOne(ctx, bson.D{{Key: "email", Value: e}}).Decode(&user)
	if err != nil {
		log.Println("models: user not found")
		log.Println(err)
		switch err {
		case mongo.ErrNoDocuments:
			return nil, helpers.ErrUserNotFound
		default:
			return nil, helpers.ErrGeneric
		}
	}

	// If the user is found, return user.
	return &user, nil
}

// ByRememberToken looks up user from database by provided remember token.
// Remember token is set while signing up user with cookie. Remember token
// is retrieved via r.Cookie("remember_token") in handlers.
func (*User) ByRememberToken(token string) (*User, error) {
	var user User
	hashedToken := helpers.HMACHashString(token)

	// Connect to the database.
	ctx := context.Background()
	usersColl := ConnectToDB().Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLL"))
	err := usersColl.FindOne(ctx, bson.D{{Key: "remember_hash", Value: hashedToken}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateFields updates user document fields in the database,
// by the provided key to find user and the fields to be updated.
func (*User) UpdateFields(key bson.D, fields bson.D) error {
	// Connect to the database.
	ctx := context.Background()
	usersColl := ConnectToDB().Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLL"))
	// Disconnect from the database.
	defer func() {
		if err := ConnectToDB().Disconnect(ctx); err != nil {
			log.Println("models: could not disconnect from the database")
			log.Println(err)
		}
	}()

	// Update user fields in the database.
	_, err := usersColl.UpdateOne(ctx, key, fields)
	if err != nil {
		log.Println("models: could not update user fields")
		log.Println(err)
		return helpers.ErrGeneric
	}

	// If user fields updated successfully, return nil for the error.
	return nil
}

// Delete user from the database.
func (*User) Delete(e string) error {
	// Validate user email.
	if err := helpers.ValidateUserEmail(e); err != nil {
		return err
	}

	// Connect to the database and delete user.
	ctx := context.Background()
	usersColl := ConnectToDB().Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLL"))
	_, err := usersColl.DeleteOne(ctx, bson.D{{Key: "email", Value: e}})
	if err != nil {
		return errors.New("models: could not delete user")
	}

	return nil
}

// AuthenticateUser checks if email and password are correct at login.
// If correct - it returns user, if not - it returns an error.
func (u *User) Authenticate(e string, p string) (*User, error) {
	// Normalize user email and password.
	e, p = helpers.NormalizeUserAuth(e, p)

	// Validate user email and password.
	if err := helpers.ValidateUserAuth(e, p); err != nil {
		return nil, err
	}

	// After successful validation, search user by email in the database.
	userOk, err := u.ByEmail(e)
	if err != nil {
		return nil, err
	}

	// Compare users hashed password in the database with the provided password.
	err = bcrypt.CompareHashAndPassword(
		[]byte(userOk.PasswordHash), []byte(p+os.Getenv("HASH_PEPPER")),
	)
	if err != nil {
		log.Println("models: password and password hash don't match")
		log.Println(err)
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, helpers.ErrPasswordMatch
		default:
			return nil, helpers.ErrGeneric
		}
	}

	return userOk, nil
}
