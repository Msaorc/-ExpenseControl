package database

import (
	"errors"

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
	expenseOrigin := NewExpenseOriginDB(e.DB)
	_, err := expenseOrigin.FindByID(expense.ExpenseOriginID)
	if err != nil {
		return errors.New("Origem despesa inválida, verifique!")
	}
	expenseLevel := NewExpenseLevelDB(e.DB)
	_, err = expenseLevel.FindByID(expense.ExpenseLevelID)
	if err != nil {
		return errors.New("Nivel despesa inválida, verifique!")
	}
	return e.DB.Create(expense).Error
}

func (e *Expense) FindAll(page, limit int, sort string) ([]entity.Expense, error) {
	var expense []entity.Expense
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page >= 0 && limit >= 0 {
		err = e.DB.Preload("ExpenseLevel").Preload("ExpenseOrigin").Limit(limit).Offset((page - 1) * limit).Order("description " + sort).Find(&expense).Error
		return expense, err
	}
	err = e.DB.Preload("ExpenseLevel").Preload("ExpenseOrigin").Find(&expense).Order("description " + sort).Error
	return expense, err
}

func (e *Expense) FindByID(id string) (*entity.Expense, error) {
	var expense *entity.Expense
	err := e.DB.First(&expense, "id = ?", id).Error
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
