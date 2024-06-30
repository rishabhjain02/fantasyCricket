package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rishabh/fantasyCricket/usecase"
)

type HTTPHandler struct {
	teamService usecase.TeamService
}

func NewHTTPHandler(teamService usecase.TeamService) *HTTPHandler {
	return &HTTPHandler{teamService: teamService}
}

func (h *HTTPHandler) CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	playerIDs := r.Form["player_ids"]
	var playerIDsInt []int
	for _, id := range playerIDs {
		pid, _ := strconv.Atoi(id)
		playerIDsInt = append(playerIDsInt, pid)
	}
	err := h.teamService.CreateTeam(userID, playerIDsInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) GetTeamPointsHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	points, err := h.teamService.GetTeamPoints(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

func (h *HTTPHandler) UpdatePlayerPointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, _ := strconv.Atoi(mux.Vars(r)["player_id"])
	points, _ := strconv.Atoi(r.FormValue("points"))
	err := h.teamService.UpdatePlayerPoints(playerID, points)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
