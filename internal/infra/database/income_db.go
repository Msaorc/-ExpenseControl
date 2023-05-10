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

func (i *IncomeDB) Create(income *entity.Income) error {
	return i.DB.Create(income).Error
}

func (i *IncomeDB) FindAll() ([]entity.Income, error) {
	var income []entity.Income
	if err := i.DB.Find(&income).Error; err != nil {
		return nil, err
	}
	return income, nil
}

func (i *IncomeDB) FindByID(id string) (*entity.Income, error) {
	var income *entity.Income
	if err := i.DB.First(&income, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return income, nil
}

func (i *IncomeDB) Update(income *entity.Income) error {
	_, err := i.FindByID(income.ID.String())
	if err != nil {
		return err
	}
	return i.DB.Save(income).Error
}

func (i *IncomeDB) Delete(id string) error {
	income, err := i.FindByID(id)
	if err != nil {
		return err
	}
	return i.DB.Delete(income).Error
}
