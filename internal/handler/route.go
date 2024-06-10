package handler

import (
	"banner-service/internal/handler/middleware"
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/model/response"
	"github.com/gorilla/mux"
	"net/http"
)

type Service interface {
	GetUserBanner(tagId, featureId uint64, useLast bool) (content string, err error)
	GetAdminBanner(tag_id, featureId, limmit, offset *uint64) (res []response.Banner, err error)
	PostBanner(ban request.Banner) (bannerId uint64, err error)
	DeleteBanner(bannerId uint64) error
	PatchBanner(id uint64, ban request.Banner) error
}

type Handler struct {
	services Service
}

func NewHandler(s Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	r.Handle("/user_banner", middleware.WithUserAuth(http.HandlerFunc(h.UserBanner))).Methods("GET")
	r.Handle("/banner", middleware.WithAdminAuth(http.HandlerFunc(h.GetAdminBanner))).Methods("GET")
	r.Handle("/banner", middleware.WithAdminAuth(http.HandlerFunc(h.PostBanner))).Methods("POST")
	r.Handle("/banner/{id}", middleware.WithAdminAuth(http.HandlerFunc(h.DeleteBanner))).Methods("DELETE")
	r.Handle("/banner/{id}", middleware.WithAdminAuth(http.HandlerFunc(h.PatchBanner))).Methods("PATCH")

	r.Use(middleware.TokenGen)

	return r
}
