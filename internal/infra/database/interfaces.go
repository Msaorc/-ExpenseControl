package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ExpenseOriginInterface interface {
	Create(expenserOrigin *entity.ExpenseOrigin) error
	FindAll() ([]entity.ExpenseOrigin, error)
	FindById(id string) (*entity.ExpenseOrigin, error)
	Update(expenseOrigin *entity.ExpenseOrigin, id string) error
	Delete(id string) error
}

type ExpenseLevelInterface interface {
	Create(expenseLevel *entity.ExpenseLevel) error
	FindAll() ([]entity.ExpenseLevel, error)
	FindById(id string) (*entity.ExpenseLevel, error)
	Update(expenseLevel *entity.ExpenseLevel, id string) error
	Delete(id string) error
}

type ExpenseInterface interface {
	Create(expense *entity.Expense) error
	FindAll(page, limit int, sort string) ([]entity.Expense, error)
	FindById(id string) (*entity.ExpenseLevel, error)
	Update(expense *entity.ExpenseLevel, id string) error
	Delete(id string) error
}