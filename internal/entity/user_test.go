package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marcos Teste", "email@exepensecontrol.com", "expense")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Marcos Teste", user.Name)
	assert.Equal(t, "email@exepensecontrol.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Marcos Teste", "email@exepensecontrol.com", "expense")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("expense"))
	assert.False(t, user.ValidatePassword("control"))
	assert.NotEqual(t, user.Password, "expense")

}
