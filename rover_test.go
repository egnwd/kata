package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Move Forward

func TestForwardNorth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: North}

	r.Command("F")
	assert.Equal(t, Rover{x: 10, y: 11, d: North}, r)
}

func TestForwardEast(t *testing.T) {
	r := Rover{x: 10, y: 10, d: East}

	r.Command("F")
	assert.Equal(t, Rover{x: 11, y: 10, d: East}, r)
}

func TestForwardSouth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: South}

	r.Command("F")
	assert.Equal(t, Rover{x: 10, y: 9, d: South}, r)
}

func TestForwardWest(t *testing.T) {
	r := Rover{x: 10, y: 10, d: West}

	r.Command("F")
	assert.Equal(t, Rover{x: 9, y: 10, d: West}, r)
}

// Move Backward

func TestBackwardNorth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: North}

	r.Command("B")
	assert.Equal(t, Rover{x: 10, y: 9, d: North}, r)
}

func TestBackwardEast(t *testing.T) {
	r := Rover{x: 10, y: 10, d: East}

	r.Command("B")
	assert.Equal(t, Rover{x: 9, y: 10, d: East}, r)
}

func TestBackwardSouth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: South}

	r.Command("B")
	assert.Equal(t, Rover{x: 10, y: 11, d: South}, r)
}

func TestBackwardWest(t *testing.T) {
	r := Rover{x: 10, y: 10, d: West}

	r.Command("B")
	assert.Equal(t, Rover{x: 11, y: 10, d: West}, r)
}

// Rotate Left

func TestLeft(t *testing.T) {
	r := Rover{x: 10, y: 10, d: North}

	r.Command("L")
	assert.Equal(t, Rover{x: 10, y: 10, d: West}, r)

	r.Command("L")
	assert.Equal(t, Rover{x: 10, y: 10, d: South}, r)

	r.Command("L")
	assert.Equal(t, Rover{x: 10, y: 10, d: East}, r)

	r.Command("L")
	assert.Equal(t, Rover{x: 10, y: 10, d: North}, r)
}

// Rotate Right

func TestRight(t *testing.T) {
	r := Rover{x: 10, y: 10, d: North}

	r.Command("R")
	assert.Equal(t, Rover{x: 10, y: 10, d: East}, r)

	r.Command("R")
	assert.Equal(t, Rover{x: 10, y: 10, d: South}, r)

	r.Command("R")
	assert.Equal(t, Rover{x: 10, y: 10, d: West}, r)

	r.Command("R")
	assert.Equal(t, Rover{x: 10, y: 10, d: North}, r)
}
