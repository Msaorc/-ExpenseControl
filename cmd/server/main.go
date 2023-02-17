package main

import (
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/configs"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/Msaorc/ExpenseControlAPI/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	routers := chi.NewRouter()
	routers.Use(middleware.Logger)
	expenseLevelDB := database.NewExpenseLevelDB(db)
	expenseLevelHander := handlers.NewExpenseLevelHandler(expenseLevelDB)
	expenseOriginDB := database.NewExpenseOrigin(db)
	expenseOriginHander := handlers.NewExpenseOriginHandler(expenseOriginDB)
	expenseDB := database.NewExpenseDB(db)
	expenseHander := handlers.NewExpenseHandler(expenseDB)
	routers.Post("/expenselevel", expenseLevelHander.CreateExpenseLevel)
	routers.Post("/expenseorigin", expenseOriginHander.CreateExpenseOrigin)
	routers.Get("/expenseorigin/{id}", expenseOriginHander.FindExpenseOriginById)
	routers.Put("/expenseorigin/{id}", expenseOriginHander.UpdateExpenseOrigin)
	routers.Post("/expense", expenseHander.CreateExpense)
	http.ListenAndServe(":8081", routers)
}
