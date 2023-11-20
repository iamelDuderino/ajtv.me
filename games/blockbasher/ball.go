package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	Img        *ebiten.Image
	ImgOpts    *ebiten.DrawImageOptions
	ConfigOpts *ConfigOpts
	IsMoving   bool // checks for a dead ball
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(ball.Img, ball.ImgOpts)
}

func (ball *Ball) SwitchDirX() {
	if ball.ConfigOpts.DirX == 1 {
		ball.ConfigOpts.DirX = -1
		return
	}
	ball.ConfigOpts.DirX = 1
}

func (ball *Ball) SwitchDirY() {
	if ball.ConfigOpts.DirY == 1 {
		ball.ConfigOpts.DirY = -1
		return
	}
	ball.ConfigOpts.DirY = 1
}

func (ball *Ball) EdgeDetectionRight() bool {
	return ball.ConfigOpts.X > float64(DefaultWindowSizeWidth)-10
}

func (ball *Ball) EdgeDetectionLeft() bool {
	return ball.ConfigOpts.X < 10
}

func (ball *Ball) EdgeDetectionTop() bool {
	return ball.ConfigOpts.Y < 10
}

func (ball *Ball) EdgeDetectionBottom() bool {
	return ball.ConfigOpts.Y > float64(DefaultWindowSizeHeight)-10
}

func (ball *Ball) Move() {
	var (
		geo    = ball.ImgOpts.GeoM.String()
		x      = ball.ImgOpts.GeoM.Element(0, 2)
		y      = ball.ImgOpts.GeoM.Element(1, 2)
		tx, ty float64

		colPaddle, colRight, colLeft, colTop, colBottom bool
	)
	if Blockbasher.Logging {
		log.Print("Ball GeoM:", geo)
	}

	ball.ConfigOpts.X = x
	ball.ConfigOpts.Y = y

	colRight = ball.EdgeDetectionRight()
	colLeft = ball.EdgeDetectionLeft()
	colBottom = ball.EdgeDetectionBottom()
	colTop = ball.EdgeDetectionTop()

	if colBottom {
		ball.IsMoving = false
		Blockbasher.Player.Lives -= 1
		return
	}

	colPaddle = Blockbasher.CollisionDetection(ball.Hitbox(), Blockbasher.Paddle.Hitbox())

	if !colRight && !colLeft && !colTop && !colBottom {
		colTop, colBottom, colLeft, colRight = Blockbasher.Level.BlockLayout.CollisionDetection()
	}

	if colPaddle || colTop || colBottom {
		ball.SwitchDirY()
	}
	if colRight || colLeft {
		ball.SwitchDirX()
	}

	if Blockbasher.Logging {
		if colTop {
			log.Println("COL TOP")
		}
		if colBottom {
			log.Println("COL BOTTOM")
		}
		if colLeft {
			log.Println("COL LEFT")
		}
		if colRight {
			log.Println("COL RIGHT")
		}
	}

	ty = ball.ConfigOpts.GetMovementSpeedY()
	tx = ball.ConfigOpts.GetMovementSpeedX()

	// Draw New X Position
	ball.ImgOpts.GeoM.Translate(tx, ty)

}

func (ball *Ball) Hitbox() image.Rectangle {
	p := image.Point{int(ball.ConfigOpts.X), int(ball.ConfigOpts.Y)}
	return ball.Img.Bounds().Add(p)
}

func (ball *Ball) SetDefaultXY() {
	ball.ImgOpts.GeoM.Translate(ball.ConfigOpts.X, ball.ConfigOpts.Y)
}

func NewBall() *Ball {
	img, _, _ := ebitenutil.NewImageFromFile(BallPNG)
	ball := &Ball{
		Img:     img,
		ImgOpts: &ebiten.DrawImageOptions{GeoM: ebiten.GeoM{}},
		ConfigOpts: &ConfigOpts{
			X:               10,
			Y:               10,
			DirX:            1,
			DirY:            1,
			Hidden:          false,
			Speed:           2,
			SpeedMultiplier: 1,
		},
		IsMoving: true,
	}
	ball.SetDefaultXY()
	return ball
}
