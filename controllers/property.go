package controllers

import (
	"CSI2132_A/models"
	"fmt"
	"strconv"
	"strings"

	"net/http"
)

// CreateProperty creates a new property
func (u *Users) CreateProperty(w http.ResponseWriter, r *http.Request) {

	form := PropertyForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	var cookieIntValue int
	cookie, err := GetCookie(w, r)
	switch err {
	case ErrCookieGone:
		http.Redirect(w, r, "/expired", 301)
	case ErrEmptyCookie:
		http.Redirect(w, r, "/expired", 301)
	case nil:
		cookieIntValue, err = strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("CreatePropety(controllers) cookie string to int problems: ", err)
			http.Redirect(w, r, "/expired", 301)
			return
		}
	default:
		fmt.Println("CreatePropety(controllers) cookie problems: ", err)
		http.Redirect(w, r, "/expired", 301)
	}
	if err != nil {
		return
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
		HostID:       cookieIntValue,
	}

	if err := u.us.AddProperty(property); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/host", 301)
}

// SearchData is used to return list of properties to users
type SearchData struct {
	Property []models.Property
}

// GetProperty gets a lists of properties
func (u *Users) GetProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	props, err := u.us.GetProperty()

	Search := &SearchData{
		Property: props,
	}

	if err != nil {
		panic(err)
	}
	// fmt.Println(&Search.Property)

	u.Search.Render(w, Search)
}

// GetPropertyForHost gets a lists of properties
func (u *Users) GetPropertyForHost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var cookieUserID int
	cookie, err := GetCookie(w, r)
	switch err {
	case ErrCookieGone:
		http.Redirect(w, r, "/expired", 301)
	case ErrEmptyCookie:
		http.Redirect(w, r, "/expired", 301)
	case nil:
		cookieUserID, err = strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("CreatePropety(controllers) cookie string to int problems: ", err)
			http.Redirect(w, r, "/expired", 301)
			return
		}
	default:
		fmt.Println("CreatePropety(controllers) cookie problems: ", err)
		http.Redirect(w, r, "/expired", 301)
	}
	if err != nil {
		return
	}

	props, err := u.us.GetPropertyForHost(cookieUserID)

	Search := &SearchData{
		Property: props,
	}

	if err != nil {
		panic(err)
	}
	// fmt.Println(&Search.Property)

	u.HostSearch.Render(w, Search)
}

// RentalInfo is used to generate the rental info for a guest
type RentalInfo struct {
	PropertyID    int
	GuestFullName string
	HostFullName  string
	GuestID       int
	HostID        int
	Number        string
	Street        string
	City          string
	Province      string
	Postal        string
	Country       string
	PropType      string
	Amenities     string
	Bedrooms      int
	Bathrooms     int
	Accommodates  int
	DayRate       float64
	WeekRate      float64
}

// CreateRental generates a rental agreement
func (u *Users) CreateRental(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	form := RentalForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	var cookieIntValue int
	cookie, err := GetCookie(w, r)
	switch err {
	case ErrCookieGone:
		http.Redirect(w, r, "/expired", 301)
	case ErrEmptyCookie:
		http.Redirect(w, r, "/expired", 301)
	case nil:
		cookieIntValue, err = strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("CreateRental(controllers) cookie string to int problems: ", err)
			http.Redirect(w, r, "/expired", 301)
			return
		}
	default:
		fmt.Println("CreateRental(controllers) cookie problems: ", err)
		http.Redirect(w, r, "/expired", 301)
	}
	if err != nil {
		return
	}

	FullName, err := u.us.GetUserName(uint(cookieIntValue))

	if err != nil {
		fmt.Println("CreateRental(controllers) GetUserName: ", err)
		http.Redirect(w, r, "/expired", 301)
		return
	}

	address, pInfo, err := u.us.GetPropertyByID(form.PropertyID)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "There appears to be a problem, please go back and try again.")
	}
	address = strings.ReplaceAll(address, "(", "")
	address = strings.ReplaceAll(address, ")", "")
	a := strings.Split(address, ",")

	rates, err := u.us.GetRates(pInfo.PriceID)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "There appears to be a problem, please go back and try again.")
	}

	HostName, err := u.us.GetUserName(uint(pInfo.HostID))
	if err != nil {
		fmt.Println("CreateRental(controllers) GetUserName: ", err)
		http.Redirect(w, r, "/expired", 301)
		return
	}

	renting := &RentalInfo{
		PropertyID:    form.PropertyID,
		GuestFullName: FullName,
		HostFullName:  HostName,
		GuestID:       cookieIntValue,
		HostID:        pInfo.HostID,
		Number:        a[0],
		Street:        a[1],
		City:          a[2],
		Province:      a[3],
		Postal:        a[4],
		Country:       a[5],
		PropType:      pInfo.PropertyType,
		Amenities:     pInfo.Amenities,
		Bedrooms:      pInfo.Bedrooms,
		Bathrooms:     pInfo.Bathrooms,
		Accommodates:  pInfo.Accommodates,
		DayRate:       rates[0],
		WeekRate:      rates[1],
	}

	u.MakeRental.Render(w, renting)
}

//CreateAgreement will generate the rental agreement
func (u *Users) CreateAgreement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	agreement := &RentalAgreement{}
	if err := parseForm(r, agreement); err != nil {
		panic(err)
	}
	fmt.Println(agreement)

	newAgreement := &models.RentalAgreementB{
		PropertyID: agreement.PropertyID,
		GuestID:    agreement.GuestID,
		StartDate:  agreement.StartDate,
		EndDate:    agreement.EndDate,
		PriceStay:  agreement.PriceStay,
	}

	fmt.Printf("%T\n", newAgreement.StartDate)
	u.us.CreateAgreement(newAgreement)
}
