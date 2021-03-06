package models

import "github.com/simonntz/bookings/internal/forms"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
