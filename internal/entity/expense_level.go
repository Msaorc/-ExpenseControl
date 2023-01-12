package entity

import (
	"errors"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrExpenseLevelDescriptionIsRequired = errors.New("ExpenseLevel: Description is Required")

type ExpenseLevel struct {
	ID          entity.ID `json:"id'`
	Description string    `json:"description"`
}

func NewExpenseLevel(description string) (*ExpenseLevel, error) {
	exLevel := &ExpenseLevel{
		ID:          entity.NewID(),
		Description: description,
	}
	err := exLevel.validate()
	if err != nil {
		return nil, err
	}
	return exLevel, nil
}

func (ex *ExpenseLevel) validate() error {
	if ex.Description == "" {
		return ErrExpenseLevelDescriptionIsRequired
	}
	return nil
}
