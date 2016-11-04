package rover

const (
	// North is a compass point
	North = iota
	// East is a compass point
	East
	// South is a compass point
	South
	// West is a compass point
	West
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

// Move takes in a set of commands for the rover to perform
func (r *Rover) Move(cs string) {
	for _, c := range cs {
		switch c {
		case forward:
			r.forward()
		}
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
