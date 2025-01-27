package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	if len(param.Password_User) < 8 {
		response.Error(ctx, http.StatusBadRequest, "password less than 8 character", nil)
	}

	err = r.service.UserService.Register(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "successfully register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	result, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	responses := model.UserLoginResponses{
		Token:  result.Token,
		RoleID: result.RoleID,
	}

	response.Success(ctx, http.StatusOK, "successfully login to system", responses)

}
