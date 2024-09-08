package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 640
	screenHeight = 480
	length1      = 100.0 // Length of the first rod
	length2      = 50.0  // Length of the second rod
	mass1        = 100.0 // Mass of the first pendulum
	mass2        = 100.0 // Mass of the second pendulum
	gravity      = 4.    // Gravitational acceleration
	dt           = 1.0 / 60.0
)

// Pendulum structure to hold state
type Pendulum struct {
	space  *cp.Space
	body1  *cp.Body
	body2  *cp.Body
	joint1 *cp.Constraint
	joint2 *cp.Constraint
}

// Create a new pendulum with Chipmunk physics engine
func newPendulum() *Pendulum {
	space := cp.NewSpace()
	space.SetGravity(cp.Vector{X: 0, Y: -gravity})

	// Create bodies for pendulum
	body1 := cp.NewBody(mass1, cp.MomentForSegment(mass1, cp.Vector{0, 0}, cp.Vector{0, length1}, 0))
	body1.SetPosition(cp.Vector{X: 320, Y: 240}) // Set to center of screen
	body1.SetAngle(math.Pi / 4)                  // Start at an angle (45 degrees)
	space.AddBody(body1)

	body2 := cp.NewBody(mass2, cp.MomentForSegment(mass2, cp.Vector{0, 0}, cp.Vector{0, length2}, 0))
	body2.SetPosition(cp.Vector{X: 320, Y: 240 - length1})
	space.AddBody(body2)

	// Create joints to simulate pendulum arms
	joint1 := cp.NewPinJoint(space.StaticBody, body1, cp.Vector{X: 320, Y: 240}, cp.Vector{X: 0, Y: 0})
	space.AddConstraint(joint1)

	joint2 := cp.NewPinJoint(body1, body2, cp.Vector{X: 0, Y: -length1}, cp.Vector{X: 0, Y: 0})
	space.AddConstraint(joint2)
	return &Pendulum{
		space:  space,
		body1:  body1,
		body2:  body2,
		joint1: joint1,
		joint2: joint2,
	}
}

// Step the physics engine forward
func (p *Pendulum) update() {
	p.space.Step(dt)
}

// Get the positions of the two masses
func (p *Pendulum) positions() (cp.Vector, cp.Vector) {
	pos1 := p.body1.Position()
	pos2 := p.body2.Position()
	return pos1, pos2
}

// Game struct to handle Ebiten state
type Game struct {
	pendulum *Pendulum
}

// Update the game state each frame
func (g *Game) Update() error {
	g.pendulum.update()
	return nil
}

// Draw the pendulum using vector.StrokeLine for thicker rods
func (g *Game) Draw(screen *ebiten.Image) {
	// Get positions of the pendulum masses
	pos1, pos2 := g.pendulum.positions()

	// Define the line thickness
	lineWidth := float32(4.0)

	// First rod (from the pivot point to the first mass)
	vector.StrokeLine(screen, 320, 240, float32(pos1.X), float32(pos1.Y), lineWidth, colornames.Red, false)

	// Second rod (from the first mass to the second mass)
	vector.StrokeLine(screen, float32(pos1.X), float32(pos1.Y), float32(pos2.X), float32(pos2.Y), lineWidth, colornames.Green, false)
}

// Layout sets the screen dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Main function to run the game
func main() {
	pendulum := newPendulum()

	game := &Game{
		pendulum: pendulum,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Double Pendulum Simulation with StrokeLine")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
