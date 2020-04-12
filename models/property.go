package models

import "fmt"

// Property type is used to insert userts into the database
type Property struct {
	PropertyType string
	Amenities    string
	Bedrooms     string
	Bathrooms    string
	Accommodates int
	AddressN     string
	Address      string
	City         string
	Province     string
	PostalCode   string
	Country      string
	RatePerDay   float64
	RatePerWeek  float64
	// Host         int
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
		property.AddressN, property.Address, property.City, property.Province, property.PostalCode, property.Country,
		property.PropertyType, property.Accommodates, property.Amenities, property.Bathrooms, property.Bedrooms, 12, priceID)

	newProperty := Property{}
	if err := row.Scan(&newProperty.PropertyType, &newProperty.Bathrooms); err != nil {
		return err
	}

	fmt.Printf("New property of type: %s was created\n", newProperty.PropertyType)
	return nil
}

// GetProperty gets a list of all properties
func (us *UserService) GetProperty() error {

	queryInfo := `SELECT * FROM properties`

	rows, err := us.db.Query(queryInfo)
	defer rows.Close()

	for rows.Next() {
		property := &Property{}
		err = rows.Scan(&property.PropertyType, &property.Amenities, &property.Bedrooms, &property.Bathrooms, &property.Accommodates,
			&property.AddressN, &property.Address, &property.City, &property.Province, &property.PostalCode, &property.Country,
			&property.RatePerDay, &property.RatePerWeek)
		if err != nil {
			// handle this error
			return err
		}
		fmt.Println(property)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
