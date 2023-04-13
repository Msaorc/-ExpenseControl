package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	entityPKG "github.com/Msaorc/ExpenseControlAPI/pkg/entity"
	"github.com/go-chi/chi"
)

type ExpenseHandler struct {
	ExpenseDB database.ExpenseInterface
}

func NewExpenseHandler(db database.ExpenseInterface) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseDB: db,
	}
}

// Create Expense godoc
// @Summary      Create Expense
// @Description  Create Expense
// @Tags         Expense
// @Accept       json
// @Produce      json
// @Param        request   body      dto.Expense  true  "expense request"
// @Success      201
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expense [post]
// @Security ApiKeyAuth
func (e *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense dto.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expenseEntity, err := entity.NewExpense(expense.Description, expense.Value, expense.LevelID, expense.OringID, expense.Note)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = e.ExpenseDB.Create(expenseEntity)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// FindAll godoc
// @Summary      FindAll Expense
// @Description  FindAll Expense
// @Tags         Expense
// @Accept       json
// @Produce      json
// @Param        page    query    string   false  "page number"
// @Param        limit   query    string   false   "limit"
// @Success      200     {array}  entity.Expense
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expense [get]
// @Security ApiKeyAuth
func (e *ExpenseHandler) FindAllExpense(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageint, err := strconv.Atoi(page)
	if err != nil {
		pageint = 0
	}
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		limitint = 0
	}

	expenses, err := e.ExpenseDB.FindAll(pageint, limitint, sort)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var expensesOutput []dto.ExpenseAll
	for _, expense := range expenses {
		expenseOutput := dto.ExpenseAll{
			ID:                expense.ID.String(),
			Description:       expense.Description,
			Value:             expense.Value,
			LevelDescription:  expense.ExpenseLevel.Description,
			OringDescritption: expense.ExpenseOrigin.Description,
			Note:              expense.Note,
		}
		expensesOutput = append(expensesOutput, expenseOutput)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expensesOutput)
}

// FindById Expense godoc
// @Summary      FindById Expense
// @Description  FindById Expense
// @Tags         Expense
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Expense ID" Format(uuid)
// @Success      200  {object}  entity.Expense
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expense/{id} [get]
// @Security ApiKeyAuth
func (e *ExpenseHandler) FindExpenseById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: "Id inválido"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expense, err := e.ExpenseDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

// Update Expense godoc
// @Summary      Update Expense
// @Description  Update Expense
// @Tags         Expense
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Expense ID" Format(uuid)
// @Param        request   body      dto.Expense  true  "expense request"
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expense/{id} [put]
// @Security ApiKeyAuth
func (e *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: "Id inválido"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	_, err := e.ExpenseDB.FindByID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	var expense entity.Expense
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	expense.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = e.ExpenseDB.Update(&expense)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Expense godoc
// @Summary      Delete Expense
// @Description  Delete Expense
// @Tags         Expense
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Expense ID" Format(uuid)
// @Success      200
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /expense/{id} [delete]
// @Security ApiKeyAuth
func (e *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := e.ExpenseDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
