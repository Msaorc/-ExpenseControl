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

func (eo *ExpenseOrigin) FindAll(page, limit int, sort string) ([]entity.Expense, error) {
	var expenseOrigins []entity.Expense
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page >= 0 && limit >= 0 {
		err = eo.DB.Limit(limit).Offset((page - 1) * limit).Order("description " + sort).Find(expenseOrigins).Error
		return expenseOrigins, err
	}
	return nil, nil
}

func (eo *ExpenseOrigin) FindByID(id string) (*entity.ExpenseOrigin, error) {
	var expenseOrigin entity.ExpenseOrigin
	err := eo.DB.First(expenseOrigin, "id = ?", id).Error
	return &expenseOrigin, err
}

func (eo *ExpenseOrigin) Update(expenseOrigon entity.ExpenseOrigin) error {
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
