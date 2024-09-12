package handler

import (
	"football-league-simulation/internal/service"
	"net/http"
)

type SimulationHandler struct {
	service *service.SimulationService
}

func NewSimulationHandler(service *service.SimulationService) *SimulationHandler {
	return &SimulationHandler{service: service}
}

func (h *SimulationHandler) SimulateWeek(w http.ResponseWriter, r *http.Request) {
	week := 1 // In a real application, you might get this from the request
	if err := h.service.SimulateWeek(week); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Week simulated successfully"))
}

func (h *SimulationHandler) SimulateSeason(w http.ResponseWriter, r *http.Request) {
	if err := h.service.SimulateSeason(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Season simulated successfully"))
}
