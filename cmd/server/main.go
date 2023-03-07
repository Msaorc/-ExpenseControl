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
	config, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("file:expense.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.ExpenseOrigin{}, &entity.ExpenseLevel{}, &entity.Expense{}, &entity.User{})
	routers := chi.NewRouter()
	routers.Use(middleware.Logger)
	expenseLevelDB := database.NewExpenseLevelDB(db)
	expenseLevelHander := handlers.NewExpenseLevelHandler(expenseLevelDB)
	expenseOriginDB := database.NewExpenseOrigin(db)
	expenseOriginHander := handlers.NewExpenseOriginHandler(expenseOriginDB)
	expenseDB := database.NewExpenseDB(db)
	expenseHander := handlers.NewExpenseHandler(expenseDB)
	userHandler := handlers.NewUserHandler(database.NewUserDB(db), config.TokenAuth, config.JwtExperesIn)
	routers.Get("/expenselevel", expenseLevelHander.FindAllExpenseLevel)
	routers.Post("/expenselevel", expenseLevelHander.CreateExpenseLevel)
	routers.Get("/expenselevel/{id}", expenseLevelHander.FindExpenseLevelById)
	routers.Put("/expenselevel/{id}", expenseLevelHander.UpdateExpenseLevel)
	routers.Delete("/expenselevel/{id}", expenseLevelHander.DeleteExpenseLevel)
	routers.Get("/expenseorigin", expenseOriginHander.FindAllExpenseOrigin)
	routers.Post("/expenseorigin", expenseOriginHander.CreateExpenseOrigin)
	routers.Get("/expenseorigin/{id}", expenseOriginHander.FindExpenseOriginById)
	routers.Put("/expenseorigin/{id}", expenseOriginHander.UpdateExpenseOrigin)
	routers.Delete("/expenseorigin/{id}", expenseOriginHander.DeleteExpenseOrigin)
	routers.Get("/expense", expenseHander.FindAllExpense)
	routers.Post("/expense", expenseHander.CreateExpense)
	routers.Get("/expense/{id}", expenseHander.FindExpenseById)
	routers.Put("/expense/{id}", expenseHander.UpdateExpense)
	routers.Delete("/expense/{id}", expenseHander.DeleteExpense)
	routers.Post("/users", userHandler.CreateUser)
	routers.Post("/users/authenticate", userHandler.Authenticate)
	http.ListenAndServe(":8081", routers)
}
