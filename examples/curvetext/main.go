package main

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/setanarut/gog"
	"golang.org/x/image/font/gofont/gomono"
)

var (
	roboto        *text.GoTextFace
	letters       []text.Glyph
	screenSize              = image.Point{800, 400}
	dio                     = &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
	path          *gog.Path = gog.Lemniscate(300, 348)
	txt                     = "A dead simple 2D game engine for Go"
	tick, pathlen float64
)

type Game struct{}

func main() {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		log.Fatal(err)
	}
	roboto = &text.GoTextFace{Source: src, Size: 50}
	path.Translate(400, 200).Reverse().Close()
	pathlen = path.Length()
	dio.ColorScale.ScaleWithColor(color.RGBA{0, 195, 255, 255})
	letters = text.AppendGlyphs(letters, txt, roboto, &text.LayoutOptions{})
	ebiten.SetWindowSize(screenSize.X, screenSize.Y)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{DisableHiDPI: true})
}

func (g *Game) Update() error {
	if tick > pathlen {
		tick = 0
	}
	tick += 4
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{13, 17, 23, 255})

	for _, glyph := range letters {
		if glyph.Image != nil {
			l := (glyph.X + float64(glyph.Image.Bounds().Dx())/2)
			l += tick
			if l > pathlen {
				l = l - pathlen
			}
			point, angle := path.PointAngleAtLength(l)
			dio.GeoM.Reset()
			dio.GeoM.Translate(-float64(glyph.Image.Bounds().Dx())/2, -(glyph.OriginY - glyph.Y))
			dio.GeoM.Rotate(angle)
			dio.GeoM.Translate(point.X, point.Y)
			screen.DrawImage(glyph.Image, dio)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize.X, screenSize.Y
}
