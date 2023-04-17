package entity

import (
	"errors"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrPeriodIdIsRequired = errors.New("Expense: ID is Required")
var ErrPeriodIdIsInvalid = errors.New("Expense: ID invalid")
var ErrPeriodDescriptionIsRequired = errors.New("Period: Description is Required")
var ErrPeriodInitialDateIsRequired = errors.New("Period: InitialDate is Required")
var ErrPeriodFinalDateIsRequired = errors.New("Period: FinalDate is Required")

type Period struct {
	ID          entity.ID `gorm:"primaryKey" json:"id"`
	Description string    `json:"description"`
	InitalDate  time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}

func NewPeriod(description string, initialDate string, finalDate string) (*Period, error) {
	layoutDate := "2006-01-02"
	iDate, err := time.Parse(layoutDate, initialDate)
	if err != nil {
		return nil, err
	}
	fDate, err := time.Parse(layoutDate, initialDate)
	if err != nil {
		return nil, err
	}
	period := &Period{
		ID:          entity.NewID(),
		Description: description,
		InitalDate:  iDate,
		FinalDate:   fDate,
	}
	return period, nil
}

func (p *Period) Validate() error {
	if p.ID.String() == "" {
		return ErrPeriodIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrPeriodIdIsInvalid
	}
	if p.Description == "" {
		return ErrPeriodDescriptionIsRequired
	}
	if p.InitalDate.String() == "" {
		return ErrPeriodInitialDateIsRequired
	}
	if p.FinalDate.String() == "" {
		return ErrPeriodFinalDateIsRequired
	}
	return nil
}
