package usecase

import (
	"errors"

	"github.com/rishabh/fantasyCricket/entities"
	"github.com/rishabh/fantasyCricket/repository"
)

type TeamService struct {
	playerRepo repository.PlayerRepository
	teamRepo   repository.TeamRepository
}

func NewTeamService(playerRepo repository.PlayerRepository, teamRepo repository.TeamRepository) *TeamService {
	return &TeamService{playerRepo: playerRepo, teamRepo: teamRepo}
}

func (ts *TeamService) CreateTeam(userID int, playerIDs []int) error {
	var players []entities.Player
	totalCredits := 0

	for _, playerID := range playerIDs {
		player, err := ts.playerRepo.GetPlayerByID(playerID)
		if err != nil {
			return err
		}
		players = append(players, player)
		totalCredits += player.Credits
	}

	if totalCredits > 100 {
		return errors.New("team exceeds budget of 100 credits")
	}

	team := entities.Team{
		UserID:  userID,
		Players: players,
	}

	if err := team.Validate(); err != nil {
		return err
	}

	return ts.teamRepo.SaveTeam(team)
}

func (ts *TeamService) GetTeamPoints(userID int) (int, error) {
	team, err := ts.teamRepo.GetTeamByUserID(userID)
	if err != nil {
		return 0, err
	}

	totalPoints := 0
	for _, player := range team.Players {
		updatedPlayer, err := ts.playerRepo.GetPlayerByID(player.ID)
		if err != nil {
			return 0, err
		}
		totalPoints += updatedPlayer.Points
	}

	return totalPoints, nil
}

func (ts *TeamService) UpdatePlayerPoints(playerID int, points int) error {
	player, err := ts.playerRepo.GetPlayerByID(playerID)
	if err != nil {
		return err
	}
	player.Points = points
	return ts.playerRepo.UpdatePlayer(player)
}
