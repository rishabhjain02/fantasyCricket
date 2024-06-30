package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishabh/fantasyCricket/delivery"
	"github.com/rishabh/fantasyCricket/entities"
	"github.com/rishabh/fantasyCricket/repository"
	"github.com/rishabh/fantasyCricket/usecase"
)

func main() {
	playerRepo := repository.NewInMemoryPlayerRepository()
	teamRepo := repository.NewInMemoryTeamRepository()
	teamService := usecase.NewTeamService(playerRepo, teamRepo)
	httpHandler := delivery.NewHTTPHandler(*teamService)

	// Prepopulate player repository with some players
	players := []entities.Player{
		{ID: 1, Name: "Player 1", Type: entities.Batsman, Credits: 10, Points: 0},
		{ID: 2, Name: "Player 2", Type: entities.Batsman, Credits: 9, Points: 0},
		{ID: 3, Name: "Player 3", Type: entities.Batsman, Credits: 8, Points: 0},
		{ID: 4, Name: "Player 4", Type: entities.Bowler, Credits: 10, Points: 0},
		{ID: 5, Name: "Player 5", Type: entities.Bowler, Credits: 9, Points: 0},
		{ID: 6, Name: "Player 6", Type: entities.Bowler, Credits: 8, Points: 0},
		{ID: 7, Name: "Player 7", Type: entities.WicketKeeper, Credits: 10, Points: 0},
		{ID: 8, Name: "Player 8", Type: entities.Batsman, Credits: 7, Points: 0},
		{ID: 9, Name: "Player 9", Type: entities.Bowler, Credits: 7, Points: 0},
		{ID: 10, Name: "Player 10", Type: entities.Batsman, Credits: 6, Points: 0},
		{ID: 11, Name: "Player 11", Type: entities.Bowler, Credits: 6, Points: 0},
	}

	for _, player := range players {
		playerRepo.UpdatePlayer(player)
	}

	r := mux.NewRouter()
	r.HandleFunc("/create-team", httpHandler.CreateTeamHandler).Methods("POST")
	r.HandleFunc("/team-points", httpHandler.GetTeamPointsHandler).Methods("GET")
	r.HandleFunc("/update-player-points/{player_id}", httpHandler.UpdatePlayerPointsHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
