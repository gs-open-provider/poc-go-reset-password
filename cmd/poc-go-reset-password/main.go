package main

import (
	"log"
	"net/http"

	"github.com/gs-open-provider/poc-go-reset-password/internal/configs"
	"github.com/gs-open-provider/poc-go-reset-password/internal/logger"
	"github.com/gs-open-provider/poc-go-reset-password/internal/services"
	"github.com/spf13/viper"
)

func main() {
	configs.InitializeViper()
	logger.InitializeZapCustomLogger()

	http.HandleFunc("/index", services.HandleIndex)
	http.HandleFunc("/reset-password", services.ResetPassword)
	http.HandleFunc("/reset-password-callback", services.ResetPasswordCallback)

	logger.Log.Info("Starting server on http://localhost:" + viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":9090", nil))
}
