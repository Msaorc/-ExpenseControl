package entity

import (
	"errors"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrExpenseLevelDescriptionIsRequired = errors.New("ExpenseLevel: Description is Required")
var ErrExpenseLevelEntityIDIsRequired = errors.New("ExpenseLevel: ID is Required")
var ErrExpenseLevelEntityIDIsInvalid = errors.New("ExpenseLevel: ID is Invalid")

type ExpenseLevel struct {
	ID          entity.ID `json:"id"`
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
	if ex.ID.String() == "" {
		return ErrExpenseLevelEntityIDIsRequired
	}
	if _, err := entity.ParseID(ex.ID.String()); err != nil {
		return ErrExpenseLevelEntityIDIsInvalid
	}
	if ex.Description == "" {
		return ErrExpenseLevelDescriptionIsRequired
	}
	return nil
}