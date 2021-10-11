// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Controller struct {
	Serial    int32     `json:"serial"`
	CreatedAt time.Time `json:"createdAt"`
}

type Indication struct {
	ID               int64     `json:"id"`
	Indication       string    `json:"indication"`
	ControllerSerial int32     `json:"controllerSerial"`
	SentAt           time.Time `json:"sentAt"`
	CreatedAt        time.Time `json:"createdAt"`
}
