package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseOrigin(t *testing.T) {
	exOrigin, err := NewExpenseOrigin("credit card")
	assert.Nil(t, err)
	assert.NotNil(t, exOrigin)
	assert.NotEmpty(t, exOrigin.ID)
	assert.Equal(t, "credit card", exOrigin.Description)
}

func TestExpenseOriginWhenDescriptionIsRequired(t *testing.T) {
	exOrigin, err := NewExpenseOrigin("")
	assert.Nil(t, exOrigin)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDescriptionIsRequired, err)
}

func TestExpenseOriginValidate(t *testing.T) {
	exOrigin, err := NewExpenseOrigin("credit card")
	assert.Nil(t, err)
	assert.NotNil(t, exOrigin)
	assert.Nil(t, exOrigin.validate())
}