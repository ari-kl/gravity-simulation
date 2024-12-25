package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/rand"
)

type Mass struct {
	mass         float32
	x, y         float32
	velx, vely   float32
	acclx, accly float32
	colour       color.RGBA
	pathpos      [][]float32
}

func NewMass(mass float32, x, y float32) *Mass {
	colour := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	return &Mass{mass, x, y, 0, 0, 0, 0, colour, [][]float32{}}
}

func (m *Mass) Update() {
	m.velx += m.acclx
	m.vely += m.accly

	m.x += m.velx
	m.y += m.vely

	m.pathpos = append(m.pathpos, []float32{m.x, m.y})
	if len(m.pathpos) > 50 {
		m.pathpos = m.pathpos[1:]
	}
}

func GStrengthAt(x, y float32, game *Game) (float32, float32) {
	masses := game.masses
	var total_forcex, total_forcey float32

	for _, m := range masses {
		dx := m.x - x
		dy := m.y - y
		r_squared := dx*dx + dy*dy
		force := (game.G * m.mass) / r_squared
		forcex := force * dx
		forcey := force * dy
		total_forcex += forcex
		total_forcey += forcey
	}

	return total_forcex, total_forcey
}

func (m *Mass) ApplyGravitation(game *Game) {
	masses := game.masses
	var total_forcex, total_forcey float32

	for _, other := range masses {
		if other != m {
			dx := other.x - m.x
			dy := other.y - m.y
			r_squared := dx*dx + dy*dy

			// Prevent division by zero so the simulation doesn't break
			if r_squared == 0 {
				r_squared = 0.0001
			}

			force := (game.G * m.mass * other.mass) / r_squared
			forcex := force * dx
			forcey := force * dy
			total_forcex += forcex
			total_forcey += forcey
		}
	}

	m.acclx = total_forcex / m.mass
	m.accly = total_forcey / m.mass
}

func (m *Mass) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, m.x, m.y, m.mass*4, m.colour, true)

	for i := 0; i < len(m.pathpos)-1; i++ {
		vector.StrokeLine(screen, m.pathpos[i][0], m.pathpos[i][1], m.pathpos[i+1][0], m.pathpos[i+1][1], m.mass*2, m.colour, true)
	}
}
