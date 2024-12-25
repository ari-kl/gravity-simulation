package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	masses      []*Mass
	G           float32
	fieldLines  bool
	paused      bool
	currentMass int
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.currentMass = 1
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.currentMass = 2
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.currentMass = 3
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.currentMass = 4
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.currentMass = 5
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
		g.currentMass = 6
	} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
		g.currentMass = 7
	} else if inpututil.IsKeyJustPressed(ebiten.Key8) {
		g.currentMass = 8
	} else if inpututil.IsKeyJustPressed(ebiten.Key9) {
		g.currentMass = 9
	} else if inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.currentMass = 10
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.fieldLines = !g.fieldLines
	} else if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		g.masses = []*Mass{}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		g.G -= 0.05
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		g.G += 0.05
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.masses = append(g.masses, NewMass(float32(g.currentMass), float32(mx), float32(my)))
	}

	dx, dy := ebiten.Wheel()

	for _, m := range g.masses {
		m.x += float32(dx)
		m.y += float32(dy)

		for pos := range m.pathpos {
			m.pathpos[pos][0] += float32(dx)
			m.pathpos[pos][1] += float32(dy)
		}
	}

	if !g.paused {
		for _, m := range g.masses {
			m.ApplyGravitation(g)
			m.Update()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.fieldLines {
		arrowSpacing := 20

		for x := 0; x < screen.Bounds().Dx()+arrowSpacing; x += arrowSpacing {
			for y := 0; y < screen.Bounds().Dy()+arrowSpacing; y += arrowSpacing {
				sx, sy := GStrengthAt(float32(x), float32(y), g)

				// Normalize the vector (sx, sy)
				magnitude := float32(math.Sqrt(float64(sx*sx + sy*sy)))

				if magnitude > 0 {
					sx /= magnitude
					sy /= magnitude
				}

				lineColour := color.RGBA{0, 255, 0, 255}
				lineColour.G = uint8(math.Min(255, float64(magnitude*40*255)))

				vector.StrokeLine(screen, float32(x), float32(y), float32(x)+sx*10, float32(y)+sy*10, 1, lineColour, false)
				vector.DrawFilledCircle(screen, float32(x)+sx*10, float32(y)+sy*10, 2, lineColour, false)
			}
		}
	}

	for _, m := range g.masses {
		m.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Click to place masses\nPress number keys to set mass value, backspace to clear all masses, F to toggle field lines, Space to pause, -/+ to decrease/increase force strength\nCurrent Mass: %d\nG: %.2f", g.currentMass, g.G))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gravity!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{
		masses:      []*Mass{},
		G:           0.1,
		fieldLines:  true,
		paused:      false,
		currentMass: 1,
	}); err != nil {
		log.Fatal(err)
	}
}
