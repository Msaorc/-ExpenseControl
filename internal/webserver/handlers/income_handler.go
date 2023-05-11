package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	entityID "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/Msaorc/ExpenseControlAPI/pkg/handler"
	"github.com/go-chi/chi"
)

type IncomeHandler struct {
	IncomeDB database.IncomeInterface
}

func NewINcomeHandler(db database.IncomeInterface) *IncomeHandler {
	return &IncomeHandler{
		IncomeDB: db,
	}
}

// Create Income godoc
// @Summary      Create Income
// @Description  Create Income
// @Tags         Income
// @Accept       json
// @Produce      json
// @Param        request   body      dto.IncomeInput  true  "income request"
// @Success      201  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /income [post]
// @Security ApiKeyAuth
func (i *IncomeHandler) CreateIncome(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	var income dto.IncomeInput
	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	incomeEntity, err := entity.NewIncome(income.Description, income.Value, income.Date)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = i.IncomeDB.Create(incomeEntity)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "Income created successfully.", w)
}

// FindAll Income godoc
// @Summary      FindAll Income
// @Description  FindAll Income
// @Tags         Income
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.IncomeOutput
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /income [get]
// @Security ApiKeyAuth
func (i *IncomeHandler) FindAllIncome(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	incomes, err := i.IncomeDB.FindAll()
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(incomes)
}

// FindById Income godoc
// @Summary      FindById Income
// @Description  FindById Income
// @Tags         Income
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Income ID" Format(uuid)
// @Success      200  {object}  dto.IncomeOutput
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /income/{id} [get]
// @Security ApiKeyAuth
func (i *IncomeHandler) FindIncomeById(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	income, err := i.IncomeDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(income)
}

// Update Income godoc
// @Summary      Update Income
// @Description  Update Income
// @Tags         Income
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Income ID" Format(uuid)
// @Param        request   body      dto.IncomeInput  true  "expense request"
// @Success      200
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /income/{id} [put]
// @Security ApiKeyAuth
func (i *IncomeHandler) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := i.IncomeDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	var income dto.IncomeInput
	err = json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	incomeID, err := entityID.ParseID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	incomeEntity, err := entity.UpdateIncome(income, incomeID)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = i.IncomeDB.Update(incomeEntity)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Income updated successfully.", w)
}

// Delete Income godoc
// @Summary      Delete Income
// @Description  Delete Income
// @Tags         Income
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Income ID" Format(uuid)
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /income/{id} [delete]
// @Security ApiKeyAuth
func (i *IncomeHandler) DeleteIncome(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	err := i.IncomeDB.Delete(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Successfully deleted Income.", w)
}
