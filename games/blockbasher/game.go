package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Name       string
	Level      *Level
	Config     *Config
	Controller *Controller
	Paddle     *Paddle
	Player     *Player
	Ball       *Ball
	Scenes     *Scenes
	Started    bool
	Paused     bool
	Logging    bool
	Over       bool
	Time       time.Time
}

/*

<------------------ BLOCKBASHER SPECIFIC ------------------>

*/

func (game *Game) Start() {

	// Default Game Config
	Blockbasher = &Game{
		Name:       GameName,
		Level:      NewLevel(1),
		Started:    false,
		Paused:     false,
		Player:     NewPlayer(),
		Paddle:     NewPaddle(),
		Ball:       NewBall(),
		Controller: NewController(),
		Config:     NewConfig(),
		Scenes: &Scenes{
			MainMenu: NewMainMenuScene(),
			GameOver: NewGameOverScene(),
			Paused:   NewPausedScene(),
		},
		Time: time.Now(),
	}

	// Other Default Settings
	Blockbasher.SetWindowTitle(GameName)
	Blockbasher.SetWindowSize()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

// ebiten.SetWindowSize(game.Configame.WindowSize.Width, game.Configame.WindowSize.Height)
func (game *Game) SetWindowSize() {
	ebiten.SetWindowSize(game.Config.WindowSize.Width, game.Config.WindowSize.Height)
}

// ebiten.SetWindowTitle(title)
func (game *Game) SetWindowTitle(title string) {
	ebiten.SetWindowTitle(title)
}

// detects if img1 collides with img2
func (game *Game) CollisionDetection(img1, img2 image.Rectangle) bool {
	return img1.Overlaps(img2)
}

// detects if img1 collides with img2 and returns the location hit in
func (game *Game) CollisionDetectionAt(img1, img2 image.Rectangle) (top, bottom, left, right bool) {
	if img1.Empty() || img2.Empty() {
		return false, false, false, false
	}

	// Img1 Collision Detection Left Side of Img2
	if img1.Min.Y > img2.Min.Y && img1.Max.Y < img2.Max.Y && img1.Max.X > img2.Min.X && img1.Max.X < img2.Max.X {
		return false, false, true, false
	}

	// Img1 Collision Detection Right Side of Img2
	if img1.Min.Y > img2.Min.Y && img1.Max.Y < img2.Max.Y && img1.Min.X < img2.Max.X && img1.Min.X > img2.Min.X {
		return false, false, false, true
	}

	// Img 1 Collision Detection Top Side of Img2
	if img1.Max.Y > img2.Min.Y && img1.Max.Y < img2.Max.Y && img1.Min.X > img2.Min.X && img1.Max.X < img2.Max.X {
		return true, false, false, false
	}

	// Img1 Collision Detection Bottom Side of Img2
	if img1.Min.Y < img2.Max.Y && img1.Min.Y > img2.Min.Y && img1.Min.X > img2.Min.X && img1.Max.X < img2.Max.X {
		return false, true, false, false
	}

	return false, false, false, false
}

/*

<------------------ EBITEN SPECIFIC ------------------>

*/

func (game *Game) Update() error {

	game.Time = time.Now()

	// Esc | Quit Game
	if inpututil.IsKeyJustPressed(game.Controller.Quit) {
		os.Exit(0)
	}

	// Checking for Game Started
	if !game.Started {
		if inpututil.IsKeyJustPressed(game.Controller.Start) {
			game.Started = true
			game.Ball.IsMoving = true
		}
		return nil
	}

	// Un/Pause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		switch game.Paused {
		case true:
			game.Paused = false
		case false:
			game.Paused = true
			return nil
		}
	}
	if game.Paused {
		return nil
	}

	// Check For Game Over
	if game.Player.Lives == 0 {
		game.Over = true
	}
	if game.Over {
		if inpututil.IsKeyJustPressed(game.Controller.Start) {
			game.Started = false
			game.Over = false
			game.Player = NewPlayer()
			game.Ball = NewBall()
			game.Paddle = NewPaddle()
			game.Level = NewLevel(1)
		}
		return nil
	}

	// Otherwise, Continue To Update The Game:

	// Move the Paddle
	if game.Paddle.IsMoving() {
		game.Paddle.Move()
	}
	// Move the Ball
	if game.Ball.IsMoving {
		go game.Ball.Move()
	} else {
		game.Ball = NewBall()
	}

	return nil

}

func (game *Game) Draw(screen *ebiten.Image) {

	screen.Clear()

	// background color
	screen.Fill(BgColor)

	// Check Game States
	switch {
	case !game.Started:
		game.Scenes.MainMenu.Draw(screen)
		return
	case game.Paused:
		if game.Time.Second()%2 == 0 {
			game.Scenes.Paused.Draw(screen)
		}
		// Paused continues to draw underlying game
	case game.Over:
		game.Scenes.GameOver.Draw(screen)
		return
	}

	// Otherwise Continue To Draw The Game:
	// Update paddle position
	game.Paddle.Draw(screen)

	// Update the ball position
	game.Ball.Draw(screen)

	// Draw Blocks
	for _, row := range game.Level.BlockLayout.Rows {
		for _, block := range row.Blocks {
			if !block.ConfigOpts.Hidden {
				block.Draw(screen)
			}
		}
	}

	// Update the screen
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d\nLives: %d", game.Player.Score, game.Player.Lives))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return DefaultWindowSizeWidth, DefaultWindowSizeHeight
}
