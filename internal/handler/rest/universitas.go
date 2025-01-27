package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetUnivNegeri(ctx *gin.Context) {
	namaUniv := ctx.Query("universitas")
	univ, err := r.service.UniversitasService.GetUnivNegeri(namaUniv)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get university", err)
		return
	}
	var univResponse []model.GetUniversitas
	for _, b := range univ {
		univResponse = append(univResponse, model.GetUniversitas{
			Nama_Universitas:      b.Nama_Universitas,
			Alamat_Universitas:    b.Alamat_Universitas,
			Deskripsi_Universitas: b.Deskripsi_Universitas,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get university", univResponse)
}

func (r *Rest) GetUnivSwasta(ctx *gin.Context) {
	namaUniv := ctx.Query("universitas")
	univ, err := r.service.UniversitasService.GetUnivSwasta(namaUniv)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get university", err)
		return
	}
	var univResponse []model.GetUniversitas
	for _, b := range univ {
		univResponse = append(univResponse, model.GetUniversitas{
			Nama_Universitas:      b.Nama_Universitas,
			Alamat_Universitas:    b.Alamat_Universitas,
			Deskripsi_Universitas: b.Deskripsi_Universitas,
		})
	}
	response.Success(ctx, http.StatusOK, "success to get university", univResponse)
}

func (r *Rest) GetAllUniv(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	univ, err := r.service.UniversitasService.GetAllUniv(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get short news", err)
		return
	}

	var univResponse []model.GetAllUniv
	for _, b := range univ {
		univResponse = append(univResponse, model.GetAllUniv{
			Nama_Universitas:   b.Nama_Universitas,
			Alamat_Universitas: b.Alamat_Universitas,
		})
	}

	response.Success(ctx, http.StatusOK, "University retrieved", univResponse)
}

func (r *Rest) GetUnivDetail(ctx *gin.Context) {
	idParam := ctx.Query("id_universitas")

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "University ID is required", nil)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid university ID", err)
		return
	}

	univ, err := r.service.UniversitasService.GetUnivDetail(idInt)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to get university", err)
		return
	}

	responses := model.GetUniversitas{
		Nama_Universitas:      univ.Nama_Universitas,
		Alamat_Universitas:    univ.Alamat_Universitas,
		Deskripsi_Universitas: univ.Deskripsi_Universitas,
	}

	response.Success(ctx, http.StatusOK, "University retrieved", responses)
}

func (r *Rest) AddUniv(ctx *gin.Context) {
	var univReq model.CreateUniv
	err := ctx.ShouldBindJSON(&univReq)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind input", err)
		return
	}

	univ, err := r.service.UniversitasService.AddUniv(&univReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to add university", err)
		return
	}

	responses := model.CreateUnivResponse{
		Nama_Universitas:      univ.Nama_Universitas,
		Alamat_Universitas:    univ.Alamat_Universitas,
		Deskripsi_Universitas: univ.Deskripsi_Universitas,
		ID_Kepemilikan:        univ.ID_Kepemilikan,
	}

	response.Success(ctx, http.StatusOK, "Successfully to add university", responses)
}
