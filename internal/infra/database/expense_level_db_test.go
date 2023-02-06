package database

import (
	"fmt"
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseLevelDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseLevel{})
	expenseLevelDB := NewExpenseLevelDB(db)
	expenseLevel, err := entity.NewExpenseLevel("Medium")
	assert.Nil(t, err)
	err = expenseLevelDB.Create(expenseLevel)
	assert.Nil(t, err)
	assert.NotNil(t, expenseLevel)
	assert.NotEmpty(t, expenseLevel.ID)
	assert.Equal(t, "Medium", expenseLevel.Description)
}

func TestFindAllExpenseLevelDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseLevel{})
	expenseLevelDB := NewExpenseLevelDB(db)
	for i := 1; i < 4; i++ {
		expenseLevel, err := entity.NewExpenseLevel(fmt.Sprintf("Medium %d", i))
		assert.Nil(t, err)
		err = expenseLevelDB.Create(expenseLevel)
		assert.Nil(t, err)
	}
	expensesLevel, err := expenseLevelDB.FindAll()
	assert.Nil(t, err)
	assert.Len(t, expensesLevel, 3)
	assert.Equal(t, "Medium 1", expensesLevel[0].Description)
	assert.Equal(t, "Medium 2", expensesLevel[1].Description)
	assert.Equal(t, "Medium 3", expensesLevel[2].Description)
}

func TestFindExpenseLevelDBByID(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseLevel{})
	expenseLevelDB := NewExpenseLevelDB(db)
	expenseLevel, err := entity.NewExpenseLevel("Medium")
	assert.Nil(t, err)
	err = expenseLevelDB.Create(expenseLevel)
	assert.Nil(t, err)
	expenseLeveblyId, err := expenseLevelDB.FindByID(expenseLevel.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, expenseLeveblyId)
	assert.Equal(t, expenseLevel.Description, expenseLeveblyId.Description)
}

func TestUpdateExpenseLevel(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseLevel{})
	expenseLevelDB := NewExpenseLevelDB(db)
	expenseLevel, err := entity.NewExpenseLevel("Medium")
	assert.Nil(t, err)
	err = expenseLevelDB.Create(expenseLevel)
	assert.Nil(t, err)
	expenseLevel.Description = "HIGH"
	err = expenseLevelDB.Update(expenseLevel)
	assert.Nil(t, err)
	assert.Equal(t, "HIGH", expenseLevel.Description)
}

func TestDeleteExpenseLevel(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseLevel{})
	expenseLevelDB := NewExpenseLevelDB(db)
	expenseLevel, err := entity.NewExpenseLevel("Medium")
	assert.Nil(t, err)
	err = expenseLevelDB.Create(expenseLevel)
	assert.Nil(t, err)
	err = expenseLevelDB.Delete(expenseLevel.ID.String())
	assert.Nil(t, err)
	_, err = expenseLevelDB.FindByID(expenseLevel.ID.String())
	assert.NotNil(t, err)
}
