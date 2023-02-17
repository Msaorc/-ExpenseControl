package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
)

type ExpenseLevelHandler struct {
	ExpenseLevelDB database.ExpenseLevelInterface
}

func NewExpenseLevelHandler(db database.ExpenseLevelInterface) *ExpenseLevelHandler {
	return &ExpenseLevelHandler{
		ExpenseLevelDB: db,
	}
}

func (elh *ExpenseLevelHandler) CreateExpenseLevel(w http.ResponseWriter, r *http.Request) {
	var expenseLevel dto.ExepnseLevel
	err := json.NewDecoder(r.Body).Decode(&expenseLevel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ExpenseLevelEntity, err := entity.NewExpenseLevel(expenseLevel.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = elh.ExpenseLevelDB.Create(ExpenseLevelEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
