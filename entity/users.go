package entity

import "time"

type User struct {
	ID             int64     `json:"id,omitempty"`
	PassportNumber string    `json:"passportNumber,omitempty"`
	PassportNum    int64     `json:"passport_num,omitempty"`
	PassportSeries int64     `json:"passport_series,omitempty"`
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserFilter struct {
	Surname   string    `json:"surname"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
