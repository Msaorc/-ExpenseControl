package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	entityPKG "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/go-chi/chi"
)

type ExpenseOriginlHandler struct {
	ExpenseOriginDB database.ExpenseOriginInterface
}

func NewExpenseOriginHandler(db database.ExpenseOriginInterface) *ExpenseOriginlHandler {
	return &ExpenseOriginlHandler{
		ExpenseOriginDB: db,
	}
}

// Create ExpenseOrigin godoc
// @Summary      Create ExpenseOrigin
// @Description  Create ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Param        request   body      dto.ExepnseOrigin  true  "expenseorigin request"
// @Success      201
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenseorigin [post]
// @Security ApiKeyAuth
func (eoh *ExpenseOriginlHandler) CreateExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	var expenseOrigin dto.ExepnseOrigin
	err := json.NewDecoder(r.Body).Decode(&expenseOrigin)
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
	expenseOriginEntity, err := entity.NewExpenseOrigin(expenseOrigin.Description)
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
	err = eoh.ExpenseOriginDB.Create(expenseOriginEntity)
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
	w.WriteHeader(http.StatusCreated)
}

// FindAll ExpenseOrigin godoc
// @Summary      FindAll ExpenseOrigin
// @Description  FindAll ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ExpenseOrigin
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenseorigin [get]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) FindAllExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	expensesOrigin, err := eo.ExpenseOriginDB.FindAll()
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expensesOrigin)
}

// FindById ExpenseOrigin godoc
// @Summary      FindById ExpenseOrigin
// @Description  FindById ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseOrigin ID" Format(uuid)
// @Success      200  {object}  entity.ExpenseOrigin
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenseorigin/{id} [get]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) FindExpenseOriginById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID Inválido.",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expenseOrigin, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenseOrigin)
}

// Update ExpenseOrigin godoc
// @Summary      Update ExpenseOrigin
// @Description  Update ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseOrigin ID" Format(uuid)
// @Param        request   body      dto.ExepnseOrigin  true  "ExpenseOrigin request"
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenseorigin/{id} [put]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) UpdateExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID Inválido.",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var expenseOrigin entity.ExpenseOrigin
	err = json.NewDecoder(r.Body).Decode(&expenseOrigin)
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
	expenseOrigin.ID, err = entityPKG.ParseID(id)
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
	err = eo.ExpenseOriginDB.Update(&expenseOrigin)
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
	w.WriteHeader(http.StatusOK)
}

// Delete ExpenseOrigin godoc
// @Summary      Delete ExpenseOrigin
// @Description  Delete ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseOrigin ID" Format(uuid)
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenseorigin/{id} [delete]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) DeleteExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: "ID Inválido.",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = eo.ExpenseOriginDB.Delete(id)
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
	w.WriteHeader(http.StatusOK)
}
