package dto

import "time"

type ExpenseLevel struct {
	Description string `json:"description"`
	Color       string `json:"color"`
}

type ExepnseOrigin struct {
	Description string `json:"description"`
}

type Expense struct {
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	Date        time.Time `json:"date"`
	LevelID     string    `json:"level_id"`
	OringID     string    `json:"origin_id"`
	Note        string    `json:"note"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthenticateOutput struct {
	AccessToken string `json:"access_token"`
}

type Error struct {
	Message string `json:"message"`
}
