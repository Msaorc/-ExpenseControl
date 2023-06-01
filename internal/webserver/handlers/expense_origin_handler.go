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
// @Success      201  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenseorigin [post]
// @Security ApiKeyAuth
func (eoh *ExpenseOriginlHandler) CreateExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	var expenseOrigin dto.ExepnseOrigin
	err := json.NewDecoder(r.Body).Decode(&expenseOrigin)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	expenseOriginEntity, err := entity.NewExpenseOrigin(expenseOrigin.Description)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = eoh.ExpenseOriginDB.Create(expenseOriginEntity)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusCreated)
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "ExpenseLevel created successfully.", w)
}

// FindAll ExpenseOrigin godoc
// @Summary      FindAll ExpenseOrigin
// @Description  FindAll ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ExpenseOrigin
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenseorigin [get]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) FindAllExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	expensesOrigin, err := eo.ExpenseOriginDB.FindAll()
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
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
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenseorigin/{id} [get]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) FindExpenseOriginById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	expenseOrigin, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
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
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenseorigin/{id} [put]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) UpdateExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	var expenseOrigin entity.ExpenseOrigin
	err = json.NewDecoder(r.Body).Decode(&expenseOrigin)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	expenseOrigin.ID, err = entityPKG.ParseID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = eo.ExpenseOriginDB.Update(&expenseOrigin)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "ExpenseOrigin updated successfully.", w)
}

// Delete ExpenseOrigin godoc
// @Summary      Delete ExpenseOrigin
// @Description  Delete ExpenseOrigin
// @Tags         ExpensesOrigin
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseOrigin ID" Format(uuid)
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenseorigin/{id} [delete]
// @Security ApiKeyAuth
func (eo *ExpenseOriginlHandler) DeleteExpenseOrigin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := eo.ExpenseOriginDB.FindByID(id)
	if err != nil {
		handler.SetHeader(w, http.StatusNotFound)
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	err = eo.ExpenseOriginDB.Delete(id)
	if err != nil {
		handler.SetHeader(w, http.StatusInternalServerError)
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetHeader(w, http.StatusOK)
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Successfully deleted ExpenseOrigin.", w)
}
