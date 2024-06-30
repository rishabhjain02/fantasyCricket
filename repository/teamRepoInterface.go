package repository

import "github.com/rishabh/fantasyCricket/entities"

type TeamRepository interface {
	SaveTeam(team entities.Team) error
	GetTeamByUserID(userID int) (entities.Team, error)
}
