package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GameName    = "Block Basher"
	PaddlePNG   = "./img/paddle.png"
	BallPNG     = "./img/ball.png"
	MainMenuPNG = "./img/mainmenu2.png"
	BlockPNG    = "./img/block.png"
	GameOverPNG = "./img/gameover.png"
	PausedPNG   = "./img/paused.png"
)

var (
	Blockbasher             *Game
	GamesConfig             *Config
	BgColor                 = color.White
	PaddleColor             = color.Black
	Red                     = color.RGBA{255, 99, 71, 0}
	TextColor               = color.RGBA{0, 0, 0, 0}
	DefaultWindowSizeWidth  = 480
	DefaultWindowSizeHeight = 320
)

type Config struct {
	WindowSize *Size
	PlayArea   ebiten.Image
	StatsArea  *Size // still need to figure this out
}

type Size struct {
	Width  int
	Height int
}

type ConfigOpts struct {
	X               float64
	Y               float64
	DirX            int
	DirY            int
	Hidden          bool
	Speed           float64
	SpeedMultiplier float64
}

func (opt *ConfigOpts) GetMovementSpeedX() float64 {
	return (opt.Speed * opt.SpeedMultiplier) * float64(opt.DirX)
}

func (opt *ConfigOpts) GetMovementSpeedY() float64 {
	return (opt.Speed * opt.SpeedMultiplier) * float64(opt.DirY)
}

func NewConfig() *Config {
	return &Config{
		WindowSize: &Size{
			Width:  DefaultWindowSizeWidth,
			Height: DefaultWindowSizeHeight,
		},
	}
}
