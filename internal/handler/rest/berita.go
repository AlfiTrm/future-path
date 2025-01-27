package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateBerita(ctx *gin.Context) {
	var beritaReq model.CreateBerita

	if err := ctx.ShouldBindJSON(&beritaReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid Request", err)
		return
	}

	berita, err := r.service.BeritaService.CreateBerita(&beritaReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create news", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "News created", berita)

}

func (r *Rest) GetBeritaSingkat(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	berita, err, totalData := r.service.BeritaService.GetBeritaSingkat(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get short news", err)
		return
	}

	responseData := struct {
		TotalData int64             `json:"total_data"`
		Berita    []model.GetBerita `json:"berita"`
	}{
		TotalData: totalData,
	}

	for _, b := range berita {
		responseData.Berita = append(responseData.Berita, model.GetBerita{
			ID_Berita:    b.ID_Berita,
			Judul_Berita: b.Judul_Berita,
			Isi_Berita:   b.Isi_Berita,
		})
	}

	response.Success(ctx, http.StatusOK, "Short news retrieved", responseData)

}

func (r *Rest) GetBeritaFull(ctx *gin.Context) {
	idParam := ctx.Query("id_berita")

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "News ID is required", nil)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid news ID", err)
		return
	}

	berita, err := r.service.BeritaService.GetBeritaFull(idInt)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to get news", err)
		return
	}

	responses := model.GetBerita{
		ID_Berita:    berita.ID_Berita,
		Judul_Berita: berita.Judul_Berita,
		Isi_Berita:   berita.Isi_Berita,
	}

	response.Success(ctx, http.StatusOK, "Full news retrieved", responses)
}

func (r *Rest) UpdateBerita(ctx *gin.Context) {
	idParam := ctx.Param("id_berita")
	var beritaRequest model.UpdateBerita

	if err := ctx.ShouldBindJSON(&beritaRequest); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
	}

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "News ID is required", nil)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid ID news", err)
	}

	berita, err := r.service.BeritaService.UpdateBerita(idInt, &beritaRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update news", err)
		return
	}

	response.Success(ctx, http.StatusOK, "News updated", berita)

}

func (r *Rest) DeleteBerita(ctx *gin.Context) {
	idParam := ctx.Param("id_berita")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid news ID", err)
		return
	}

	if err := r.service.BeritaService.DeleteBerita(idInt); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete news", err)
	}

	response.Success(ctx, http.StatusOK, "News deleted successfully", nil)

}
