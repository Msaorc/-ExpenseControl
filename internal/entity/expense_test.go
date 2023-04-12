package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExpense(t *testing.T) {
	expense, err := NewExpense("Gasoline", 100.00, "level_id_generete", "origin_id_generete", "")
	assert.Nil(t, err)
	assert.NotNil(t, expense)
	assert.NotEmpty(t, expense.ID)
	assert.NotEmpty(t, expense.Description)
	assert.Equal(t, "Gasoline", expense.Description)
	assert.NotEmpty(t, expense.Value)
	assert.Equal(t, 100.00, expense.Value)
	assert.NotEmpty(t, expense.ExpenseLevelID)
	assert.Equal(t, "level_id_generete", expense.ExpenseLevelID)
	assert.NotEmpty(t, expense.ExpenseOriginID)
	assert.Equal(t, "origin_id_generete", expense.ExpenseOriginID)
}

func TestExpenseWhenDescriptionIsRequired(t *testing.T) {
	expense, err := NewExpense("", 100.00, "level_id_generete", "origin_id_generete", "")
	assert.Nil(t, expense)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseDescriptionIsRequired, err)
}

func TestExpenseWhenValueIsRequired(t *testing.T) {
	expense, err := NewExpense("Gasoline", 0, "level_id_generete", "origin_id_generete", "")
	assert.Nil(t, expense)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseValueIsRequired, err)
}

func TestExpenseWhenValueIsInvalid(t *testing.T) {
	expense, err := NewExpense("Gasoline", -1, "level_id_generete", "origin_id_generete", "")
	assert.Nil(t, expense)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseValueIsInvalid, err)

}

func TestExpenseWhenLevelIDIsRequired(t *testing.T) {
	expense, err := NewExpense("Gasoline", 100.00, "", "origin_id_generete", "")
	assert.Nil(t, expense)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseLevelIsRequired, err)
}

func TestExpenseWhenOriginIDIsRequired(t *testing.T) {
	expense, err := NewExpense("Gasoline", 100.00, "level_id_generete", "", "")
	assert.Nil(t, expense)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseOriginIsRequired, err)
}

func TestExpenseValidate(t *testing.T) {
	expense, err := NewExpense("Gasoline", 100.00, "level_id_generete", "origin_id_generete", "")
	assert.Nil(t, err)
	assert.NotNil(t, expense)
	assert.Nil(t, expense.Validate())
}
