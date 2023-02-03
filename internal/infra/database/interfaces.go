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
	Update(expense *entity.ExpenseOrigin, id string) error
	Delete(id string) error
}
