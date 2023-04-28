package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	entityPKG "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/Msaorc/ExpenseControlAPI/pkg/handler"
	"github.com/go-chi/chi"
)

type PeriodHandler struct {
	PeriodDB database.PeriodInterface
}

func NewPeriodHandler(db database.PeriodInterface) *PeriodHandler {
	return &PeriodHandler{PeriodDB: db}
}

func (ph *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	var period dto.PeriodInput
	err := json.NewDecoder(r.Body).Decode(&period)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	periodEntity, err := entity.NewPeriod(period.Description, period.InitialDate, period.FinalDate)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = ph.PeriodDB.Create(periodEntity)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "Period created successfully.", w)
}

func (ph *PeriodHandler) FindAllPeriod(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	period, err := ph.PeriodDB.FindAll()
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(period)
}

func (ph *PeriodHandler) FindPeriodByID(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, "invalid ID", w)
		return
	}
	period, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(period)
}

func (ph *PeriodHandler) UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	var periodDto dto.PeriodInput
	err = json.NewDecoder(r.Body).Decode(&periodDto)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	uuidEntity, err := entityPKG.ParseID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	period := entity.UpdatePeriod(uuidEntity, periodDto)
	err = ph.PeriodDB.Update(period)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Period updated successfully.", w)
}

func (ph *PeriodHandler) DeletePeriod(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	err = ph.PeriodDB.Delete(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Successfully deleted Period.", w)
}
