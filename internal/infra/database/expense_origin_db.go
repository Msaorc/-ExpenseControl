package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/gorm"
)

type ExpenseOrigin struct {
	DB *gorm.DB
}

func NewExpenseOrigin(db *gorm.DB) *ExpenseOrigin {
	return &ExpenseOrigin{DB: db}
}

func (eo *ExpenseOrigin) Create(expenserOrigin *entity.ExpenseOrigin) error {
	return eo.DB.Create(expenserOrigin).Error
}

func (eo *ExpenseOrigin) FindAll() ([]entity.ExpenseOrigin, error) {
	var expenseOrigins []entity.ExpenseOrigin
	err := eo.DB.Find(&expenseOrigins).Error
	return expenseOrigins, err
}

func (eo *ExpenseOrigin) FindByID(id string) (*entity.ExpenseOrigin, error) {
	var expenseOrigin entity.ExpenseOrigin
	err := eo.DB.First(&expenseOrigin, "id = ?", id).Error
	return &expenseOrigin, err
}

func (eo *ExpenseOrigin) Update(expenseOrigon *entity.ExpenseOrigin) error {
	_, err := eo.FindByID(expenseOrigon.ID.String())
	if err != nil {
		return err
	}
	return eo.DB.Save(expenseOrigon).Error
}

func (eo *ExpenseOrigin) Delete(id string) error {
	expenseOrigin, err := eo.FindByID(id)
	if err != nil {
		return err
	}
	return eo.DB.Delete(expenseOrigin).Error
}
