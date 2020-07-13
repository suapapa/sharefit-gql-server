package model

import "time"

type Membership struct {
	ID       string    `json:"id"`
	Training string    `json:"training"`
	CurrCnt  int       `json:"currCnt"`
	TotalCnt int       `json:"totalCnt"`
	Expiry   time.Time `json:"expiry"`
}
