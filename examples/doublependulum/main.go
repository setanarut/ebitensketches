package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	height = 512
	width  = 512
)

type Game struct {
	angle1, angle2                     float64
	angularVelocity1, angularVelocity2 float64
	length1, length2                   float64
	mass1, mass2                       float64
	gravity                            float64
	dt                                 float64
}

func (g *Game) Update() error {
	// Handle pendulum physics
	numerator1 := -g.gravity * (2*g.mass1 + g.mass2) * math.Sin(g.angle1)
	numerator2 := -g.mass2 * g.gravity * math.Sin(g.angle1-2*g.angle2)
	numerator3 := -2 * math.Sin(g.angle1-g.angle2) * g.mass2
	numerator4 := g.angularVelocity2*g.angularVelocity2*g.length2 + g.angularVelocity1*g.angularVelocity1*g.length1*math.Cos(g.angle1-g.angle2)
	denominator := g.length1 * (2*g.mass1 + g.mass2 - g.mass2*math.Cos(2*g.angle1-2*g.angle2))
	angularAcceleration1 := (numerator1 + numerator2 + numerator3*numerator4) / denominator

	numerator1 = 2 * math.Sin(g.angle1-g.angle2)
	numerator2 = g.angularVelocity1 * g.angularVelocity1 * g.length1 * (g.mass1 + g.mass2)
	numerator3 = g.gravity * (g.mass1 + g.mass2) * math.Cos(g.angle1)
	numerator4 = g.angularVelocity2 * g.angularVelocity2 * g.length2 * g.mass2 * math.Cos(g.angle1-g.angle2)
	denominator = g.length2 * (2*g.mass1 + g.mass2 - g.mass2*math.Cos(2*g.angle1-2*g.angle2))
	angularAcceleration2 := (numerator1 * (numerator2 + numerator3 + numerator4)) / denominator

	g.angularVelocity1 += angularAcceleration1 * g.dt
	g.angularVelocity2 += angularAcceleration2 * g.dt
	g.angle1 += g.angularVelocity1 * g.dt
	g.angle2 += g.angularVelocity2 * g.dt

	// Normalize angles
	g.angle1 = math.Mod(g.angle1, 2*math.Pi)
	g.angle2 = math.Mod(g.angle2, 2*math.Pi)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.Gray{30})

	// Calculate pendulum positions
	x1 := g.length1 * math.Sin(g.angle1)
	y1 := g.length1 * math.Cos(g.angle1)
	x2 := x1 + g.length2*math.Sin(g.angle2)
	y2 := y1 + g.length2*math.Cos(g.angle2)

	// Draw pendulum rods using vector.StrokeLine
	vector.StrokeLine(screen, float32(width/2), float32(height/2), float32(width/2+x1), float32(height/2+y1), 2, color.White, true)
	vector.StrokeLine(screen, float32(width/2+x1), float32(height/2+y1), float32(width/2+x2), float32(height/2+y2), 2, color.White, true)

	// Draw pendulum masses using vector.DrawFilledCircle
	vector.DrawFilledCircle(screen, float32(width/2+x1), float32(height/2+y1), 10, color.RGBA{255, 125, 220, 255}, true)
	vector.DrawFilledCircle(screen, float32(width/2+x2), float32(height/2+y2), 10, color.RGBA{255, 125, 220, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	game := &Game{
		angle1:           math.Pi * 1.2,
		angle2:           math.Pi / 3,
		angularVelocity1: 0,
		angularVelocity2: 0,
		length1:          100,
		length2:          100,
		mass1:            1,
		mass2:            1,
		gravity:          5,
		dt:               0.2,
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Double Pendulum")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
