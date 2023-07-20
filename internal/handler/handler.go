package handler

import (
	"github.com/gin-gonic/gin"
	"redirect_api/internal/service"
)

type Handler struct {
	router  *gin.Engine
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		router:  gin.New(),
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// admin routes
	h.router.GET("/admin/redirects", h.getRecords)
	h.router.GET("/admin/redirects/", h.getOneRecord)
	h.router.POST("/admin/redirects", h.insertNewRecord)
	h.router.PATCH("/admin/redirects/", h.editRecord)
	h.router.DELETE("/admin/redirects/", h.deleteRecords)

	// user routes
	h.router.GET("/redirects", h.redirectUser)

	return h.router
}
