package main

type gameState struct {
	Started bool
	Paused  bool
	Level   int
	Lives   int
	Score   int
	Quit    bool
}

type ball struct {
	x       float64
	y       float64
	Length  int
	Width   int
	Speed   int
	Visible bool
}

type bumper struct {
	x       float64
	y       float64
	Length  int
	Width   int
	Visible bool
	Moving  bool
}

type blocks struct {
	Total     int
	Remaining int
	Blocks    []block
}

type block struct {
	Length  int
	Width   int
	Visible bool
	Cracked bool
}
