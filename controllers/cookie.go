package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	//ErrCookieGone means an error occurred when looking for that cookie
	ErrCookieGone = errors.New("Invalid Cookie or Cookie Not Found")

	//ErrEmptyCookie means the cookie value is an empty string
	ErrEmptyCookie = errors.New("No value was found for this cookie")
)

//NewCookie makes and sets a new cookie
func NewCookie(w http.ResponseWriter, r *http.Request, userid int) {

	cookie := &http.Cookie{Name: "userid", Value: strconv.Itoa(userid)}
	http.SetCookie(w, cookie)
}

//ResetCookie deletes a cookie
func ResetCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RESETING COOKIE............................")

	cookie := &http.Cookie{
		Name:    "userid",
		Value:   "",
		Expires: time.Now(),
		MaxAge:  -1,
	}

	http.SetCookie(w, cookie) // resets a cookie
	// http.Redirect(w, r, "/signout", 301)
}

// GetCookie returns a cookie
func GetCookie(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("userid")

	if err != nil {
		fmt.Println("GetCookie error............................")
		return nil, ErrCookieGone
	}

	if cookie.Value == "" {
		fmt.Println("Empty Cookie error............................")
		return nil, ErrEmptyCookie
	}

	return cookie, nil
}
