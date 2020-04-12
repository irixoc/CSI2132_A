package controllers

import (
	"fmt"
	"net/http"
)

//NewCookie makes and sets a new cookie
func NewCookie(w http.ResponseWriter, r *http.Request, email string) {
	cookie := &http.Cookie{Name: "email", Value: email}
	http.SetCookie(w, cookie)
}

//ResetCookie deletes a cookie
func ResetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	fmt.Println("RESETING COOKIE")
	if err != nil || cookie.Value == "" {
		fmt.Println("****** COOKIE ERROR ******")
		http.Redirect(w, r, "/", 301)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
