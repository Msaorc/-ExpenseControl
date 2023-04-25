package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
)

type PeriodHandler struct {
	PeriodOriginDB database.PeriodInterface
}

func NewPeriodDB(db database.PeriodInterface) *PeriodHandler {
	return &PeriodHandler{PeriodOriginDB: db}
}

func (ph *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	var period dto.PeriodInput
	err := json.NewDecoder(r.Body).Decode(&period)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	periodEntity, err := entity.NewPeriod(period.Description, period.InitialDate, period.FinalDate)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = ph.PeriodOriginDB.Create(periodEntity)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
