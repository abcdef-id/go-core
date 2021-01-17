package config

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

//JwtCustomClaims JwtCustomClaims
type JwtCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	jwt.StandardClaims
	IsAdmin bool `json:"is_admin"`
}

//JwtConfig JwtConfig
var JwtConfig middleware.JWTConfig

func init() {
	JwtConfig = middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(viper.GetString("jwtSign")),
	}
}
