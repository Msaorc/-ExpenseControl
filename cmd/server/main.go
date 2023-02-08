package main

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/configs"
	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("file:expense.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.ExpenseOrigin{}, &entity.ExpenseLevel{}, &entity.Expense{})
	http.ListenAndServe(":8080", nil)
}

type ExpenseLevelHandler struct {
	ExpenseLevelDB database.ExpenseLevelInterface
}

func NewExpenseLevelHandler(db database.ExpenseLevelInterface) *ExpenseLevelHandler {
	return &ExpenseLevelHandler{
		ExpenseLevelDB: db,
	}
}

type ExpenseOriginlHandler struct {
	ExpenseOriginDB database.ExpenseOriginInterface
}

func NewExpenseOriginHandler(db database.ExpenseOriginInterface) *ExpenseOriginlHandler {
	return &ExpenseOriginlHandler{
		ExpenseOriginDB: db,
	}
}

type ExpenseHandler struct {
	ExpenseDB database.ExpenseInterface
}

func NewExpenseHandler(db database.ExpenseInterface) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseDB: db,
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
