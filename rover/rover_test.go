package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutObstacles(t *testing.T) {
	// Setup pluto
	var noObstacles = map[int]bool{}
	pluto = noObstacles

	// Move Forward
	t.Run("ForwardNorth", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: North}

		r.Command("F")
		assert.Equal(t, Rover{X: 10, Y: 11, D: North}, r)
	})

	t.Run("ForwardEast", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: East}

		r.Command("F")
		assert.Equal(t, Rover{X: 11, Y: 10, D: East}, r)
	})

	t.Run("ForwardSouth", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: South}

		r.Command("F")
		assert.Equal(t, Rover{X: 10, Y: 9, D: South}, r)
	})

	t.Run("ForwardWest", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: West}

		r.Command("F")
		assert.Equal(t, Rover{X: 9, Y: 10, D: West}, r)
	})

	// Move Backward
	t.Run("BackwardNorth", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: North}

		r.Command("B")
		assert.Equal(t, Rover{X: 10, Y: 9, D: North}, r)
	})

	t.Run("BackwardEast", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: East}

		r.Command("B")
		assert.Equal(t, Rover{X: 9, Y: 10, D: East}, r)
	})

	t.Run("BackwardSouth", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: South}

		r.Command("B")
		assert.Equal(t, Rover{X: 10, Y: 11, D: South}, r)
	})

	t.Run("BackwardWest", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: West}

		r.Command("B")
		assert.Equal(t, Rover{X: 11, Y: 10, D: West}, r)
	})

	// Rotate Left

	t.Run("RotateLeft", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: North}

		r.Command("L")
		assert.Equal(t, Rover{X: 10, Y: 10, D: West}, r)

		r.Command("L")
		assert.Equal(t, Rover{X: 10, Y: 10, D: South}, r)

		r.Command("L")
		assert.Equal(t, Rover{X: 10, Y: 10, D: East}, r)

		r.Command("L")
		assert.Equal(t, Rover{X: 10, Y: 10, D: North}, r)
	})

	// Rotate Right

	t.Run("RotateRight", func(t *testing.T) {
		r := Rover{X: 10, Y: 10, D: North}

		r.Command("R")
		assert.Equal(t, Rover{X: 10, Y: 10, D: East}, r)

		r.Command("R")
		assert.Equal(t, Rover{X: 10, Y: 10, D: South}, r)

		r.Command("R")
		assert.Equal(t, Rover{X: 10, Y: 10, D: West}, r)

		r.Command("R")
		assert.Equal(t, Rover{X: 10, Y: 10, D: North}, r)
	})
}

func TestWithObstacles(t *testing.T) {
	// Setup pluto
	/*
	  R = Rover starting position
	  * = Obstacle
	  +---------+
	  | | | | | |
	  | | |*| | |
	  |*| | | | |
	  | | | |*| |
	  |R| | | | |
	  +---------+
	*/
	var obstacles = map[int]bool{
		5:  true, // (0, 2)
		11: true, // (3, 1)
		18: true, // (2, 3)
	}
	pluto = obstacles

	t.Run("ObstaclePass", func(t *testing.T) {
		r := Rover{X: 0, Y: 0, D: North}

		r.Command("FRFFLFRFLF")
		assert.Equal(t, Rover{X: 3, Y: 3, D: North}, r)
	})

	t.Run("ObstacleFail", func(t *testing.T) {
		r := Rover{X: 0, Y: 0, D: North}

		r.Command("FRFFFFLFF")
		assert.Equal(t, Rover{X: 2, Y: 1, D: East}, r)
	})
}

func TestString(t *testing.T) {
	r := Rover{X: 43, Y: 87, D: West}

	assert.Equal(t, "Rover: (43, 87), facing West", r.String())
}
