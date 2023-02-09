package main

import (
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/configs"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/Msaorc/ExpenseControlAPI/internal/webserver/handlers"
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
	mux := http.NewServeMux()
	expenseLevelDB := database.NewExpenseLevelDB(db)
	expenseLevelHander := handlers.NewExpenseLevelHandler(expenseLevelDB)
	expenseOriginDB := database.NewExpenseOrigin(db)
	expenseOriginHander := handlers.NewExpenseOriginHandler(expenseOriginDB)
	expenseDB := database.NewExpenseDB(db)
	expenseHander := handlers.NewExpenseHandler(expenseDB)
	mux.HandleFunc("/expenselevel", expenseLevelHander.CreateExpenseLevel)
	mux.HandleFunc("/expenseorigin", expenseOriginHander.CreateExpenseOrigin)
	mux.HandleFunc("/expense", expenseHander.CreateExpense)
	http.ListenAndServe(":8081", mux)
}
