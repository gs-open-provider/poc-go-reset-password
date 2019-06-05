package main

import (
	"github.com/gs-open-provider/poc-go-reset-password/internal/configs"
	"github.com/gs-open-provider/poc-go-reset-password/internal/logger"
)

func main() {
	configs.InitializeViper()
	logger.InitializeZapCustomLogger()

}
