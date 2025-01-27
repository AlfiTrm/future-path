package middleware

import (
	"future-path/internal/service"
	"future-path/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Timeout() gin.HandlerFunc
	AuthenticateUser(ctx *gin.Context)
	OnlyAdmin(ctx *gin.Context)
	Cors() gin.HandlerFunc
}

type middleware struct {
	service *service.Service
	jwtAuth jwt.Interface
}

func Init(service *service.Service, jwtAuth jwt.Interface) Interface {
	return &middleware{
		service: service,
		jwtAuth: jwtAuth,
	}
}
