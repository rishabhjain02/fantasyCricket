package entities

type PlayerType string

const (
	Batsman      PlayerType = "Batsman"
	Bowler       PlayerType = "Bowler"
	WicketKeeper PlayerType = "WicketKeeper"
)

type Player struct {
	ID      int
	Name    string
	Type    PlayerType
	Credits int
	Points  int
}
