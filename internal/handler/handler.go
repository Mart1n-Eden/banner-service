package handler

import (
	"banner-service/internal/handler/tools"
	"banner-service/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user_banner", h.userBanner).Methods("GET")
	r.HandleFunc("/banner", h.userBanner).Methods("GET")
	r.HandleFunc("/banner", h.userBanner).Methods("POST")

	return r
}

func (h *Handler) userBanner(w http.ResponseWriter, r *http.Request) {
	var tagId, featureId uint64
	//var useLast bool
	var err error

	if strTagId := r.URL.Query().Get("tag_id"); strTagId == "" {
		tools.SendError(w, http.StatusBadRequest, "tag_id is empty")
		return
	} else if tagId, err = strconv.ParseUint(strTagId, 10, 64); err != nil {
		tools.SendError(w, http.StatusBadRequest, "tag_id is not a number")
		return
	}

	if strFeatureId := r.URL.Query().Get("feature_id"); strFeatureId == "" {
		tools.SendError(w, http.StatusBadRequest, "feature_id id empty")
		return
	} else if featureId, err = strconv.ParseUint(strFeatureId, 10, 64); err != nil {
		tools.SendError(w, http.StatusBadRequest, "feature_id is not a number")
		return
	}

	// TODO: use_last_version realisation

	// TODO: check token

	// TODO: take banner from cache

	if content, err := h.services.GetUserBanner(tagId, featureId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tools.SendError(w, http.StatusNotFound, "banner not found")
		} else {
			// TODO: correct return InternalServerError
			tools.SendError(w, http.StatusInternalServerError, err.Error())
		}
	} else {
		tools.SendSucsess(w, http.StatusOK, content)
	}
}

/*
func (h *Handler) getAdminBanner(w http.ResponseWriter, r *http.Request) {
	var tagId, featureId, limit, offset uint64
	//var useLast bool
	var err error

	if strTagId := r.URL.Query().Get("tag_id"); strTagId != "" {
		if tagId, err = strconv.ParseUint(strTagId, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "tag_id is not a number")
			return
		}
	}

	if strFeatureId := r.URL.Query().Get("feature_id"); strFeatureId != "" {
		if featureId, err = strconv.ParseUint(strFeatureId, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "feature_id is not a number")
			return
		}
	}

	if strLimit := r.URL.Query().Get("limit"); strLimit != "" {
		if limit, err = strconv.ParseUint(strLimit, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "limit is not a number")
			return
		}
	}

	if strOffset := r.URL.Query().Get("offset"); strOffset != "" {
		if offset, err = strconv.ParseUint(strOffset, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "offset is not a number")
			return
		}
	}

	// TODO: continue

}

func (h *Handler) postAdminBanner(w http.ResponseWriter, r *http.Request) {
	var ban request.PostBanner

	if err := json.NewDecoder(r.Body).Decode(&ban); err != nil {
		tools.SendError(w, http.StatusBadRequest, "incorrect data") //check
	}

}
*/
