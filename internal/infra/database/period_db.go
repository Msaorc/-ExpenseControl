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

func (p *Period) FindAll() ([]entity.Period, error) {
	var period []entity.Period
	if err := p.DB.Find(&period).Order("description asc").Error; err != nil {
		return nil, err
	}
	return period, nil
}

func (p *Period) FindByID(id string) (*entity.Period, error) {
	var period *entity.Period
	err := p.DB.First(&period, "id = ?", id).Error
	return period, err
}

func (p *Period) Update(period *entity.Period) error {
	_, err := p.FindByID(period.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(&period).Error
}

func (p *Period) Delete(id string) error {
	period, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(period).Error
}
