package controllers

import (
	"CSI2132_A/views"
)

// Landing type is used to create views for all the landing
// pages that users will use to sign in or sign up
type Landing struct {
	WelcomeView *views.View
	GuestView   *views.View
	HostView    *views.View
}

// NewLanding is a factory that is used to create a new landing controller
func NewLanding() *Landing {
	return &Landing{
		WelcomeView: views.NewView("landing", "views/landingPages/welcome.html"),
		GuestView:   views.NewView("dashboard", "views/dashboard/guest.html"),
		HostView:    views.NewView("dashboard", "views/dashboard/host.html"),
	}
}
