package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (r *Rest) OAuthLogin(ctx *gin.Context) {
	provider := ctx.Param("provider")
	if provider == "" {
		response.Error(ctx, http.StatusBadRequest, "Provider is required", nil)
		return
	}

	ctx.Request = gothic.GetContextWithProvider(ctx.Request, provider)
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (r *Rest) OAuthCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	if provider == "" {
		response.Error(ctx, http.StatusBadRequest, "Provider is required", nil)
		return
	}

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "OAuth authentication failed", err)
		return
	}

	authUser, err := r.service.OAuthService.HandleGoogleLogin(user.Email, user.FirstName, user.AccessToken)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to process user", err)
		return
	}

	responsesAuth := model.UserLoginResponses{
		Token:  authUser.Token,
		RoleID: authUser.RoleID,
	}

	response.Success(ctx, http.StatusOK, "Successfully authenticated user", responsesAuth)
}
