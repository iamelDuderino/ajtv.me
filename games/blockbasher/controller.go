package main

import "github.com/hajimehoshi/ebiten/v2"

type Controller struct {
	Left   ebiten.Key
	Right  ebiten.Key
	Up     ebiten.Key
	Down   ebiten.Key
	Start  ebiten.Key
	Action ebiten.Key
	Back   ebiten.Key
	Quit   ebiten.Key
}

func (controller *Controller) UpdateLeft(key ebiten.Key) {
	controller.Left = key
}

func (controller *Controller) UpdateRight(key ebiten.Key) {
	controller.Right = key
}

func (controller *Controller) UpdateUp(key ebiten.Key) {
	controller.Up = key
}

func (controller *Controller) UpdateDown(key ebiten.Key) {
	controller.Down = key
}

func NewController() *Controller {
	return &Controller{
		Left:   ebiten.KeyA,
		Right:  ebiten.KeyD,
		Up:     ebiten.KeyW,
		Down:   ebiten.KeyS,
		Start:  ebiten.KeyEnter,
		Action: ebiten.KeyE,
		Back:   ebiten.KeyQ,
		Quit:   ebiten.KeyEscape,
	}
}
