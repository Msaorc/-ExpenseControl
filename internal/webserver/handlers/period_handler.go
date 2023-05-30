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

// Create Period godoc
// @Summary      Create Period
// @Description  Create Period
// @Tags         Period
// @Accept       json
// @Produce      json
// @Param        request   body      dto.PeriodInput  true  "period request"
// @Success      201  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /period [post]
// @Security ApiKeyAuth
func (ph *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	var period dto.PeriodInput
	err := json.NewDecoder(r.Body).Decode(&period)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	periodEntity, err := entity.NewPeriod(period.Description, period.InitialDate, period.FinalDate)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = ph.PeriodDB.Create(periodEntity)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusCreated)
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "Period created successfully.", w)
}

// FindAll Period godoc
// @Summary      FindAll Period
// @Description  FindAll Period
// @Tags         Period
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.Period
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /period [get]
// @Security ApiKeyAuth
func (ph *PeriodHandler) FindAllPeriod(w http.ResponseWriter, r *http.Request) {
	period, err := ph.PeriodDB.FindAll()
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(period)
}

// FindById Period godoc
// @Summary      FindById Period
// @Description  FindById Period
// @Tags         Period
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Period ID" Format(uuid)
// @Success      200  {object}  entity.Period
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /period/{id} [get]
// @Security ApiKeyAuth
func (ph *PeriodHandler) FindPeriodByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, "invalid ID", w)
		return
	}
	period, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(period)
}

// Update Period godoc
// @Summary      Update Period
// @Description  Update Period
// @Tags         Period
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Period ID" Format(uuid)
// @Param        request   body      dto.PeriodInput  true  "Period request"
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /period/{id} [put]
// @Security ApiKeyAuth
func (ph *PeriodHandler) UpdatePeriod(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	var periodDto dto.PeriodInput
	err = json.NewDecoder(r.Body).Decode(&periodDto)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	uuidEntity, err := entityPKG.ParseID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	period, err := entity.UpdatePeriod(uuidEntity, periodDto)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = ph.PeriodDB.Update(period)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Period updated successfully.", w)
}

// Delete Period godoc
// @Summary      Delete Period
// @Description  Delete Period
// @Tags         Period
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Period ID" Format(uuid)
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /period/{id} [delete]
// @Security ApiKeyAuth
func (ph *PeriodHandler) DeletePeriod(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := ph.PeriodDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	err = ph.PeriodDB.Delete(id)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Successfully deleted Period.", w)
}
