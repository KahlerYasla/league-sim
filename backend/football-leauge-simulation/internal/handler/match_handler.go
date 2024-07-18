package handler

import (
	"encoding/json"
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MatchHandler struct {
	service *service.MatchService
}

func NewMatchHandler(service *service.MatchService) *MatchHandler {
	return &MatchHandler{service: service}
}

func (h *MatchHandler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match domain.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateMatch(&match); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func (h *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.service.GetMatches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

func (h *MatchHandler) GetMatchByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	match, err := h.service.GetMatchByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}
