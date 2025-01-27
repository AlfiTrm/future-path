package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secret := os.Getenv("SESSION_SECRET")
		gothic.Store = sessions.NewCookieStore([]byte(secret))
		ctx.Next()
	}
}
