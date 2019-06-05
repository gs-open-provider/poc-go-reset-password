package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gs-open-provider/poc-go-jwt/models"
	"github.com/gs-open-provider/poc-go-reset-password/internal/logger"
	"github.com/spf13/viper"
)

// JWTKey from Config/environment
// Create the JWT key used to create the signature
var JWTKey = []byte(viper.GetString("jwt-secret"))

// RegisterRequestParams Struct
type RegisterRequestParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Claims Struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateJWT Function
func CreateJWT(w http.ResponseWriter, r *http.Request, params *RegisterRequestParams) (string, error) {
	logger.Log.Info(params.Username)

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &models.Claims{
		Username: params.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		logger.Log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}

	logger.Log.Info(tokenString)
	return tokenString, nil
}
