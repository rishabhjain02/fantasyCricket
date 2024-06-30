package repository

import (
	"errors"

	"github.com/rishabh/fantasyCricket/entities"
)

type InMemoryTeamRepository struct {
	teams map[int]entities.Team
}

func NewInMemoryTeamRepository() *InMemoryTeamRepository {
	return &InMemoryTeamRepository{teams: make(map[int]entities.Team)}
}

func (repo *InMemoryTeamRepository) SaveTeam(team entities.Team) error {
	repo.teams[team.UserID] = team
	return nil
}

func (repo *InMemoryTeamRepository) GetTeamByUserID(userID int) (entities.Team, error) {
	team, exists := repo.teams[userID]
	if !exists {
		return team, errors.New("team not found")
	}
	return team, nil
}
