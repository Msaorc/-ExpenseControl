package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        db,
		Jwt:           jwt,
		JwtExperiesIn: jwtExperiesIn,
	}
}

// Create User godoc
// @Summary      Create User
// @Description  Create User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.User  true  "user request"
// @Success      201
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.User
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusNotFound,
			Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	user, _ := u.UserDB.FindByEmail(userInput.Email)
	if user != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("Já existe um usuário com o email (%s) cadastrado!", userInput.Email)
		errorMessage := dto.Error{
			Code:    http.StatusAlreadyReported,
			Message: message}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	user, err = entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = u.UserDB.Create(user)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		errorMessage := dto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Authenticate User godoc
// @Summary      Authenticate User
// @Description  Authenticate User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.UserAuthenticate  true  "authenticate request"
// @Success      200  {object}  dto.UserAuthenticateOutput
// @Failure      404  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /users/authenticate [post]
func (u *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var userAuthenticate dto.UserAuthenticate
	err := json.NewDecoder(r.Body).Decode(&userAuthenticate)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	user, err := u.UserDB.FindByEmail(userAuthenticate.Email)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	if !user.ValidatePassword(userAuthenticate.Password) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	_, token, _ := u.Jwt.Encode(map[string]interface{}{
		"sub":  user.ID.String(),
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * time.Duration(u.JwtExperiesIn)).Unix(),
	})

	fmt.Println("Passou em td, vamos devolver o token jwt")
	accessToken := dto.UserAuthenticateOutput{AccessToken: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
