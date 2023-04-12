package entity

import (
	"errors"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/pkg/entity"
)

var ErrExpenseIdIsRequired = errors.New("Expense: ID is Required")
var ErrExpenseIdIsInvalid = errors.New("Expense: ID invalid")
var ErrExpenseDescriptionIsRequired = errors.New("Expense: Description is Required")
var ErrExpenseValueIsRequired = errors.New("Expense: Value is Required")
var ErrExpenseValueIsInvalid = errors.New("Expense: Value is Invalid")
var ErrExpenseLevelIsRequired = errors.New("Expense: LevelID is Required")
var ErrExpenseOriginIsRequired = errors.New("Expense: OriginID is Required")

type Expense struct {
	ID              entity.ID `gorm:"primaryKey" json:"id"`
	Description     string    `json:"description"`
	Value           float64   `json:"value"`
	Date            time.Time `json:"date"`
	ExpenseLevelID  string    `gorm:"size:36" json:"level_id"`
	ExpenseLevel    ExpenseLevel
	ExpenseOriginID string `gorm:"size:36" json:"origin_id"`
	ExpenseOrigin   ExpenseOrigin
	Note            string `json:"note"`
}

func NewExpense(description string, value float64, levelID string, originID string, note string) (*Expense, error) {
	expense := &Expense{
		ID:              entity.NewID(),
		Description:     description,
		Value:           value,
		Date:            time.Now(),
		ExpenseLevelID:  levelID,
		ExpenseOriginID: originID,
		Note:            note,
	}
	err := expense.Validate()
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (e *Expense) Validate() error {
	if e.ID.String() == "" {
		return ErrExpenseIdIsRequired
	}
	if _, err := entity.ParseID(e.ID.String()); err != nil {
		return ErrExpenseIdIsInvalid
	}
	if e.Description == "" {
		return ErrExpenseDescriptionIsRequired
	}
	if e.Value == 0 {
		return ErrExpenseValueIsRequired
	}
	if e.Value < 0 {
		return ErrExpenseValueIsInvalid
	}
	if e.ExpenseLevelID == "" {
		return ErrExpenseLevelIsRequired
	}
	if e.ExpenseOriginID == "" {
		return ErrExpenseOriginIsRequired
	}
	return nil
}
