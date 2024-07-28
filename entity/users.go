package entity

import "time"

type User struct {
	ID             int64     `json:"id,omitempty"`
	PassportNum    int64     `json:"passport_num,omitempty"`
	PassportSeries int64     `json:"passport_series,omitempty"`
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserPassport struct {
	PassportNumber string `json:"passportNumber"`
}

type UserFilter struct {
	Surname    string    `json:"surname"`
	Name       string    `json:"name"`
	Patronymic string    `json:"patronymic""`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserInfo struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic""`
	Address    string `json:"address"`
}

func Validation(passport string) error {
	count := 0
	for i := range passport {
		if string(passport[i]) == " " {
			count += 1
		}
	}
	if count != 1 || string(passport[4]) != " " {
		return ErrValidate
	}
	return nil
}
