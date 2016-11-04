package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForwardNorth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: North}

	r.Move("F")
	assert.Equal(t, Rover{x: 10, y: 11, d: North}, r)
}

func TestForwardEast(t *testing.T) {
	r := Rover{x: 10, y: 10, d: East}

	r.Move("F")
	assert.Equal(t, Rover{x: 11, y: 10, d: East}, r)
}

func TestForwardSouth(t *testing.T) {
	r := Rover{x: 10, y: 10, d: South}

	r.Move("F")
	assert.Equal(t, Rover{x: 10, y: 9, d: South}, r)
}

func TestForwardWest(t *testing.T) {
	r := Rover{x: 10, y: 10, d: West}

	r.Move("F")
	assert.Equal(t, Rover{x: 9, y: 10, d: West}, r)
}
