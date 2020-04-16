package models

import (
	"fmt"
)

// RentalAgreement uses cookies and form data to generate a rental agreement
type RentalAgreement struct {
	PropertyID int
	GuestID    int
}

// Property type is used to insert userts into the database
type Property struct {
	PropertyID   int
	PropertyType string
	Amenities    string
	Bedrooms     int
	Bathrooms    int
	Accommodates int
	AddressN     int
	Address      string
	City         string
	Province     string
	PostalCode   string
	Country      string
	RatePerDay   float64
	RatePerWeek  float64
	HostID       int
	PriceID      int
}

// AddProperty adds a new property into the database
func (us *UserService) AddProperty(property *Property) error {
	priceQuery := `
	INSERT INTO pricing(rate_per_day, rate_per_week)
	VALUES($1,$2) RETURNING price_id`
	priceRow := us.db.QueryRow(priceQuery, property.RatePerDay, property.RatePerWeek)

	var priceID int
	if err := priceRow.Scan(&priceID); err != nil {
		return err
	}

	queryInfo := `
	INSERT INTO properties (property_address, property_type, accommodates, amenities, bathrooms, bedrooms, host_id, pricing_id)
	VALUES (
		ROW($1, $2, $3, $4, $5, $6)::address, 
		$7, $8, $9, $10, $11, $12, $13
	)
	RETURNING property_type, bathrooms`

	row := us.db.QueryRow(queryInfo,
		property.AddressN, property.Address, property.City, property.Province, property.PostalCode,
		property.Country, property.PropertyType, property.Accommodates, property.Amenities,
		property.Bathrooms, property.Bedrooms, property.HostID, priceID)

	newProperty := Property{}
	if err := row.Scan(&newProperty.PropertyType, &newProperty.Bathrooms); err != nil {
		return err
	}

	fmt.Printf("New property of type: %s was created\n", newProperty.PropertyType)
	return nil
}

// GetProperty gets a list of all properties
func (us *UserService) GetProperty() ([]Property, error) {
	var property []Property

	queryInfo := `
	SELECT property_id, property_type, amenities, bedrooms, bathrooms, accommodates 
	FROM properties`

	rows, err := us.db.Query(queryInfo)
	defer rows.Close()

	for rows.Next() {
		prop := Property{}

		err = rows.Scan(&prop.PropertyID, &prop.PropertyType, &prop.Amenities, &prop.Bedrooms, &prop.Bathrooms, &prop.Accommodates)
		if err != nil {
			return nil, err
		}

		property = append(property, prop)
	}

	err = rows.Err() // Get any error that occurred during iterattion
	if err != nil {
		return nil, err
	}

	return property, nil
}

// GetPropertyForHost returns a list of all properties belonging to a host
func (us *UserService) GetPropertyForHost(userid int) ([]Property, error) {
	var property []Property

	queryInfo := `
	SELECT property_id, property_type, amenities, bedrooms, bathrooms, accommodates 
	FROM properties
	WHERE host_id=$1`

	rows, err := us.db.Query(queryInfo, userid)
	defer rows.Close()

	for rows.Next() {
		prop := Property{}

		err = rows.Scan(&prop.PropertyID, &prop.PropertyType, &prop.Amenities, &prop.Bedrooms, &prop.Bathrooms, &prop.Accommodates)

		if err != nil {
			return nil, err
		}

		property = append(property, prop)
	}

	err = rows.Err() // Get any error that occurred during iterattion
	if err != nil {
		return nil, err
	}

	return property, nil
}

// GetPropertyByID returns a list of all properties belonging to a host
func (us *UserService) GetPropertyByID(propertyID int) (string, *Property, error) {

	//property_id, property_type, amenities, bedrooms, bathrooms, accommodates
	queryInfo := `
	SELECT property_address, property_id, property_type, amenities, bedrooms, bathrooms, accommodates, host_id, pricing_id
	FROM properties
	WHERE property_id=$1`

	prop := &Property{}
	row := us.db.QueryRow(queryInfo, propertyID)
	var address string

	if err := row.Scan(&address, &prop.PropertyID, &prop.PropertyType, &prop.Amenities, &prop.Bedrooms, &prop.Bathrooms, &prop.Accommodates, &prop.HostID, &prop.PriceID); err != nil {
		return "", nil, err
	}

	return address, prop, nil
}

// GetRates returns the rates for that property
func (us *UserService) GetRates(priceID int) ([]float64, error) {
	queryPrice := `
	SELECT rate_per_day, rate_per_week
	FROM pricing
	WHERE price_id = $1`

	priceRow := us.db.QueryRow(queryPrice, priceID)
	var dayRate, weekRate float64
	if err := priceRow.Scan(&dayRate, &weekRate); err != nil {
		return nil, err
	}

	return []float64{dayRate, weekRate}, nil
}

//RentalAgreementB is used to insert a rental agreement
type RentalAgreementB struct {
	PropertyID int
	GuestID    int
	HostID     int
	StartDate  string
	EndDate    string
	PriceStay  string
}

// CreateAgreement pulls information from the database to generate a rental agreement
func (us *UserService) CreateAgreement(rental *RentalAgreementB) {
	queryInfo := `
	INSERT INTO rental_agreement(property_id, guest_id, signing_date, start_date, end_date)
	VALUES($1, $2, current_timestamp, current_timestamp, current_timestamp)
	RETURNING rental_id;`

	row := us.db.QueryRow(queryInfo, rental.PropertyID, rental.GuestID /*starter, ender*/)

	var rentalid int
	if err := row.Scan(&rentalid); err != nil {
		panic(err)
	}
	fmt.Println("You have successfully created a rental agreement")
}
