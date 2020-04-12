// Package models (errors.go): all the error variables are stored in this file
package models

import "errors"

// The reason I defined custom errors for this package is so that we know that an error
// comes directly from the models package. This will make things easier if I decide to
// replace GORM with another sql package.
var (
	// ErrNotFound is returned when a resource in the db could not be found
	ErrNotFound = errors.New("models(users): Resource could not be found in the database")

	// ErrDuplicateEmail is returned when an email being used to create a new user already exists in the database
	ErrDuplicateEmail = errors.New("models(users): Email already exists in the database")

	// ErrDuplicateUsername is returned when an email being used to create a new user already exists in the database
	ErrDuplicateUsername = errors.New("models(users): Username already exists in the database")

	// ErrEmailNotFound is returned when an email cannot be found in the database
	ErrEmailNotFound = errors.New("models(users): Email could not be found in the database")

	// ErrUsernameNotFound is returned when an email cannot be found in the database
	ErrUsernameNotFound = errors.New("models(users): Username could not be found in the database")

	// ErrInvalidID is returned when an invalid ID is provided
	ErrIDNotFound = errors.New("models(users): ID could not be found in the database")

	// ErrInvalidID is returned when an invalid ID is provided (ex. id <= 0)
	ErrInvalidID = errors.New("models(users): Invalid ID entered")

	//ErrInvalidPassword indicates that a user has entered the wrong password
	ErrInvalidPassword = errors.New("models(users): Invalid password entered")
)
