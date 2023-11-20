package main

type Player struct {
	Name  string
	Score int
	Lives int
	// ActiveBuff
}

func NewPlayer() *Player {
	return &Player{
		Score: 0,
		Lives: 3,
	}
}
