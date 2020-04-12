package controllers

import (
	"CSI2132/views"
)

// Dashboard type is used to create all the dashboard views
// which are the pages that users will use to interact with the data
type Dashboard struct {
	HomeView  *views.View
	GuestView *views.View
	HostView  *views.View
}

// NewDashboard is a factory that is used to create a new dashboard controller
func NewDashboard() *Dashboard {
	return &Dashboard{
		HomeView:  views.NewView("dashboard", "views/dashboard/home.html"),
		GuestView: views.NewView("dashboard", "views/dashboard/guest.html"),
		HostView:  views.NewView("dashboard", "views/dashboard/host.html"),
	}
}
