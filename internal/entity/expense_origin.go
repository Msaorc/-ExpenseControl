package entity

import (
	"errors"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrDescriptionIsRequiredExOrigin = errors.New("description is required")

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
	if ex.Description == "" {
		return ErrDescriptionIsRequiredExOrigin
	}
	return nil
}
