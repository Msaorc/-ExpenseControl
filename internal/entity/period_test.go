package entity

import (
	"testing"

	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	period, err := NewPeriod("Maio", "2023-04-11", "2023-05-10")
	assert.NotEmpty(t, period)
	assert.Empty(t, err)
	assert.NotEmpty(t, period.ID.String())
	assert.Equal(t, period.Description, "Maio")
	assert.Equal(t, period.InitialDate.String(), "2023-04-11 00:00:00 +0000 UTC")
	assert.Equal(t, period.FinalDate.String(), "2023-05-10 00:00:00 +0000 UTC")
	assert.Equal(t, date.ConvertDateToString(period.InitialDate), "2023-04-11")
	assert.Equal(t, date.ConvertDateToString(period.FinalDate), "2023-05-10")
}

func TestPeriodWhenDescriptionIsRequired(t *testing.T) {
	period, err := NewPeriod("", "2023-04-11", "2023-05-10")
	assert.Nil(t, period)
	assert.Error(t, err)
	assert.Equal(t, ErrPeriodDescriptionIsRequired, err)
}

func TestPeriodWhenInitialDateIsRequired(t *testing.T) {
	period, err := NewPeriod("Maio", "", "2023-05-10")
	assert.Nil(t, period)
	assert.Error(t, err)
	assert.Equal(t, ErrPeriodInitialDateIsRequired, err)
}

func TestPeriodWhenFinalDateIsRequired(t *testing.T) {
	period, err := NewPeriod("Maio", "2023-04-11", "")
	assert.Nil(t, period)
	assert.Error(t, err)
	assert.Equal(t, ErrPeriodFinalDateIsRequired, err)
}

func TestPeriodWhenInitalDateIsBefore(t *testing.T) {
	period, err := NewPeriod("Maio", "2023-04-11", "2023-04-09")
	assert.Nil(t, period)
	assert.Error(t, err)
	assert.Equal(t, ErrPeriodInitalDateNotBefore, err)
}

func TestPeriodWhenInitalDateEqualsFinalDate(t *testing.T) {
	period, err := NewPeriod("Maio", "2023-04-11", "2023-04-11")
	assert.Nil(t, period)
	assert.Error(t, err)
	assert.Equal(t, ErrPeriodInitalDateNotEqualsFinalDate, err)
}
