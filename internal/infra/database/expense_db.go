package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/gorm"
)

type Expense struct {
	DB *gorm.DB
}

func NewExpenseDB(db *gorm.DB) *Expense {
	return &Expense{DB: db}
}

func (e *Expense) Create(expense *entity.Expense) error {
	return e.DB.Create(expense).Error
}

func (e *Expense) FindAll(page, limit int, sort string) ([]entity.Expense, error) {
	var expense []entity.Expense
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page >= 0 && limit >= 0 {
		err = e.DB.Limit(limit).Offset((page - 1) * limit).Order("description " + sort).Find(&expense).Error
		return expense, err
	}
	err = e.DB.Order("description " + sort).Find(&expense).Error
	return expense, err
}

func (e *Expense) FindByID(id string) (*entity.Expense, error) {
	var expense *entity.Expense
	err := e.DB.Where("id = ?", id).First(&expense).Error
	if err != nil {
		return nil, err
	}
	return expense, err
}

func (e *Expense) Update(expense *entity.Expense) error {
	if _, err := e.FindByID(expense.ID.String()); err != nil {
		return err
	}
	return e.DB.Save(expense).Error
}

func (e *Expense) Delete(id string) error {
	expense, err := e.FindByID(id)
	if err != nil {
		return err
	}
	return e.DB.Delete(expense).Error
}
