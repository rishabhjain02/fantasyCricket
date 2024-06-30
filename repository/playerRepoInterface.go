package repository

import "github.com/rishabh/fantasyCricket/entities"

type PlayerRepository interface {
	GetPlayerByID(id int) (entities.Player, error)
	UpdatePlayer(player entities.Player) error
}
