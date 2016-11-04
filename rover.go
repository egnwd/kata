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

	cardinalDirections = 4
)

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
	x int
	y int
	d int
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
	switch r.d {
	case North:
		r.y = (r.y + 1) % GridSize
	case East:
		r.x = (r.x + 1) % GridSize
	case South:
		r.y = (r.y + GridSize - 1) % GridSize
	case West:
		r.x = (r.x + GridSize - 1) % GridSize
	}
}

func (r *Rover) backward() {
	switch r.d {
	case North:
		r.y = (r.y + GridSize - 1) % GridSize
	case East:
		r.x = (r.x + GridSize - 1) % GridSize
	case South:
		r.y = (r.y + 1) % GridSize
	case West:
		r.x = (r.x + 1) % GridSize
	}
}

func (r *Rover) left() {
	r.d = (r.d + cardinalDirections - 1) % cardinalDirections
}

func (r *Rover) right() {
	r.d = (r.d + 1) % cardinalDirections
}

func (r *Rover) detectObstacle(isForward bool) bool {
	ghost := *r
	if isForward {
		ghost.forward()
	} else {
		ghost.backward()
	}

	h := hashCoordinate(ghost.x, ghost.y)

	return pluto[h]
}

func hashCoordinate(x, y int) int {
	// This is the cantor pairing function
	return int(0.5*float64((x+y)*(x+y+1))) + y
}
