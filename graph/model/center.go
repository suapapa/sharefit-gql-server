package model

type Center struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	// Memberships []*Membership `json:"memberships"`
}
