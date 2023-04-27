package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ExpenseControlAPI/internal/dto"
)

func SetReturnStatusMessageHandlers(code int, message string, w http.ResponseWriter) {
	statusMessage := dto.StatusMessage{
		Code:    code,
		Message: message,
	}
	json.NewEncoder(w).Encode(statusMessage)
}

func SetHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("content-type", "applcation/json")
	w.WriteHeader(statusCode)
}
