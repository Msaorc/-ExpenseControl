package entity

import (
	"errors"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrExpenseOriginEntityDescriptionIsRequired = errors.New("ExpenseOrigin: description is required")
var ErrExpenseOriginEntityIDIsRequired = errors.New("ExpenseOrigin: ID is required")
var ErrExpenseOriginEntityIDIsInvalid = errors.New("ExpenseOrigin: ID is required")

type ExpenseOrigin struct {
	ID          entity.ID `json:"id"`
	Description string    `json:"description"`
}

func NewExpenseOrigin(description string) (*ExpenseOrigin, error) {
	exOrigin := &ExpenseOrigin{
		ID:          entity.NewID(),
		Description: description,
	}
	err := exOrigin.validate()
	if err != nil {
		return nil, err
	}
	return exOrigin, nil
}

func (ex *ExpenseOrigin) validate() error {
	if ex.ID.String() == "" {
		return ErrExpenseOriginEntityIDIsRequired
	}
	if _, err := entity.ParseID(ex.ID.String()); err != nil {
		return ErrExpenseOriginEntityIDIsInvalid
	}
	if ex.Description == "" {
		return ErrExpenseOriginEntityDescriptionIsRequired
	}
	return nil
}
