package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetSekolahNegeri(ctx *gin.Context) {
	namaQuery := ctx.Query("sekolah")
	sekolah, err := r.service.SekolahService.GetSekolahNegeri(namaQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get school", err)
		return
	}

	var sekolahResponses []model.GetSekolah
	for _, b := range sekolah {
		sekolahResponses = append(sekolahResponses, model.GetSekolah{
			ID_Sekolah:        b.ID_Sekolah,
			Nama_Sekolah:      b.Nama_Sekolah,
			Alamat_Sekolah:    b.Alamat_Sekolah,
			Deskripsi_Sekolah: b.Deskripsi_Sekolah,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get school", sekolahResponses)
}

func (r *Rest) GetSekolahSwasta(ctx *gin.Context) {
	namaQuery := ctx.Query("sekolah")
	sekolah, err := r.service.SekolahService.GetSekolahSwasta(namaQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get school", err)
		return
	}

	var sekolahResponses []model.GetSekolah
	for _, b := range sekolah {
		sekolahResponses = append(sekolahResponses, model.GetSekolah{
			ID_Sekolah:        b.ID_Sekolah,
			Nama_Sekolah:      b.Nama_Sekolah,
			Alamat_Sekolah:    b.Alamat_Sekolah,
			Deskripsi_Sekolah: b.Deskripsi_Sekolah,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get school", sekolahResponses)
}

func (r *Rest) GetAllSekolah(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	sekolah, err, totalData := r.service.SekolahService.GetAllSekolah(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get short news", err)
		return
	}

	responseData := struct {
		TotalData int64                 `json:"total_data"`
		Sekolah   []model.GetAllSekolah `json:"sekolah"`
	}{
		TotalData: totalData,
	}

	for _, b := range sekolah {
		responseData.Sekolah = append(responseData.Sekolah, model.GetAllSekolah{
			ID_Sekolah:     b.ID_Sekolah,
			Nama_Sekolah:   b.Nama_Sekolah,
			Alamat_Sekolah: b.Alamat_Sekolah,
		})
	}

	response.Success(ctx, http.StatusOK, "Schools retrieved", responseData)
}

func (r *Rest) GetSekolahDetail(ctx *gin.Context) {
	idParam := ctx.Query("id_sekolah")

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "School ID is required", nil)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid school ID", err)
		return
	}

	sekolah, err := r.service.SekolahService.GetSekolahDetail(idInt)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to get school", err)
		return
	}

	responses := model.GetSekolah{
		Nama_Sekolah:      sekolah.Nama_Sekolah,
		Alamat_Sekolah:    sekolah.Alamat_Sekolah,
		Deskripsi_Sekolah: sekolah.Deskripsi_Sekolah,
	}

	response.Success(ctx, http.StatusOK, "School retrieved", responses)
}

func (r *Rest) AddSekolah(ctx *gin.Context) {
	photoReq, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
	}

	sekolahReq := model.CreateSekolah{
		Nama_Sekolah:      ctx.PostForm("nama_sekolah"),
		Alamat_Sekolah:    ctx.PostForm("alamat_sekolah"),
		Deskripsi_Sekolah: ctx.PostForm("deskripsi_sekolah"),
	}

	photo := model.UploadPhoto{
		Photo: photoReq,
	}

	sekolah, err := r.service.SekolahService.AddSekolah(&sekolahReq, photo)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to add school", err)
		return
	}

	responses := model.CreateSekolahResponse{
		Nama_Sekolah:      sekolah.Nama_Sekolah,
		Alamat_Sekolah:    sekolah.Alamat_Sekolah,
		Deskripsi_Sekolah: sekolah.Deskripsi_Sekolah,
		ID_Kepemilikan:    sekolah.ID_Kepemilikan,
	}

	response.Success(ctx, http.StatusOK, "Successfully to add school", responses)
}

// func (r *Rest) AddSekolah(ctx *gin.Context) {
// 	var sekolahReq model.CreateSekolah

// 	err := ctx.ShouldBind(&sekolahReq)
// 	if err != nil {
// 		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind input", err)
// 		return
// 	}

// 	photoReq, err := ctx.FormFile("photo")
// 	if err != nil {
// 		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
// 	}

// 	photo := model.UploadPhoto{
// 		Photo: photoReq,
// 	}

// 	sekolah, err := r.service.SekolahService.AddSekolah(&sekolahReq, photo)
// 	if err != nil {
// 		response.Error(ctx, http.StatusInternalServerError, "Failed to add school", err)
// 		return
// 	}

// 	responses := model.CreateSekolahResponse{
// 		Nama_Sekolah:      sekolah.Nama_Sekolah,
// 		Alamat_Sekolah:    sekolah.Alamat_Sekolah,
// 		Deskripsi_Sekolah: sekolah.Deskripsi_Sekolah,
// 		ID_Kepemilikan:    sekolah.ID_Kepemilikan,
// 		PhotoLink:         sekolah.PhotoLink,
// 	}

// 	response.Success(ctx, http.StatusOK, "Successfully to add school", responses)
// }
