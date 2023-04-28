package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
	"github.com/Msaorc/ExpenseControlAPI/internal/entity"
	"github.com/Msaorc/ExpenseControlAPI/internal/infra/database"
	"github.com/Msaorc/ExpenseControlAPI/pkg/handler"
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
// @Success      201  {object}  dto.StatusMessage
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	var userInput dto.User
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	user, _ := u.UserDB.FindByEmail(userInput.Email)
	if user != nil {
		message := fmt.Sprintf("Já existe um usuário com o email (%s) cadastrado!", userInput.Email)
		handler.SetReturnStatusMessageHandlers(http.StatusAlreadyReported, message, w)
		return
	}
	user, err = entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	err = u.UserDB.Create(user)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	handler.SetReturnStatusMessageHandlers(http.StatusCreated, "User created successfully.", w)
}

// Authenticate User godoc
// @Summary      Authenticate User
// @Description  Authenticate User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.UserAuthenticate  true  "authenticate request"
// @Success      200  {object}  dto.UserAuthenticateOutput
// @Failure      404  {object}  dto.StatusMessage
// @Failure      500  {object}  dto.StatusMessage
// @Router       /users/authenticate [post]
func (u *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	handler.SetHeader(w, http.StatusOK)
	var userAuthenticate dto.UserAuthenticate
	err := json.NewDecoder(r.Body).Decode(&userAuthenticate)
	if err != nil {
		handler.SetReturnStatusMessageHandlers(http.StatusInternalServerError, err.Error(), w)
		return
	}
	user, err := u.UserDB.FindByEmail(userAuthenticate.Email)
	if user == nil {
		handler.SetReturnStatusMessageHandlers(http.StatusNotFound, err.Error(), w)
		return
	}
	if !user.ValidatePassword(userAuthenticate.Password) {
		handler.SetReturnStatusMessageHandlers(http.StatusUnauthorized, err.Error(), w)
		return
	}
	_, token, _ := u.Jwt.Encode(map[string]interface{}{
		"sub":  user.ID.String(),
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * time.Duration(u.JwtExperiesIn)).Unix(),
	})
	accessToken := dto.UserAuthenticateOutput{
		UserID:      user.ID.String(),
		AccessToken: token,
	}
	json.NewEncoder(w).Encode(accessToken)
}
