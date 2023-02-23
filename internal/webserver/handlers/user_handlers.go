package handlers

import "github.com/Msaorc/ExpenseControlAPI/internal/infra/database"

type UserHandler struct {
	DB *database.UserInterface
}

