package database

import (
	"fmt"
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
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

func TestFindAllIncomeDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Income{})
	incomeDB := NewIncomeDB(db)
	for i := 1; i < 4; i++ {
		incomeEntity, err := entity.NewIncome(fmt.Sprintf("Salary %d", i), 6000, "2023-05-20")
		assert.Nil(t, err)
		err = incomeDB.Create(incomeEntity)
		assert.Nil(t, err)
	}
	income, err := incomeDB.FindAll()
	assert.Nil(t, err)
	assert.Len(t, income, 3)
	assert.Equal(t, "Salary 1", income[0].Description)
	assert.Equal(t, "Salary 2", income[1].Description)
	assert.Equal(t, "Salary 3", income[2].Description)
}

func TestFindIncomeByIdDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Income{})
	incomeDB := NewIncomeDB(db)
	incomeEntity, err := entity.NewIncome("Salary", 6000, "2023-05-20")
	assert.Nil(t, err)
	err = incomeDB.Create(incomeEntity)
	assert.Nil(t, err)
	incomeFind, err := incomeDB.FindByID(incomeEntity.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, incomeFind)
	assert.Equal(t, incomeEntity.ID, incomeFind.ID)
	assert.Equal(t, incomeEntity.Description, incomeFind.Description)
	assert.Equal(t, incomeEntity.Value, incomeFind.Value)
}

func TestUpdateIncomeDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Income{})
	incomeDB := NewIncomeDB(db)
	incomeEntity, err := entity.NewIncome("Salary", 6000, "2023-05-20")
	assert.Nil(t, err)
	err = incomeDB.Create(incomeEntity)
	assert.Nil(t, err)
	incomeEntity.Description = "Overtime"
	incomeEntity.Value = 1000
	incomeEntity.Date = date.ConvertDateToTime("2023-05-22")
	err = incomeDB.Update(incomeEntity)
	assert.Nil(t, err)
	incomeUpdate, err := incomeDB.FindByID(incomeEntity.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, incomeUpdate)
	assert.Equal(t, "Overtime", incomeUpdate.Description)
	assert.Equal(t, 1000.00, incomeUpdate.Value)
	assert.Equal(t, "2023-05-22", date.ConvertDateToString(incomeUpdate.Date))
}

func TestDeleteIncomeDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Income{})
	incomeDB := NewIncomeDB(db)
	incomeEntity, err := entity.NewIncome("Salary", 6000, "2023-05-20")
	assert.Nil(t, err)
	err = incomeDB.Create(incomeEntity)
	assert.Nil(t, err)
	incomeDB.Delete(incomeEntity.ID.String())
	income, err := incomeDB.FindByID(incomeEntity.ID.String())
	assert.Nil(t, income)
	assert.NotNil(t, err)
}
