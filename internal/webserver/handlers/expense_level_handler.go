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
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenselevel [post]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) CreateExpenseLevel(w http.ResponseWriter, r *http.Request) {
	var expenseLevel dto.ExpenseLevel
	err := json.NewDecoder(r.Body).Decode(&expenseLevel)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	ExpenseLevelEntity, err := entity.NewExpenseLevel(expenseLevel.Description, expenseLevel.Color)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = el.ExpenseLevelDB.Create(ExpenseLevelEntity)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// FindAll ExpenseLevel godoc
// @Summary      FindAll ExpenseLevel
// @Description  FindAll ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.ExpenseLevel
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenselevel [get]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) FindAllExpenseLevel(w http.ResponseWriter, r *http.Request) {
	expensesLevel, err := el.ExpenseLevelDB.FindAll()
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenselevel/{id} [get]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) FindExpenseLevelById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: "Id invalido!"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expenseOrigin, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenselevel/{id} [put]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) UpdateExpenseLevel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: "Id invalido!"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var expenseLevel entity.ExpenseLevel
	err = json.NewDecoder(r.Body).Decode(&expenseLevel)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expenseLevel.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = el.ExpenseLevelDB.Update(&expenseLevel)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete ExpenseLevel godoc
// @Summary      Delete ExpenseLevel
// @Description  Delete ExpenseLevel
// @Tags         ExpenseLevel
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "ExpenseLevel ID" Format(uuid)
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expenselevel/{id} [delete]
// @Security ApiKeyAuth
func (el *ExpenseLevelHandler) DeleteExpenseLevel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: "Id invalido!"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := el.ExpenseLevelDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = el.ExpenseLevelDB.Delete(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusOK)
}
