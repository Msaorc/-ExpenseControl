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

func (i *IncomeDB) FindAll() ([]entity.Income, error) {
	var income []entity.Income
	if err := i.DB.Find(&income).error; err != nil {
		return nil, err
	}
	return income, nil
}

func (i *IncomeDB) FindByID(id string) (*entity.Income, error){
	var income *entity.Income
	if err := i.DB.First(&income, "id = ?", id).error; err != nil {
		return nil, err
	}
	return income, nil
}
