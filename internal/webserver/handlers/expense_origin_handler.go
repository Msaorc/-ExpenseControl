package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
)

type ExpenseOriginlHandler struct {
	ExpenseOriginDB database.ExpenseOriginInterface
}

func NewExpenseOriginHandler(db database.ExpenseOriginInterface) *ExpenseOriginlHandler {
	return &ExpenseOriginlHandler{
		ExpenseOriginDB: db,
	}
}

func (eoh *ExpenseOriginlHandler) CreateExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	var expenseOrigin dto.ExepnseOrigin
	err := json.NewDecoder(r.Body).Decode(&expenseOrigin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expenseOriginEntity, err := entity.NewExpenseOrigin(expenseOrigin.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = eoh.ExpenseOriginDB.Create(expenseOriginEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}