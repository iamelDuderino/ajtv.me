package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Scenes struct {
	MainMenu *Scene
	GameOver *Scene
	Paused   *Scene
}

type Scene struct {
	Img        *ebiten.Image
	ImgOpts    *ebiten.DrawImageOptions
	ConfigOpts *ConfigOpts
}

func (scene *Scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(scene.Img, scene.ImgOpts)
}

func (scene *Scene) SetDefaultXY() {
	scene.ImgOpts.GeoM.Translate(scene.ConfigOpts.X, scene.ConfigOpts.Y)
}

func NewMainMenuScene() *Scene {
	img, _, _ := ebitenutil.NewImageFromFile(MainMenuPNG)
	scene := &Scene{
		Img:     img,
		ImgOpts: &ebiten.DrawImageOptions{GeoM: ebiten.GeoM{}},
		ConfigOpts: &ConfigOpts{
			X: 0,
			Y: 0,
		},
	}
	scene.SetDefaultXY()
	return scene
}

func NewGameOverScene() *Scene {
	img, _, _ := ebitenutil.NewImageFromFile(GameOverPNG)
	scene := &Scene{
		Img:     img,
		ImgOpts: &ebiten.DrawImageOptions{GeoM: ebiten.GeoM{}},
		ConfigOpts: &ConfigOpts{
			X: 0,
			Y: 0,
		},
	}
	scene.SetDefaultXY()
	return scene
}

func NewPausedScene() *Scene {
	img, _, _ := ebitenutil.NewImageFromFile(PausedPNG)
	scene := &Scene{
		Img:     img,
		ImgOpts: &ebiten.DrawImageOptions{GeoM: ebiten.GeoM{}},
		ConfigOpts: &ConfigOpts{
			X:      float64(DefaultWindowSizeWidth)/2 - (75 / 2),
			Y:      float64(DefaultWindowSizeHeight)/2 - 17,
			Hidden: false,
		},
	}
	scene.SetDefaultXY()
	return scene
}
