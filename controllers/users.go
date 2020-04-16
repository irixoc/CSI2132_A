package controllers

import (
	"CSI2132_A/models"
	"CSI2132_A/views"
	"fmt"
	"net/http"
)

// Users struct creats a view for signing up a user
type Users struct {
	NewView      *views.View
	SignInView   *views.View
	SignOutView  *views.View
	Search       *views.View
	HostSearch   *views.View
	MakeRental   *views.View
	us           *models.UserService
	PropertyView *views.View
	Invalid      *views.View
	Gohome       *views.View
	Expired      *views.View
	Phone        *views.View
}

// NewUsers is a factory method for Users struct
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:      views.NewView("landing", "views/users/signup.html"),
		SignInView:   views.NewView("landing", "views/users/signin.html"),
		SignOutView:  views.NewView("landing", "views/users/signout.html"),
		PropertyView: views.NewView("dashboard", "views/users/properties.html"),
		Search:       views.NewView("dashboard", "views/users/search.html"),
		HostSearch:   views.NewView("dashboard", "views/users/searchhost.html"),
		MakeRental:   views.NewView("dashboard", "views/users/rentalagreement.html"),
		Invalid:      views.NewView("invalid", "views/users/invalid.html"),
		Gohome:       views.NewView("gohome", "views/users/gohome.html"),
		Expired:      views.NewView("expired", "views/users/expired.html"),
		us:           us,
		Phone:        views.NewView("dashboard", "views/users/update.html"),
	}
}

// SignUp is called to render a view for user creation form
func (u *Users) SignUp(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

//UpdatePhone updates phone number
// func (u *Users) UpdatePhone(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "Your phone number has been updated")
// }

// Create makes a POST request to process the signup form when a
// user tries to create a new user account
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	form := SignupForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	if err := u.us.VerifyEmail(form.Email); err != nil {
		if err == models.ErrDuplicateEmail {
			fmt.Fprintf(w, "A user with the email: %s already exists. Please go back and try again.\n", form.Email)
		} else {
			fmt.Println(err)
			fmt.Fprintf(w, "You have entered an invalid email or password: %s. Please go back and try again.\n", form.Email)
		}
		return
	}

	user := &models.User{
		FirstName:  form.FirstName,
		LastName:   form.LastName,
		MiddleName: form.MiddleName,
		Email:      form.Email,
		Password:   form.Password,
		AddressN:   form.AddressN,
		Address:    form.Address,
		City:       form.City,
		Province:   form.Province,
		PostalCode: form.PostalCode,
		Country:    form.Country,
		Phone:      form.Phone,
		Host:       checkBox(form.Host),
		Guest:      checkBox(form.Guest),
	}

	if err := u.us.Create(user); err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/invalid", 301)

	}

	NewCookie(w, r, user.UserID)
	http.Redirect(w, r, "/home", 301)
}

// SignIn is used to process the sign in form when a user
// tries to log in as an existing user with their email and password
func (u *Users) SignIn(w http.ResponseWriter, r *http.Request) {

	form := SignInForm{}

	if err := parseForm(r, &form); err != nil {
		fmt.Printf("Error occured at parseForm at SignIn(controllers):\n%s\n", err)
		http.Redirect(w, r, "/gohome", 301)
	}

	user := &models.User{}

	user, err := u.us.Authenticate(form.Email, form.Password)

	switch err {
	case models.ErrEmailNotFound:
		http.Redirect(w, r, "/invalid", 301)
		return
	case models.ErrInvalidPassword:
		http.Redirect(w, r, "/invalid", 301)
		return
	case nil:
		NewCookie(w, r, user.UserID)
		http.Redirect(w, r, "/home", 301)
	default:
		http.Redirect(w, r, "/gohome", 301)
		return
	}

}

// SignOut signs the user out
func (u *Users) SignOut(w http.ResponseWriter, r *http.Request) {
	ResetCookie(w, r)
	u.SignOutView.Render(w, nil)
}
