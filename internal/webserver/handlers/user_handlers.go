package handlers

import (
	"encoding/json"
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

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.User
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = u.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var userAuthenticate dto.UserAuthenticate
	err := json.NewDecoder(r.Body).Decode(&userAuthenticate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := u.UserDB.FindByEmail(userAuthenticate.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !user.ValidatePassword(userAuthenticate.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := u.Jwt.Encode(map[string]interface{}{
		"sub":  user.ID.String(),
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * time.Duration(u.JwtExperiesIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}
