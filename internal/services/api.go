package services

import (
	"net/http"

	"github.com/gs-open-provider/poc-go-reset-password/internal/helpers"
	"github.com/gs-open-provider/poc-go-reset-password/internal/logger"
)

// HandleIndex Function
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Index called..")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(helpers.IndexPage))
}

// ResetPassword Function
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Reset Password called..")
	HandleResetPassword(w, r)
}

// ResetPasswordCallback Function
func ResetPasswordCallback(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Reset Password Callback called..")
	HandleResetPasswordCallback(w, r)
}
