package repository

import (
	"errors"

	"github.com/rishabh/fantasyCricket/entities"
)

type InMemoryPlayerRepository struct {
	players map[int]entities.Player
}

func NewInMemoryPlayerRepository() *InMemoryPlayerRepository {
	return &InMemoryPlayerRepository{players: make(map[int]entities.Player)}
}

func (repo *InMemoryPlayerRepository) GetPlayerByID(id int) (entities.Player, error) {
	player, exists := repo.players[id]
	if !exists {
		return player, errors.New("player not found")
	}
	return player, nil
}

func (repo *InMemoryPlayerRepository) UpdatePlayer(player entities.Player) error {
	repo.players[player.ID] = player
	return nil
}
