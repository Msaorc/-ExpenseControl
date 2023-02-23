package dto

import "time"

type ExepnseLevel struct {
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
