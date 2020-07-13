package model

import "time"

type Membership struct {
	ID string `json:"id"`
	// training represents training course for the membership
	Training string    `json:"training"`
	CurrCnt  int       `json:"currCnt"`
	TotalCnt int       `json:"totalCnt"`
	Expiry   time.Time `json:"expiry"`
	// users reperesents users who share this membership
	UserIDs []string `json:"users"`
}
