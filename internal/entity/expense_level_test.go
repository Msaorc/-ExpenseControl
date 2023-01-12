package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExpenseLevel(t *testing.T) {
	exLevel, err := NewExpenseLevel("Medium")
	assert.Nil(t, err)
	assert.NotNil(t, exLevel)
	assert.NotEmpty(t, exLevel.ID)
	assert.Equal(t, "Medium", exLevel.Description)
}

func TestExpenseLevelWhenDescriptionIsRequired(t *testing.T) {
	exOrigin, err := NewExpenseLevel("")
	assert.Nil(t, exOrigin)
	assert.NotNil(t, err)
	assert.Equal(t, ErrExpenseLevelDescriptionIsRequired, err)
}

func TestExpenseLevelValidate(t *testing.T) {
	exLevel, err := NewExpenseLevel("Emergency")
	assert.Nil(t, err)
	assert.NotNil(t, exLevel)
	assert.Nil(t, exLevel.validate())
}
