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

// Rover represents the state of a Kata rover
type Rover struct {
	x int
	y int
	d int
}

// Move takes in a set of commands for the rover to perform
func (r Rover) Move(c string) {

}
