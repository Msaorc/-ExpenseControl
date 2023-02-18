package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	entityPKG "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/go-chi/chi"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expenseEntity, err := entity.NewExpense(expense.Description, expense.Value, expense.LevelID, expense.OringID, expense.Note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = e.ExpenseDB.Create(expenseEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (e *ExpenseHandler) FindExpenseById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expense, err := e.ExpenseDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expense)
}

func (e *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := e.ExpenseDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var expense entity.Expense
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expense.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = e.ExpenseDB.Update(&expense)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
