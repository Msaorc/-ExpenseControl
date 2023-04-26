package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/Msaorc/ExpenseControlAPI/pkg/date"
	entityPKG "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/go-chi/chi"
)

type PeriodHandler struct {
	PeriodDB database.PeriodInterface
}

func NewPeriodHandler(db database.PeriodInterface) *PeriodHandler {
	return &PeriodHandler{PeriodDB: db}
}

func (ph *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var period dto.PeriodInput
	err := json.NewDecoder(r.Body).Decode(&period)
	if err != nil {
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
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = ph.PeriodDB.Create(periodEntity)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ph *PeriodHandler) FindAllPeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "applcation/json")
	w.WriteHeader(http.StatusOK)
	period, err := ph.PeriodDB.FindAll()
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	json.NewEncoder(w).Encode(period)
}

func (ph *PeriodHandler) FindPeriodByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "ID inválido",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	period, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID inválido",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	json.NewEncoder(w).Encode(period)
}

func (ph *PeriodHandler) UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID inválido",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var periodDto dto.PeriodInput
	err = json.NewDecoder(r.Body).Decode(&periodDto)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Erro para converter json em struct",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	uuidEntity, err := entityPKG.ParseID(id)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var period entity.Period
	period.ID = uuidEntity
	period.Description = periodDto.Description
	period.InitialDate, _ = time.Parse(date.DateLayout, periodDto.InitialDate)
	period.FinalDate, _ = time.Parse(date.DateLayout, periodDto.FinalDate)
	err = ph.PeriodDB.Update(&period)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
}

func (ph *PeriodHandler) DeletePeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID inválido",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = ph.PeriodDB.Delete(id)
	if err != nil {
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
}
