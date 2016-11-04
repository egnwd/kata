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
		r := Rover{x: 10, y: 10, d: North}

		r.Command("F")
		assert.Equal(t, Rover{x: 10, y: 11, d: North}, r)
	})

	t.Run("ForwardEast", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: East}

		r.Command("F")
		assert.Equal(t, Rover{x: 11, y: 10, d: East}, r)
	})

	t.Run("ForwardSouth", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: South}

		r.Command("F")
		assert.Equal(t, Rover{x: 10, y: 9, d: South}, r)
	})

	t.Run("ForwardWest", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: West}

		r.Command("F")
		assert.Equal(t, Rover{x: 9, y: 10, d: West}, r)
	})

	// Move Backward
	t.Run("BackwardNorth", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: North}

		r.Command("B")
		assert.Equal(t, Rover{x: 10, y: 9, d: North}, r)
	})

	t.Run("BackwardEast", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: East}

		r.Command("B")
		assert.Equal(t, Rover{x: 9, y: 10, d: East}, r)
	})

	t.Run("BackwardSouth", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: South}

		r.Command("B")
		assert.Equal(t, Rover{x: 10, y: 11, d: South}, r)
	})

	t.Run("BackwardWest", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: West}

		r.Command("B")
		assert.Equal(t, Rover{x: 11, y: 10, d: West}, r)
	})

	// Rotate Left

	t.Run("RotateLeft", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: North}

		r.Command("L")
		assert.Equal(t, Rover{x: 10, y: 10, d: West}, r)

		r.Command("L")
		assert.Equal(t, Rover{x: 10, y: 10, d: South}, r)

		r.Command("L")
		assert.Equal(t, Rover{x: 10, y: 10, d: East}, r)

		r.Command("L")
		assert.Equal(t, Rover{x: 10, y: 10, d: North}, r)
	})

	// Rotate Right

	t.Run("RotateRight", func(t *testing.T) {
		r := Rover{x: 10, y: 10, d: North}

		r.Command("R")
		assert.Equal(t, Rover{x: 10, y: 10, d: East}, r)

		r.Command("R")
		assert.Equal(t, Rover{x: 10, y: 10, d: South}, r)

		r.Command("R")
		assert.Equal(t, Rover{x: 10, y: 10, d: West}, r)

		r.Command("R")
		assert.Equal(t, Rover{x: 10, y: 10, d: North}, r)
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
		r := Rover{x: 0, y: 0, d: North}

		r.Command("FRFFLFRFLF")
		assert.Equal(t, Rover{x: 3, y: 3, d: North}, r)
	})

	t.Run("ObstacleFail", func(t *testing.T) {
		r := Rover{x: 0, y: 0, d: North}

		r.Command("FRFFFFLFF")
		assert.Equal(t, Rover{x: 2, y: 1, d: East}, r)
	})
}
