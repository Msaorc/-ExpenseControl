package database

import (
	"fmt"
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Expense{})
	expenseDB := NewExpenseDB(db)
	expense, err := entity.NewExpense("Bakery", 24.00, "level_id_test", "origin_id_test", "Buying bread and egg in Avencas")
	assert.Nil(t, err)
	err = expenseDB.Create(expense)
	assert.Nil(t, err)
	assert.NotNil(t, expense)
	assert.NotEmpty(t, expense.ID)
	assert.Equal(t, "Bakery", expense.Description)
	assert.Equal(t, 24.00, expense.Value)
	assert.Equal(t, "level_id_test", expense.LevelID)
	assert.Equal(t, "origin_id_test", expense.OringID)
	assert.Equal(t, "Buying bread and egg in Avencas", expense.Note)
}

func TestFindAllExpenseDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Expense{})
	expenseDB := NewExpenseDB(db)
	for i := 1; i < 10; i++ {
		expense, err := entity.NewExpense(fmt.Sprintf("Gasoline %d", i), 100.00, "level_id_test", "origin_id_test", fmt.Sprintf("Gasoline note %d", i))
		assert.Nil(t, err)
		err = expenseDB.Create(expense)
		assert.Nil(t, err)
	}
	expenses, err := expenseDB.FindAll(1, 3, "asc")
	assert.Nil(t, err)
	assert.Len(t, expenses, 3)
	assert.Equal(t, "Gasoline 1", expenses[0].Description)
	assert.Equal(t, "Gasoline 2", expenses[1].Description)
	assert.Equal(t, "Gasoline 3", expenses[2].Description)
}

func TestFindExpenseDBByID(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Expense{})
	expenseDB := NewExpenseDB(db)
	expense, err := entity.NewExpense("Bakery", 24.00, "level_id_test", "origin_id_test", "Buying bread and egg in Avencas")
	assert.Nil(t, err)
	err = expenseDB.Create(expense)
	assert.Nil(t, err)
	expensebyId, err := expenseDB.FindByID(expense.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, expensebyId)
	assert.Equal(t, expensebyId.Description, expensebyId.Description)
	assert.Equal(t, expensebyId.Value, expensebyId.Value)
	assert.Equal(t, expense.LevelID, expensebyId.LevelID)
	assert.Equal(t, expense.OringID, expensebyId.OringID)
	assert.Equal(t, expense.Note, expensebyId.Note)
}

func TestUpdateExpense(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Expense{})
	expenseDB := NewExpenseDB(db)
	expense, err := entity.NewExpense("Bakery", 24.00, "level_id_test", "origin_id_test", "Buying bread and egg in Avencas")
	assert.Nil(t, err)
	err = expenseDB.Create(expense)
	assert.Nil(t, err)
	expense.Description = "Gasoline"
	expense.Value = 30.00
	expense.LevelID = "level_id_test_updated"
	expense.OringID = "origin_id_test_updated"
	expense.Note = "expense update"
	err = expenseDB.Update(expense)
	assert.Nil(t, err)
	assert.Equal(t, "Gasoline", expense.Description)
	assert.Equal(t, 30.00, expense.Value)
	assert.Equal(t, "level_id_test_updated", expense.LevelID)
	assert.Equal(t, "origin_id_test_updated", expense.OringID)
	assert.Equal(t, "expense update", expense.Note)
}

func TestDeleteExpense(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.Expense{})
	expenseDB := NewExpenseDB(db)
	expense, err := entity.NewExpense("Bakery", 24.00, "level_id_test", "origin_id_test", "Buying bread and egg in Avencas")
	assert.Nil(t, err)
	err = expenseDB.Create(expense)
	assert.Nil(t, err)
	err = expenseDB.Delete(expense.ID.String())
	assert.Nil(t, err)
	_, err = expenseDB.FindByID(expense.ID.String())
	assert.NotNil(t, err)
}
