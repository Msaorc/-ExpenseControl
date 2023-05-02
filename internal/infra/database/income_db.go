package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/gorm"
)

type IncomeDB struct {
	DB *gorm.DB
}

func NewIncomeDB(db *gorm.DB) *IncomeDB {
	return &IncomeDB{DB: db}
}

func (i *IncomeDB) CreateIncome(income *entity.Income) error {
	return i.DB.Create(income).Error
}
