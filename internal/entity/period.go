package entity

import (
	"errors"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrPeriodIdIsRequired = errors.New("Expense: ID is Required")
var ErrPeriodIdIsInvalid = errors.New("Expense: ID invalid")
var ErrPeriodDescriptionIsRequired = errors.New("Period: Description is Required")
var ErrPeriodInitialDateIsRequired = errors.New("Period: InitialDate is Required")
var ErrPeriodFinalDateIsRequired = errors.New("Period: FinalDate is Required")
var ErrPeriodInitalDateNotBefore = errors.New("Period: Start date cannot be greater than after all")
var ErrPeriodInitalDateNotEquals = errors.New("Period: The start date cannot be the same as the end date")

type Period struct {
	ID          entity.ID `gorm:"primaryKey" json:"id"`
	Description string    `json:"description"`
	InitalDate  time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}

func NewPeriod(description string, initialDate string, finalDate string) (*Period, error) {
	if initialDate == "" {
		return nil, ErrPeriodInitialDateIsRequired
	}
	if finalDate == "" {
		return nil, ErrPeriodFinalDateIsRequired
	}
	iDate, err := time.Parse(date.DateLayout, initialDate)
	if err != nil {
		return nil, err
	}
	fDate, err := time.Parse(date.DateLayout, finalDate)
	if err != nil {
		return nil, err
	}
	period := &Period{
		ID:          entity.NewID(),
		Description: description,
		InitalDate:  iDate,
		FinalDate:   fDate,
	}
	err = period.Validate()
	if err != nil {
		return nil, err
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
	if p.FinalDate.Before(p.InitalDate) {
		return ErrPeriodInitalDateNotBefore
	}
	if p.InitalDate.Equal(p.FinalDate) {
		return ErrPeriodInitalDateNotEquals
	}
	return nil
}
