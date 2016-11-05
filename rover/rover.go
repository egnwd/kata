package rover

import (
	"fmt"
	"math/rand"
)

const (
	// North is a compass point
	North = iota
	// East is a compass point
	East
	// South is a compass point
	South
	// West is a compass point
	West
	cardinalDirections
)

var compassPoints = []string{"North", "East", "South", "West"}

// Action constants
const (
	forward  = 'F'
	backward = 'B'
	left     = 'L'
	right    = 'R'
)

// GridSize represents the size of the grid on Pluto (NxN)
const GridSize = 100

// Rover represents the state of a Kata rover
type Rover struct {
	X int
	Y int
	D int
}

const maxObstacles = 50

var pluto = make(map[int]bool, maxObstacles)

func init() {
	for i := 0; i < maxObstacles; i++ {
		x := rand.Intn(GridSize)
		y := rand.Intn(GridSize)
		h := hashCoordinate(x, y)
		pluto[h] = true
	}
}

// Command takes in a set of commands for the rover to perform
func (r *Rover) Command(cs string) {
	for _, c := range cs {
		switch c {
		case forward, backward:
			if r.detectObstacle(c == forward) {
				fmt.Println("Danger: Obstacle Ahead")
				return
			}
			r.move(c)
		case left:
			r.left()
		case right:
			r.right()
		}
	}
}

func (r *Rover) move(c rune) {
	switch c {
	case forward:
		r.forward()
	case backward:
		r.backward()
	}
}

func (r *Rover) forward() {
	switch r.D {
	case North:
		r.Y = (r.Y + 1) % GridSize
	case East:
		r.X = (r.X + 1) % GridSize
	case South:
		r.Y = (r.Y + GridSize - 1) % GridSize
	case West:
		r.X = (r.X + GridSize - 1) % GridSize
	}
}

func (r *Rover) backward() {
	switch r.D {
	case North:
		r.Y = (r.Y + GridSize - 1) % GridSize
	case East:
		r.X = (r.X + GridSize - 1) % GridSize
	case South:
		r.Y = (r.Y + 1) % GridSize
	case West:
		r.X = (r.X + 1) % GridSize
	}
}

func (r *Rover) left() {
	r.D = (r.D + cardinalDirections - 1) % cardinalDirections
}

func (r *Rover) right() {
	r.D = (r.D + 1) % cardinalDirections
}

func (r *Rover) detectObstacle(isForward bool) bool {
	ghost := *r
	if isForward {
		ghost.forward()
	} else {
		ghost.backward()
	}

	h := hashCoordinate(ghost.X, ghost.Y)

	return pluto[h]
}

func (r *Rover) String() string {
	return fmt.Sprintf("Rover: (%d, %d), facing %s", r.X, r.Y, compassPoints[r.D])
}

func hashCoordinate(x, y int) int {
	// This is the cantor pairing function
	return int(0.5*float64((x+y)*(x+y+1))) + y
}
