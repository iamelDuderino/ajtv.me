package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	Blockbasher.Start()
	Blockbasher.Logging = true
	log.Fatal(ebiten.RunGame(Blockbasher))
}
