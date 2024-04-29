package priceplans

import (
	"github.com/julienschmidt/httprouter"
	"joi-energy-golang/api"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CompareAll(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	smartMeterId := urlParams.ByName("smartMeterId")
	err := validateSmartMeterId(smartMeterId)
	if err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := h.service.CompareAllPricePlans(smartMeterId)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.SuccessJson(w, r, result)
}

func (h *Handler) Recommend(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	smartMeterId := urlParams.ByName("smartMeterId")
	err := validateSmartMeterId(smartMeterId)
	if err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.ParseUint(limitString, 10, 64)
	if limitString != "" && err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := h.service.RecommendPricePlans(smartMeterId, limit)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.SuccessJson(w, r, result)
}
