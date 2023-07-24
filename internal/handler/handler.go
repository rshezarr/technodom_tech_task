package handler

import (
	"github.com/gin-gonic/gin"
	"redirect_api/internal/service"
)

// handler class
type Handler struct {
	router  *gin.Engine
	service *service.Service
}

// handler object
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		router:  gin.New(),
		service: service,
	}
}

// all endpoints
func (h *Handler) InitRoutes() *gin.Engine {
	//admin endpoints
	h.router.GET("/admin/redirects", h.getRecords)
	h.router.GET("/admin/redirects/", h.getOneRecord)
	h.router.POST("/admin/redirects", h.insertNewRecord)
	h.router.PATCH("/admin/redirects/", h.editRecord)
	h.router.DELETE("/admin/redirects/", h.deleteRecords)

	//user endpoints
	h.router.GET("/redirects", h.redirectUser)

	return h.router
}
