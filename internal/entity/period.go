package entity

import (
	"errors"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrPeriodIdIsRequired = errors.New("Expense: ID is Required")
var ErrPeriodIdIsInvalid = errors.New("Expense: ID invalid")
var ErrPeriodDescriptionIsRequired = errors.New("Period: Description is Required")
var ErrPeriodInitialDateIsRequired = errors.New("Period: InitialDate is Required")
var ErrPeriodFinalDateIsRequired = errors.New("Period: FinalDate is Required")
var ErrPeriodInitalDateNotBefore = errors.New("Period: Start date cannot be greater than after all")
var ErrPeriodInitalDateNotEqualsFinalDate = errors.New("Period: The start date cannot be the same as the end date")

type Period struct {
	ID          entity.ID `gorm:"primaryKey" json:"id"`
	Description string    `json:"description"`
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}

func NewPeriod(description string, initialDate string, finalDate string) (*Period, error) {
	if initialDate == "" {
		return nil, ErrPeriodInitialDateIsRequired
	}
	if finalDate == "" {
		return nil, ErrPeriodFinalDateIsRequired
	}
	period := &Period{
		ID:          entity.NewID(),
		Description: description,
		InitialDate: date.ConvertDateToTime(initialDate),
		FinalDate:   date.ConvertDateToTime(finalDate),
	}
	err := period.Validate()
	if err != nil {
		return nil, err
	}
	return period, nil
}

func UpdatePeriod(id entity.ID, periodDto dto.PeriodInput) (*Period, error) {
	period := &Period{
		ID:          id,
		Description: periodDto.Description,
		InitialDate: date.ConvertDateToTime(periodDto.InitialDate),
		FinalDate:   date.ConvertDateToTime(periodDto.FinalDate),
	}
	err := period.Validate()
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
	if p.InitialDate.String() == "" {
		return ErrPeriodInitialDateIsRequired
	}
	if p.FinalDate.String() == "" {
		return ErrPeriodFinalDateIsRequired
	}
	if p.FinalDate.Before(p.InitialDate) {
		return ErrPeriodInitalDateNotBefore
	}
	if p.InitialDate.Equal(p.FinalDate) {
		return ErrPeriodInitalDateNotEqualsFinalDate
	}
	return nil
}
