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
	FindByID(id string) (*entity.ExpenseOrigin, error)
	Update(expenseOrigin *entity.ExpenseOrigin) error
	Delete(id string) error
}

type ExpenseLevelInterface interface {
	Create(expenseLevel *entity.ExpenseLevel) error
	FindAll() ([]entity.ExpenseLevel, error)
	FindByID(id string) (*entity.ExpenseLevel, error)
	Update(expenseLevel *entity.ExpenseLevel) error
	Delete(id string) error
}

type ExpenseInterface interface {
	Create(expense *entity.Expense) error
	FindAll(page, limit int, sort string) ([]entity.Expense, error)
	FindByID(id string) (*entity.Expense, error)
	Update(expense *entity.Expense) error
	Delete(id string) error
}

type PeriodInterface interface {
	Create(period *entity.Period) error
	FindAll() ([]entity.Period, error)
	FindByID(id string) (*entity.Period, error)
	Update(period *entity.Period) error
	Delete(id string) error
}

type IncomeInterface interface {
	Create(income *entity.Income) error
	FindAll() ([]entity.Income, error)
	FindByID(id string) (*entity.Income, error)
	Update(income *entity.Income) error
	Delete(id string) error
}
