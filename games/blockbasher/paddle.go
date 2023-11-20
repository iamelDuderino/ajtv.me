package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Paddle struct {
	Img        *ebiten.Image
	ImgOpts    *ebiten.DrawImageOptions
	ConfigOpts *ConfigOpts
}

func (paddle *Paddle) Draw(screen *ebiten.Image) {
	if paddle.ConfigOpts.Hidden {
		return
	}
	screen.DrawImage(paddle.Img, paddle.ImgOpts)
}

func (paddle *Paddle) SetDefaultXY() {
	paddle.ImgOpts.GeoM.Translate(paddle.ConfigOpts.X, paddle.ConfigOpts.Y)
}

func (paddle *Paddle) Hitbox() image.Rectangle {
	xy := image.Point{int(Blockbasher.Paddle.ConfigOpts.X), int(Blockbasher.Paddle.ConfigOpts.Y)}
	return Blockbasher.Paddle.Img.Bounds().Add(xy)
}

func (paddle *Paddle) Move() {
	var (
		geo = paddle.ImgOpts.GeoM.String()
		x   = paddle.ImgOpts.GeoM.Element(0, 2)
		y   = paddle.ImgOpts.GeoM.Element(1, 2)
		tx  float64
		ty  float64 = 0
	)
	if Blockbasher.Logging {
		log.Print("Paddle GeoM:", geo)
	}

	paddle.ConfigOpts.X = x
	paddle.ConfigOpts.Y = y

	switch paddle.ConfigOpts.DirX {
	// Check For Right Boundary
	case 1:
		if x <= float64(DefaultWindowSizeWidth)-115 {
			tx = paddle.ConfigOpts.GetMovementSpeedX()
		}
	// Check For Left Boundary
	case -1:
		if x >= 5 {
			tx = paddle.ConfigOpts.GetMovementSpeedX()
		}
	default:
		tx = 0
	}
	paddle.ImgOpts.GeoM.Translate(tx, ty)
}

// checks for Left <A> or Right <D> to be held by Controller
func (paddle *Paddle) IsMoving() bool {
	switch {
	case ebiten.IsKeyPressed(Blockbasher.Controller.Left):
		paddle.ConfigOpts.DirX = -1
		return true
	case ebiten.IsKeyPressed(Blockbasher.Controller.Right):
		paddle.ConfigOpts.DirX = 1
		return true
	default:
		paddle.ConfigOpts.DirX = 0
		paddle.ConfigOpts.DirY = 0
		return false
	}
}

func NewPaddle() *Paddle {
	img, _, _ := ebitenutil.NewImageFromFile(PaddlePNG)
	paddle := &Paddle{
		Img: img,
		ImgOpts: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		}, ConfigOpts: &ConfigOpts{
			X:               float64((DefaultWindowSizeWidth / 2) - 50),
			Y:               float64(DefaultWindowSizeHeight - 20),
			DirX:            0,
			DirY:            0,
			Speed:           5,
			SpeedMultiplier: 1,
			Hidden:          false,
		},
	}
	paddle.SetDefaultXY()
	return paddle
}
