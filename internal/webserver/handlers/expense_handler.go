package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
)

type ExpenseHandler struct {
	ExpenseDB database.ExpenseInterface
}

func NewExpenseHandler(db database.ExpenseInterface) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseDB: db,
	}
}

func (e *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense dto.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	expenseEntity, err := entity.NewExpense(expense.Description, expense.Value, expense.LevelID, expense.OringID, expense.Note)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	err = e.ExpenseDB.Create(expenseEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}