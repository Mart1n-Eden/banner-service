package handler

import (
	"banner-service/internal/handler/middleware"
	"banner-service/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	r.Handle("/user_banner", middleware.WithUserAuth(http.HandlerFunc(h.UserBanner))).Methods("GET")
	r.Handle("/banner", middleware.WithAdminAuth(http.HandlerFunc(h.GetAdminBanner))).Methods("GET")
	r.Handle("/banner", middleware.WithAdminAuth(http.HandlerFunc(h.PostAdminBanner))).Methods("POST")
	r.Handle("/banner/{id}", middleware.WithAdminAuth(http.HandlerFunc(h.DeleteBanner))).Methods("DELETE")
	r.Handle("/banner/{id}", middleware.WithAdminAuth(http.HandlerFunc(h.PatchBanner))).Methods("PATCH")

	r.Use(middleware.TokenGen)

	return r
}
