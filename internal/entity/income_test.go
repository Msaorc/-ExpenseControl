package entity

import (
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestCreateIncome(t *testing.T) {
	income, err := NewIncome("Salary", 6000, "2023-05-01")
	assert.Nil(t, err)
	assert.NotNil(t, income)
	assert.NotNil(t, income.ID.String())
	assert.Equal(t, "Salary", income.Description)
	assert.Equal(t, 6000.00, income.Value)
	assert.Equal(t, date.ConvertDateToTime("2023-05-01"), income.Date)
}

func TestErrorWhenDescriptionIsRequired(t *testing.T) {
	income, err := NewIncome("", 6000, "2023-05-01")
	assert.Nil(t, income)
	assert.Equal(t, ErrIncomeDescriptionIsRequired, err)
}

func TestErrorWhenValueIsRequired(t *testing.T) {
	income, err := NewIncome("Salary", 0, "2023-05-01")
	assert.Nil(t, income)
	assert.Equal(t, ErrIncomeValueIsRequired, err)
}

func TestErrorWhenValueIsInvalid(t *testing.T) {
	income, err := NewIncome("Salary", -20, "2023-05-01")
	assert.Nil(t, income)
	assert.Equal(t, ErrIncomeValueIsInvalid, err)
}

func TestErrorWhenDateIsInvalid(t *testing.T) {
	income, err := NewIncome("Salary", 250, "252556-52")
	assert.Nil(t, income)
	assert.Equal(t, ErrIncomeDateIsInvalid, err)
}
