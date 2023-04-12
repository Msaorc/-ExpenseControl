package main

import (
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/configs"
	_ "github.com/Msaorc/ExpenseControlAPI/docs"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/Msaorc/ExpenseControlAPI/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           ExpenseControl API
// @version         1.0
// @description     API for controlling day-to-day expenses.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcos Augusto
// @contact.url    http://M&ASistem.com.br
// @contact.email  msaorc@hotmail.com

// @license.name  M&ASistem
// @license.url   http://M&ASistem.com.br

// @host      localhost:8081
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	expenseLevelDB := database.NewExpenseLevelDB(db)
	expenseLevelHander := handlers.NewExpenseLevelHandler(expenseLevelDB)
	expenseOriginDB := database.NewExpenseOriginDB(db)
	expenseOriginHander := handlers.NewExpenseOriginHandler(expenseOriginDB)
	expenseDB := database.NewExpenseDB(db)
	expenseHander := handlers.NewExpenseHandler(expenseDB)
	userHandler := handlers.NewUserHandler(database.NewUserDB(db), config.TokenAuth, config.JwtExperesIn)
	routers := chi.NewRouter()
	routers.Use(middleware.Logger)
	// routers.Use(middleware.Recoverer)

	// routers.Use(middleware.RouteHeaders().Route("Origin", "*", cors.Handler(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Content-Type"},
	// 	AllowCredentials: true, // <----------<<< do not allow credentials
	// })).Handler)

	// cors := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// })
	// routers.Use(cors.Handler)

	routers.Route("/expenselevel", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", expenseLevelHander.FindAllExpenseLevel)
		r.Post("/", expenseLevelHander.CreateExpenseLevel)
		r.Get("/{id}", expenseLevelHander.FindExpenseLevelById)
		r.Put("/{id}", expenseLevelHander.UpdateExpenseLevel)
		r.Delete("/{id}", expenseLevelHander.DeleteExpenseLevel)
	})

	routers.Route("/expenseorigin", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", expenseOriginHander.FindAllExpenseOrigin)
		r.Post("/", expenseOriginHander.CreateExpenseOrigin)
		r.Get("/{id}", expenseOriginHander.FindExpenseOriginById)
		r.Put("/{id}", expenseOriginHander.UpdateExpenseOrigin)
		r.Delete("/{id}", expenseOriginHander.DeleteExpenseOrigin)
	})

	routers.Route("/expense", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", expenseHander.FindAllExpense)
		r.Post("/", expenseHander.CreateExpense)
		r.Get("/{id}", expenseHander.FindExpenseById)
		r.Put("/{id}", expenseHander.UpdateExpense)
		r.Delete("/{id}", expenseHander.DeleteExpense)
	})

	routers.Post("/users", userHandler.CreateUser)
	routers.Post("/users/authenticate", userHandler.Authenticate)
	routers.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/docs/doc.json")))
	http.ListenAndServe(":8081", routers)
}
