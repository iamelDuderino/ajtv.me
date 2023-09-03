package main

import "fmt"

/*

This is all currently pseudo-code from my original JS bumpball sample

Prepping for update to be en Ebiten application that runs in UI via IFrame

*/

func main() {
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
