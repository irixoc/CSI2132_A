package controllers

import (
	"CSI2132/models"

	"net/http"
)

// CreateProperty creates a new property
func (u *Users) CreateProperty(w http.ResponseWriter, r *http.Request) {

	form := PropertyForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	property := &models.Property{
		PropertyType: form.PropertyType,
		Amenities:    form.Amenities,
		Bedrooms:     form.Bedrooms,
		Bathrooms:    form.Bathrooms,
		Accommodates: form.Accommodates,
		AddressN:     form.AddressN,
		Address:      form.Address,
		City:         form.City,
		Province:     form.Province,
		PostalCode:   form.PostalCode,
		Country:      form.Country,
		RatePerDay:   form.RatePerDay,
		RatePerWeek:  form.RatePerWeek,
	}

	if err := u.us.AddProperty(property); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/host", 301)
}

// GetProperty gets a lists of properties
func (u *Users) GetProperty(w http.ResponseWriter, r *http.Request) {

	form := PropertyForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	property := &models.Property{
		PropertyType: form.PropertyType,
		Amenities:    form.Amenities,
		Bedrooms:     form.Bedrooms,
		Bathrooms:    form.Bathrooms,
		Accommodates: form.Accommodates,
		AddressN:     form.AddressN,
		Address:      form.Address,
		City:         form.City,
		Province:     form.Province,
		PostalCode:   form.PostalCode,
		Country:      form.Country,
		RatePerDay:   form.RatePerDay,
		RatePerWeek:  form.RatePerWeek,
	}

	if err := u.us.AddProperty(property); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/host", 301)
}
