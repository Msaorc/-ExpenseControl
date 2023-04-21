package database

import (
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"gorm.io/gorm"
)

type Period struct {
	DB *gorm.DB
}

func NewPeriodDB(db *gorm.DB) *Period {
	return &Period{DB: db}
}

func (p *Period) Create(period *entity.Period) error {
	return p.DB.Create(&period).Error
}
