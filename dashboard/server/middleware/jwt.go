package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/yijia-cc/grouplive/dashboard/config"
)

func StartupJWT(cfg *config.Config) *jwtmiddleware.JWTMiddleware {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.App.JwtSigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		UserProperty: "auth_token",
	})

	return jwtMiddleware
}
