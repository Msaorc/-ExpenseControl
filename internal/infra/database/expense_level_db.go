package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/gorm"
)

type ExpenseLevel struct {
	DB *gorm.DB
}

func NewExpenseLevelDB(db *gorm.DB) *ExpenseLevel {
	return &ExpenseLevel{DB: db}
}

func (el *ExpenseLevel) Create(expenseLevel *entity.ExpenseLevel) error {
	return el.DB.Create(expenseLevel).Error
}

func (el *ExpenseLevel) FindAll() ([]entity.ExpenseLevel, error) {
	var expenseLevel []entity.ExpenseLevel
	if err := el.DB.Find(&expenseLevel).Error; err != nil {
		return nil, err
	}
	return expenseLevel, nil
}

func (el *ExpenseLevel) FindByID(id string) (*entity.ExpenseLevel, error) {
	var expenseLevel *entity.ExpenseLevel
	err := el.DB.First(&expenseLevel, "id = ?", id).Error
	return expenseLevel, err
}

func (el *ExpenseLevel) Update(expenseLevel *entity.ExpenseLevel) error {
	_, err := el.FindByID(expenseLevel.ID.String())
	if err != nil {
		return err
	}
	return el.DB.Save(expenseLevel).Error
}

func (el *ExpenseLevel) Delete(id string) error {
	expenseLevel, err := el.FindByID(id)
	if err != nil {
		return err
	}
	return el.DB.Delete(expenseLevel).Error
}
