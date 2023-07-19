package handler

import (
	"net/http"

	"redirect_api/internal/service"
)

type Handler struct {
	router  *http.ServeMux
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		router:  http.NewServeMux(),
		service: service,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.router.HandleFunc("/admin/redirects", h.getRecords)
	h.router.HandleFunc("/admin/redirects/", h.getOneRecord)
	h.router.HandleFunc("/admin/redirects", h.insertNewRecord)
	h.router.HandleFunc("/admin/redirects/", h.editRecord)
	h.router.HandleFunc("/admin/redirects/", h.deleteRecords)

	// user route
	h.router.HandleFunc("/redirects", h.redirectUser)

	return h.router
}
