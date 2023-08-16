package bumpball

import (
	"fmt"
	"net/http"
)

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
	x       float64
	y       float64
	Length  int
	Width   int
	Visible bool
	Cracked bool
}

func Play(writer http.ResponseWriter, req *http.Request) {

	var (
		gameState = new(gameState)
		ball      = &ball{
			x:       0.0,
			y:       0.0,
			Length:  2,
			Width:   2,
			Speed:   1,
			Visible: false,
		}
		bumper = &bumper{
			x:       0.0,
			y:       0.0,
			Length:  2,
			Width:   2,
			Visible: false,
		}
		blocks = &blocks{
			Total:     0,
			Remaining: 0,
		}
	)

	fmt.Fprint(writer, "<h1>Future Page For Bump Ball</h1>")

	// Loop game until Quit
	for !gameState.Quit {
		// Loop game until Started
		for !gameState.Started {
			// display start screen,
			// await input
			// once we receive it, switch gameState.Started = true
			gameState.Started = true
		}
		// Loop game while Started
		for gameState.Started {
			// Break if the game was exited/quit
			if gameState.Quit {
				break
			}
			for gameState.Paused {
				// Pause the game

				// Unpause the game
				gameState.Paused = false
			}
			// else continue to draw the next frame
			for _, b := range blocks.Blocks {
				if b.Visible {
					// draw it
					fmt.Println("draw block")
				}
			}
			if ball.Visible {
				// draw it
				fmt.Println("draw ball")
			}
			if bumper.Visible {
				// draw it
				fmt.Println("draw bumper")
			}
		}
	}

}
