package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// PropertyForm is used to create new listing
type PropertyForm struct {
	PropertyType string  `schema:"propertytype"`
	Amenities    string  `schema:"amenities"`
	Bedrooms     string  `schema:"bedrooms"`
	Bathrooms    string  `schema:"bathrooms"`
	Accommodates int     `schema:"accommodates"`
	AddressN     string  `schema:"addressn"`
	Address      string  `schema:"address"`
	City         string  `schema:"city"`
	Province     string  `schema:"province"`
	PostalCode   string  `schema:"postalcode"`
	Country      string  `schema:"country"`
	RatePerDay   float64 `schema:"rpday"`
	RatePerWeek  float64 `schema:"rpweek"`
}

// SignupForm type is used by parseForm to create a struct
// that stores values to be used in the creation of a new user
type SignupForm struct {
	FirstName  string `schema:"firstname"`
	LastName   string `schema:"lastname"`
	MiddleName string `schema:"middlename"`
	Email      string `schema:"email"`
	Password   string `schema:"password"`
	AddressN   string `schema:"addressn"`
	Address    string `schema:"address"`
	City       string `schema:"city"`
	Province   string `schema:"province"`
	PostalCode string `schema:"postalcode"`
	Country    string `schema:"country"`
	Phone      string `schema:"phone"`
	Host       string `schema:"host"`
	Guest      string `schema:"guest"`
}

// SignInForm type used by parseForm to create a struct
// that stores values to be used to sign in a user
type SignInForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// parseForm takes a request pointer and uses to parse data into a form
// It must receive two pointers as arguments
func parseForm(r *http.Request, destination interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	deco := schema.NewDecoder()
	if err := deco.Decode(destination, r.PostForm); err != nil {
		return err
	}

	return nil
}

func checkBox(value string) bool {
	if value == "on" {
		return true
	}
	return false
}
