package database

import (
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.User{})
	userDB := NewUserDB(db)
	user, _ := entity.NewUser("Marcos Augusto", "marcos@email.com", "200")
	err := userDB.Create(user)
	assert.Nil(t, err)
	var userFinded entity.User

	err = db.First(&userFinded, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFinded.ID)
	assert.Equal(t, user.Name, userFinded.Name)
	assert.Equal(t, user.Email, userFinded.Email)
	assert.NotNil(t, userFinded.Password)
}

func TestFindByEmail(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.User{})
	userDB := NewUserDB(db)
	user, _ := entity.NewUser("Marcos Augusto", "email@email.com", "200")
	err := userDB.Create(user)
	assert.Nil(t, err)
	userFindByEmail, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.NotNil(t, userFindByEmail)
	assert.Equal(t, user.ID, userFindByEmail.ID)
	assert.Equal(t, user.Name, userFindByEmail.Name)
	assert.Equal(t, user.Email, userFindByEmail.Email)
	assert.NotNil(t, userFindByEmail.Password)
}
