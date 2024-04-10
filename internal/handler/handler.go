package handler

import (
	"banner-service/internal/handler/model/request"
	"banner-service/internal/handler/tools"
	"banner-service/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/banner", h.getAdminBanner).Methods("GET")
	r.HandleFunc("/banner", h.postAdminBanner).Methods("POST")
	r.HandleFunc("/banner/{id}", h.deleteBanner).Methods("DELETE")
	r.HandleFunc("/banner/{id}", h.patchBanner).Methods("PATCH")

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
			return
		} else {
			// TODO: correct return InternalServerError
			tools.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		tools.SendSucsessContent(w, http.StatusOK, content)
	}
}

func (h *Handler) getAdminBanner(w http.ResponseWriter, r *http.Request) {
	//tagId, featureId, limit, offset := new(uint64), new(uint64), new(uint64), new(uint64)
	var tagId, featureId, limit, offset *uint64
	var err error
	fmt.Println(tagId, featureId, limit, offset)

	if strTagId := r.URL.Query().Get("tag_id"); strTagId != "" {
		tagId = new(uint64)
		fmt.Println(tagId)
		if *tagId, err = strconv.ParseUint(strTagId, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "tag_id is not a number")
			return
		}
	}

	if strFeatureId := r.URL.Query().Get("feature_id"); strFeatureId != "" {
		featureId = new(uint64)
		if *featureId, err = strconv.ParseUint(strFeatureId, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "feature_id is not a number")
			return
		}
	}

	if strLimit := r.URL.Query().Get("limit"); strLimit != "" {
		limit = new(uint64)
		if *limit, err = strconv.ParseUint(strLimit, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "limit is not a number")
			return
		}
	}

	if strOffset := r.URL.Query().Get("offset"); strOffset != "" {
		offset = new(uint64)
		if *offset, err = strconv.ParseUint(strOffset, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "offset is not a number")
			return
		}
	}

	if res, err := h.services.GetAdminBanner(tagId, featureId, limit, offset); err != nil {
		tools.SendError(w, http.StatusBadRequest, err.Error()) // TODO:
	} else {
		switch {
		case res != nil:
			tools.SendSucsessArray(w, http.StatusOK, res)
		case res == nil:
			// TODO: return empty array for json
			tools.SendSucsessArray(w, http.StatusOK, []byte{})
		}
	}
}

func (h *Handler) postAdminBanner(w http.ResponseWriter, r *http.Request) {
	var ban request.Banner

	if err := json.NewDecoder(r.Body).Decode(&ban); err != nil {
		tools.SendError(w, http.StatusBadRequest, "incorrect data") //check
		return
	}

	if bannerId, err := h.services.PostBanner(ban); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: correct return error message
			tools.SendError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			// TODO: correct return InternalServerError
			tools.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		tools.SendSucsessId(w, http.StatusCreated, bannerId)
	}

}

func (h *Handler) deleteBanner(w http.ResponseWriter, r *http.Request) {
	var bannerId uint64
	var err error

	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		if bannerId, err = strconv.ParseUint(id, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "id is not a number")
			return
		}
	} else {
		tools.SendError(w, http.StatusBadRequest, "id is empty")
		return
	}

	if err := h.services.DeleteBanner(bannerId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tools.SendError(w, http.StatusNotFound, "banner not found")
			return
		} else {
			// TODO: correct return InternalServerError
			tools.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		tools.SendStatus(w, http.StatusNoContent)
	}
}

func (h *Handler) patchBanner(w http.ResponseWriter, r *http.Request) {
	var ban request.Banner

	if err := json.NewDecoder(r.Body).Decode(&ban); err != nil {
		tools.SendError(w, http.StatusBadRequest, "incorrect data") //check
		return
	}

	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		if bannerId, err := strconv.ParseUint(id, 10, 64); err != nil {
			tools.SendError(w, http.StatusBadRequest, "id is not a number")
			return
		} else {
			if err := h.services.PatchBanner(bannerId, ban); err != nil {
				// TODO: correct return error message
				tools.SendError(w, http.StatusBadRequest, err.Error())
			} else {
				tools.SendStatus(w, http.StatusOK)
			}
		}
	} else {
		tools.SendError(w, http.StatusBadRequest, "id is empty")
		return
	}
}
