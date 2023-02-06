package database

import (
	"fmt"
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseOriginDB(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseOrigin{})
	expenseOriginDB := NewExpenseOrigin(db)
	expenseOrigin, err := entity.NewExpenseOrigin("CreditCard")
	assert.Nil(t, err)
	err = expenseOriginDB.Create(expenseOrigin)
	assert.Nil(t, err)
	assert.NotNil(t, expenseOrigin)
	assert.NotEmpty(t, expenseOrigin.ID)
	assert.Equal(t, "CreditCard", expenseOrigin.Description)
}

func TestFindAllExpenseOrigin(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseOrigin{})
	expenseOriginDB := NewExpenseOrigin(db)
	for i := 1; i < 4; i++ {
		expenseOrigin, err := entity.NewExpenseOrigin(fmt.Sprintf("Ticket %d", i))
		assert.Nil(t, err)
		err = expenseOriginDB.Create(expenseOrigin)
		assert.Nil(t, err)
	}

	expensesOrigin, err := expenseOriginDB.FindAll()
	assert.Nil(t, err)
	assert.Len(t, expensesOrigin, 3)
	assert.Equal(t, "Ticket 1", expensesOrigin[0].Description)
	assert.Equal(t, "Ticket 2", expensesOrigin[1].Description)
	assert.Equal(t, "Ticket 3", expensesOrigin[2].Description)
}

func TestFindExpenseOriginByID(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseOrigin{})
	expenseOriginDB := NewExpenseOrigin(db)
	expenseOrigin, err := entity.NewExpenseOrigin("Pix")
	assert.Nil(t, err)
	err = expenseOriginDB.Create(expenseOrigin)
	assert.Nil(t, err)
	expenseOriginFind, err := expenseOriginDB.FindByID(expenseOrigin.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, expenseOriginFind)
	assert.Equal(t, expenseOrigin.Description, expenseOriginFind.Description)
}

func TestUpdateExpenseOrigin(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseOrigin{})
	expenseOriginDB := NewExpenseOrigin(db)
	expenseOrigin, err := entity.NewExpenseOrigin("Ticket")
	assert.Nil(t, err)
	err = expenseOriginDB.Create(expenseOrigin)
	assert.Nil(t, err)
	expenseOrigin.Description = "CreditCard"
	err = expenseOriginDB.Update(expenseOrigin)
	assert.Nil(t, err)
	assert.Equal(t, "CreditCard", expenseOrigin.Description)
}

func TestDeleteExpenseOrigin(t *testing.T) {
	db := CreateTableAndConnectionBD(entity.ExpenseOrigin{})
	expenseOriginDB := NewExpenseOrigin(db)
	expenseOrigin, err := entity.NewExpenseOrigin("Ticket")
	assert.Nil(t, err)
	err = expenseOriginDB.Create(expenseOrigin)
	assert.Nil(t, err)
	err = expenseOriginDB.Delete(expenseOrigin.ID.String())
	assert.Nil(t, err)
	_, err = expenseOriginDB.FindByID(expenseOrigin.ID.String())
	assert.NotNil(t, err)
}
