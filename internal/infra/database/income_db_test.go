package database

import (
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateIncomeDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Income{})
	incomeDB := NewIncomeDB(db)
	incomeEntity, err := entity.NewIncome("Salary", 6000, "2023-05-20")
	assert.Nil(t, err)
	assert.NotNil(t, incomeEntity)
	err = incomeDB.Create(incomeEntity)
	assert.Nil(t, err)
	assert.NotNil(t, incomeEntity)
	assert.NotEmpty(t, incomeEntity.ID)
	assert.Equal(t, "Salary", incomeEntity.Description)
}
