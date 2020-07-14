// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewMembership struct {
	Training string    `json:"training"`
	CurrCnt  int       `json:"currCnt"`
	TotalCnt int       `json:"totalCnt"`
	Expiry   time.Time `json:"expiry"`
	CenterID *string   `json:"centerID"`
}

type NewUser struct {
	Name         string  `json:"name"`
	PhoneNumber  string  `json:"phoneNumber"`
	MembershipID *string `json:"membershipID"`
}
