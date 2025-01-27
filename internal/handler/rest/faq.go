package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateFAQ(ctx *gin.Context) {
	var FAQReq model.CreateFAQ

	if err := ctx.ShouldBindJSON(&FAQReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid Request", err)
		return
	}

	faq, err := r.service.FAQService.CreateFAQ(&FAQReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create FAQ", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "FAQ Created", faq)
}

func (r *Rest) GetFAQ(ctx *gin.Context) {
	faq, err := r.service.FAQService.GetFAQ()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get FAQ", err)
		return
	}

	var FAQResponse []model.GetFAQ
	for _, b := range faq {
		FAQResponse = append(FAQResponse, model.GetFAQ{
			ID_FAQ:    b.ID_FAQ,
			Judul_FAQ: b.Judul_FAQ,
			Isi_FAQ:   b.Isi_FAQ,
		})
	}

	response.Success(ctx, http.StatusOK, "FAQ retrieved", FAQResponse)
}

func (r *Rest) UpdateFAQ(ctx *gin.Context) {
	idParam := ctx.Param("id_faq")
	var FAQRequest model.UpdateFAQ

	if err := ctx.ShouldBindJSON(&FAQRequest); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
	}

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "FAQ ID is required", nil)
		return
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid ID FAQ", err)
	}

	faq, err := r.service.FAQService.UpdateFAQ(idInt, &FAQRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Faild to update FAQ", err)
		return
	}

	response.Success(ctx, http.StatusOK, "FAQ Updated", faq)
}

func (r *Rest) DeleteFAQ(ctx *gin.Context) {
	idParam := ctx.Param("id_faq")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid FAQ ID", err)
		return
	}

	if err := r.service.FAQService.DeleteFAQ(idInt); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete FAQ", err)
	}

	response.Success(ctx, http.StatusOK, "FAQ deleted successfully", nil)
}
