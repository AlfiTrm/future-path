package middleware

import (
	"errors"
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userID, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to validate token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUser(model.UserParam{
		ID_User: userID,
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to get user", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
