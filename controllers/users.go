package controllers

import (
	"CSI2132/models"
	"CSI2132/views"
	"fmt"
	"net/http"
)

// Users struct creats a view for signing up a user
type Users struct {
	NewView      *views.View
	SignInView   *views.View
	SignOutView  *views.View
	Search       *views.View
	us           *models.UserService
	PropertyView *views.View
}

// NewUsers is a factory method for Users struct
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:      views.NewView("landing", "views/users/signup.html"),
		SignInView:   views.NewView("landing", "views/users/signin.html"),
		SignOutView:  views.NewView("landing", "views/users/signout.html"),
		PropertyView: views.NewView("landing", "views/users/properties.html"),
		Search:       views.NewView("landing", "views/users/search.html"),
		us:           us,
	}
}

// New is called to render a view for user creation form
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create makes a POST request to process the signup form when a
// user tries to create a new user account
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	form := SignupForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	if err := u.us.VerifyEmail(form.Email); err != nil {
		if err == models.ErrDuplicateEmail {
			fmt.Fprintf(w, "The user with the email: %s already exists.\n", form.Email)
		} else {
			panic(err)
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
	fmt.Println("FORM VALUES: Host Value -->", form.Host)
	fmt.Println("FORM VALUES: Guest Value -->", form.Guest)
	fmt.Println("Host Value", user.Host)
	fmt.Println("Guest Value", user.Guest)
	if err := u.us.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)

	}

	NewCookie(w, r, user.Email)
	http.Redirect(w, r, "/home", 301)
}

// SignIn is used to process the sign in form when a user
// tries to log in as an existing user with their email and password
func (u *Users) SignIn(w http.ResponseWriter, r *http.Request) {

	form := SignInForm{}

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user := &models.User{}

	user, err := u.us.Authenticate(form.Email, form.Password)

	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email address.")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password.")
	case nil:
		NewCookie(w, r, user.Email)
		http.Redirect(w, r, "/home", 301)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// SignOut signs the user out
func (u *Users) SignOut(w http.ResponseWriter, r *http.Request) {
	ResetCookie(w, r)
	u.SignOutView.Render(w, nil)
}
