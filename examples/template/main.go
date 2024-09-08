package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	w, h int = 800, 400
)

type Game struct{}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{DisableHiDPI: true})
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{13, 17, 23, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}
