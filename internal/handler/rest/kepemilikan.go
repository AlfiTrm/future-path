package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetKepemilikan(ctx *gin.Context) {
	kepemilikan, err := r.service.KepemilikanService.GetKepemilikan()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get ownerships", err)
		return
	}

	var kepemilikanResponse []model.KepemilikanResponse
	for _, b := range kepemilikan {
		kepemilikanResponse = append(kepemilikanResponse, model.KepemilikanResponse{
			ID_Kepemilikan:   b.ID_Kepemilikan,
			Nama_Kepemilikan: b.Nama_Kepemilikan,
		})
	}

	response.Success(ctx, http.StatusOK, "Ownerships found", kepemilikanResponse)
}
