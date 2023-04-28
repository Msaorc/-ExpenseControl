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

type ExpenseLevelHandler struct {
	ExpenseLevelDB database.ExpenseLevelInterface
}

func NewExpenseLevelHandler(db database.ExpenseLevelInterface) *ExpenseLevelHandler {
	return &ExpenseLevelHandler{
		ExpenseLevelDB: db,
	}
}

// Create ExpenseLevel godoc
// @Summary      Create ExpenseLevel
// @Description  Create ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Param        request   body      dto.ExpenseLevel  true  "expenselevel request"
// @Success      201
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenselevel [post]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) CreateExpenseLevel(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	var expenseLevel dto.ExpenseLevel
	err := json.NewDecoder(r.Body).Decode(&expenseLevel)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusOK, err.Error(), w)
	}
	ExpenseLevelEntity, err := entity.NewExpenseLevel(expenseLevel.Description, expenseLevel.Color)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusOK, err.Error(), w)
		return
	}
	err = el.ExpenseLevelDB.Create(ExpenseLevelEntity)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusOK, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "ExpenseLevel created successfully.", w)
}

// FindAll ExpenseLevel godoc
// @Summary      FindAll ExpenseLevel
// @Description  FindAll ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ExpenseLevel
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenselevel [get]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) FindAllExpenseLevel(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	expensesLevel, err := el.ExpenseLevelDB.FindAll()
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(expensesLevel)
}

// FindById ExpenseLevel godoc
// @Summary      FindById ExpenseLevel
// @Description  FindById ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseLevel ID" Format(uuid)
// @Success      200  {object}  entity.ExpenseLevel
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenselevel/{id} [get]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) FindExpenseLevelById(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	expenseOrigin, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(expenseOrigin)
}

// Update ExpenseLevel godoc
// @Summary      Update ExpenseLevel
// @Description  Update ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseLevel ID" Format(uuid)
// @Param        request   body      dto.ExpenseLevel  true  "expenselevel request"
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenselevel/{id} [put]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) UpdateExpenseLevel(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	var expenseLevel entity.ExpenseLevel
	err = json.NewDecoder(r.Body).Decode(&expenseLevel)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	expenseLevel.ID, err = entityPKG.ParseID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = el.ExpenseLevelDB.Update(&expenseLevel)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "ExpenseLevel updated successfully.", w)
}

// Delete ExpenseLevel godoc
// @Summary      Delete ExpenseLevel
// @Description  Delete ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseLevel ID" Format(uuid)
// @Success      200  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /expenselevel/{id} [delete]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) DeleteExpenseLevel(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	id := chi.URLParam(r, "id")
	if id == "" {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, "invalid ID", w)
		return
	}
	_, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	err = el.ExpenseLevelDB.Delete(id)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusOK, "Successfully deleted ExpenseLevel.", w)
}
