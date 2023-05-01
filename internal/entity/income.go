package entity

import (
	"errors"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrIncomeIDIsRequired = errors.New("Income: ID is Required")
var ErrIncomeIDIsInvalid = errors.New("Income: ID is Invalid")
var ErrIncomeDescriptionIsRequired = errors.New("Income: Description is Required")
var ErrIncomeValueIsRequired = errors.New("Income: Value is Required")
var ErrIncomeValueIsInvalid = errors.New("Income: Value is Invalid")
var ErrIncomeDateIsInvalid = errors.New("Income: Date is Invalid")

type Income struct {
	ID          entity.ID
	Description string
	Value       float64
	Date        time.Time
}

func NewIncome(description string, value float64, dateIncome string) (*Income, error) {
	if err := date.Validate(dateIncome); err != nil {
		return nil, ErrIncomeDateIsInvalid
	}
	income := &Income{
		ID:          entity.NewID(),
		Description: description,
		Value:       value,
		Date:        date.ConvertDateToTime(dateIncome),
	}
	if err := income.validate(); err != nil {
		return nil, err
	}
	return income, nil
}

func (i *Income) validate() error {
	if i.ID.String() == "" {
		return ErrIncomeIDIsRequired
	}
	if _, err := entity.ParseID(i.ID.String()); err != nil {
		return ErrIncomeIDIsInvalid
	}
	if i.Description == "" {
		return ErrIncomeDescriptionIsRequired
	}
	if i.Value < 0 {
		return ErrIncomeValueIsInvalid
	}
	if i.Value == 0 {
		return ErrIncomeValueIsRequired
	}
	return nil
}
