package rest

import (
	"fmt"
	"future-path/internal/service"
	"future-path/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Cors(), r.middleware.Timeout(), middleware.SessionMiddleware())
	routerGroup := r.router.Group("/future-path")
	// routerGroup.Use(r.middleware.Timeout())
	// routerGroup.GET("/testing", testTimeout)

	auth := routerGroup.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login", r.Login)

	auth.GET("/login/:provider", r.OAuthLogin)
	auth.GET("/callback/:provider/", r.OAuthCallback)

	user := routerGroup.Group("/user")
	user.Use(r.middleware.AuthenticateUser)
	user.GET("/berita", r.GetBeritaSingkat)
	user.GET("/full-news", r.GetBeritaFull)
	user.GET("/list-sekolah", r.GetAllSekolah)
	user.GET("/sekolah", r.GetSekolahDetail)
	user.GET("/cari-sekolah/negeri", r.GetSekolahNegeri)
	user.GET("/cari-sekolah/swasta", r.GetSekolahSwasta)
	user.GET("/list-universitas", r.GetAllUniv)
	user.GET("/universitas", r.GetUnivDetail)
	user.GET("/cari-universitas/negeri", r.GetUnivNegeri)
	user.GET("/cari-universitas/swasta", r.GetUnivSwasta)
	user.GET("/faq", r.GetFAQ)

	admin := routerGroup.Group("/admin")
	admin.Use(r.middleware.AuthenticateUser, r.middleware.OnlyAdmin)
	admin.POST("/create-berita", r.CreateBerita)
	admin.PATCH("/update-berita/:id_berita", r.UpdateBerita)
	admin.DELETE("/delete-berita/:id_berita", r.DeleteBerita)
	admin.POST("/create-faq", r.CreateFAQ)
	admin.PATCH("/update-faq/:id_faq", r.UpdateFAQ)
	admin.DELETE("/delete-faq/:id_faq", r.DeleteFAQ)
	admin.GET("/get-ownerships", r.GetKepemilikan)
	admin.POST("/add-sekolah", r.AddSekolah)
	admin.POST("add-universitas", r.AddUniv)
}

// func testTimeout(ctx *gin.Context) {
// 	time.Sleep(3 * time.Second)

// 	response.Success(ctx, http.StatusOK, "success", nil)
// }

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
