// Package models implements the user model
package models

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	_ "github.com/lib/pq" // This package loads the Postgres driver
)

// User type
type User struct {
	UserID     int
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
	Password   string
	AddressN   string
	Address    string
	City       string
	Province   string
	PostalCode string
	Country    string
	Phone      string
	Host       bool
	Guest      bool
}

// UserService type provides functionality
// such as querying, creating, and updating client.
// It is the abstraction layer for our client database
type UserService struct {
	db *sql.DB
}

// NewUserService is a Factory for UserService
// that sets up a connection to our database using GORM
// It receives a string as an argument that contains all the info needed
// to make a connection to the database
// If the connection is successful, it returns a pointer to a UserService struct
func NewUserService() (*UserService, error) {
	fmt.Printf("Connecting to the db...\n")

	password, err := base64.StdEncoding.DecodeString(dbcode)
	if err != nil {
		panic(err)
	}

	// connectionInfo is a string with all the information needed to connect to a given database
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil
}

// Close is a UserService method that can be called to close the database connection
// If the connection is successfully closed, nil will be returned
// otherwise, an error is returned, which will be handled by the caller
func (us *UserService) Close() error {
	return us.db.Close()
}

// Ping is a UserService method that can be called to check if our application has
// successfully connected to the intended database.
func (us *UserService) Ping() error {
	return us.db.Ping()
}

/* ------------ Query methods and its helper functions start here ------------ */

// GetUserName is a UserService method that is called to query a single row based on a user id
// If the user is found, it returns a pointer to a user, and a nil
// If the user is not found, it returns an ErrNotFound error (part of the models package)
// If non-models error occurs, that error is returned
func (us *UserService) GetUserName(id uint) (string, error) {
	queryInfo := `
	SELECT first_name, last_name 
	FROM users
	WHERE user_id = $1`

	var firstName, lastName string
	row := us.db.QueryRow(queryInfo, id)
	if err := row.Scan(&firstName, &lastName); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrIDNotFound
		}
		return "", err
	}

	FullName := fmt.Sprintf("%s %s", strings.TrimSpace(firstName), strings.TrimSpace(lastName))

	return FullName, nil
}

// ByEmail is a UserService method that is called to query the user from the db
// using the given email of the user
// If the user is found, nil is returned
// If the user cannot be found, ErrNotFound is returned
// If non-models error occurs, that error is returned
func (us *UserService) ByEmail(email string) (*User, error) {
	queryInfo := `
	SELECT user_id, email, password FROM users
	WHERE email = $1`

	user := User{}
	row := us.db.QueryRow(queryInfo, email)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEmailNotFound
		}
		return nil, err
	}

	return &user, nil
}

/* ------------ Query methods and its helper functions end here ------------ */

/* ------------ CRUD methods and its helper functions start here ------------ */

// Create is a UserService method that receives a pointer to a User struct
// It uses the pointer to directly create a user,
// Also, it hashes a password with it using the bycrpt package
// and inserts all the user info into the database
// Returns an error, or nil if the user was successfully created and inserted into the db
func (us *UserService) Create(user *User) error {

	queryInfo := `
	INSERT INTO users (user_address, first_name, last_name, email, password, phone_number, host, guest, middle_name, branch_id)
	VALUES (
		ROW($1, $2, $3, $4, $5, $6)::address, 
		$7, $8, $9, $10, $11, $12, $13, $14, 1
	)
	RETURNING user_id;`

	row := us.db.QueryRow(queryInfo,
		user.AddressN, user.Address, user.City, user.Province, user.PostalCode, user.Country,
		user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.Host, user.Guest, user.MiddleName)

	if err := row.Scan(&user.UserID); err != nil {
		return err
	}

	fmt.Printf("New user with ID: %d was created\n", user.UserID)
	return nil
}

// Authenticate is called by SignIn method in controllers package
func (us *UserService) Authenticate(email, password string) (*User, error) {
	user := &User{}
	user, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidPassword
	}
	return user, nil
}

// VerifyEmail is used to check if the email already exists in the database
func (us *UserService) VerifyEmail(email string) error {
	user, err := us.ByEmail(email)
	if err != nil {
		if err == ErrEmailNotFound {
			return nil
		}
		return err
	}

	if user != nil {
		return ErrDuplicateEmail
	}

	return nil
}

/* ------------ CRUD methods and its helper functions end here ------------ */
