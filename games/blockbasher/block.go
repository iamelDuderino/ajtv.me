package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Block struct {
	Img        *ebiten.Image
	ImgOpt     *ebiten.DrawImageOptions
	ConfigOpts *ConfigOpts
	Strength   int
}

func (block *Block) Draw(screen *ebiten.Image) {
	screen.DrawImage(block.Img, block.ImgOpt)
}

func (block *Block) Hitbox() image.Rectangle {
	p := image.Point{int(block.ConfigOpts.X), int(block.ConfigOpts.Y)}
	return block.Img.Bounds().Add(p)
}

func NewBlock(x, y float64, strength int) *Block {
	img, _, _ := ebitenutil.NewImageFromFile(BlockPNG)
	block := &Block{
		Img:    img,
		ImgOpt: &ebiten.DrawImageOptions{GeoM: ebiten.GeoM{}},
		ConfigOpts: &ConfigOpts{
			X: x,
			Y: y,
		},
		Strength: strength,
	}
	block.ImgOpt.GeoM.Translate(block.ConfigOpts.X, block.ConfigOpts.Y)
	return block
}

type BlockLayout struct {
	Rows []*BlockRow
}

type BlockRow struct {
	Blocks []*Block
}

// func (blockLayout *BlockLayout) CollisionDetection() bool {
func (blockLayout *BlockLayout) CollisionDetection() (bool, bool, bool, bool) {
	for _, row := range Blockbasher.Level.BlockLayout.Rows {
		for _, block := range row.Blocks {
			if !block.ConfigOpts.Hidden {
				top, bottom, left, right := Blockbasher.CollisionDetectionAt(Blockbasher.Ball.Hitbox(), block.Hitbox())
				if top || bottom || left || right {
					block.ConfigOpts.Hidden = true
					Blockbasher.Player.Score += 1 // Update Player Score
					if Blockbasher.Logging {
						log.Println("BALL COLLIDING WITH BLOCK")
					}
					return top, bottom, left, right
				}
			}
		}
	}
	return false, false, false, false
}

func NewBlockLayout(strength int) *BlockLayout {
	var (
		layout = &BlockLayout{}
		n, nn  int
		x, y   float64 = 20, 10
		block  *Block
	)
	block = NewBlock(x, y, strength)
	blockSizeW, blockSizeH := float64(block.Hitbox().Dx()), float64(block.Hitbox().Dy())

	for n < 10 {
		nn = 0
		x = 20
		layout.Rows = append(layout.Rows, &BlockRow{})
		for nn < 11 {
			x += blockSizeW
			layout.Rows[n].Blocks = append(layout.Rows[n].Blocks, NewBlock(x, y, strength))
			nn += 1
		}
		y += blockSizeH
		n += 1
	}
	return layout
}
