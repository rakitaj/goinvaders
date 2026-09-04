package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = GameConfig{XSize: 10, YSize: 10, MaxLength: 8}
var starting_position = NewPosition(0, 0)

func TestCapacityIsTheSmallerNumer(t *testing.T) {
	ring_buffer := NewRingBuffer(config, starting_position)
	assert.Equal(t, 8, len(ring_buffer.data), "Ring buffer capacity should be the smaller of MaxLength or board height * width.")
}

func TestCantExpandMoreThanCapacity(t *testing.T) {
	ring_buffer := NewRingBuffer(config, starting_position)
	for i := 0; i < 10; i++ {
		ring_buffer.Expand()
	}
	assert.Equal(t, 8, ring_buffer.Length, "Ring buffer.Length must never exceed capacity.")
}

func TestFunctional(t *testing.T) {
	rb := NewRingBuffer(config, starting_position)
	rb.Expand()
	assert.Equal(t, rb.Length, 2)

	rb.Add(Position{X: 1, Y: 1})
	rb.Add(Position{X: 2, Y: 2})
	rb.Add(Position{X: 3, Y: 3})
	assert.Equal(t, rb.Length, 2)

	assert.Equal(t, rb.Get(0), NewPosition(2, 2))
	assert.Equal(t, rb.Get(1), NewPosition(3, 3))
}

func TestExpand(t *testing.T) {
	rb := NewRingBuffer(config, starting_position)
	rb.Expand()
	rb.Expand()
	assert.Equal(t, rb.Length, 3)

	rb.Add(Position{X: 1, Y: 1})
	rb.Add(Position{X: 2, Y: 2})
	rb.Add(Position{X: 3, Y: 3})

	rb.Expand()

	rb.Add(Position{X: 4, Y: 4})

	assert.Equal(t, rb.data[0], NewPosition(1, 1))
	assert.Equal(t, rb.data[1], NewPosition(2, 2))
	assert.Equal(t, rb.data[2], NewPosition(3, 3))
	assert.Equal(t, rb.data[3], NewPosition(4, 4))
}
