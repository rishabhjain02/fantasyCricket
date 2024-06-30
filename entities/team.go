package entities

import "errors"

type Team struct {
	ID      int
	UserID  int
	Players []Player
}

func (t *Team) Validate() error {
	var batsmen, bowlers, keepers int
	for _, player := range t.Players {
		switch player.Type {
		case Batsman:
			batsmen++
		case Bowler:
			bowlers++
		case WicketKeeper:
			keepers++
		}
	}
	if keepers < 1 {
		return errors.New("at least one wicket keeper is required")
	}
	if batsmen < 3 {
		return errors.New("at least three batsmen are required")
	}
	if bowlers < 3 {
		return errors.New("at least three bowlers are required")
	}
	return nil
}

func (t *Team) TotalCredits() int {
	totalCredits := 0
	for _, player := range t.Players {
		totalCredits += player.Credits
	}
	return totalCredits
}

func (t *Team) TotalPoints() int {
	totalPoints := 0
	for _, player := range t.Players {
		totalPoints += player.Points
	}
	return totalPoints
}
