package services

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gs-open-provider/poc-go-reset-password/internal/helpers"
	"github.com/gs-open-provider/poc-go-reset-password/internal/logger"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

// HandleResetPassword Function
func HandleResetPassword(w http.ResponseWriter, r *http.Request) {
	var params RegisterRequestParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		logger.Log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Log.Info(params.Email)

	to := mail.NewEmail("Example User", params.Email)
	jwtToken, err := CreateJWT(w, r, &params)
	if err != nil {
		logger.Log.Error(err.Error())
		w.Write([]byte("Password Reset Failed.."))
		return
	}
	logger.Log.Info(jwtToken)
	message := mail.NewSingleEmail(
		helpers.RPFromAddress,
		helpers.RPEmailSubject,
		to,
		helpers.RPEmailContent,
		helpers.GetResetPasswordTemplate("http://192.168.60.11:9090/reset-password-callback?token="+jwtToken),
	)
	client := sendgrid.NewSendClient(viper.GetString("sendgrid-apikey"))
	_, err = client.Send(message)
	if err != nil {
		logger.Log.Info(err.Error())
		w.Write([]byte("Password Reset Failed.."))
		return
	}
	logger.Log.Info("Email Sent..")
	http.Redirect(w, r, "/email-sent", http.StatusTemporaryRedirect)
}

// HandleEmailSent Function
func HandleEmailSent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Password Reset Email Sent.."))
}

// HandleResetPasswordCallback Function
func HandleResetPasswordCallback(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(r.URL.Query().Get("token"))
	tknStr := r.URL.Query().Get("token")

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Password Reset Failed.."))
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Password Reset Failed.."))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password Reset Failed.."))
		return
	}

	w.Write([]byte("Reset your password here.. ==>> " + claims.Username))
}
