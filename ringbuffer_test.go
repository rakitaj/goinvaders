package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = GameConfig{XSize: 10, YSize: 10, MaxLength: 8}

func TestCapacityIsTheSmallerNumer(t *testing.T) {
	ring_buffer := NewRingBuffer(config)
	assert.Equal(t, 8, len(ring_buffer.data), "Ring buffer capacity should be the smaller of MaxLength or board height * width.")
}

func TestCantExpandMoreThanCapacity(t *testing.T) {
	ring_buffer := NewRingBuffer(config)
	for i := 0; i < 8; i++ {
		ring_buffer.Expand()
	}
	assert.Equal(t, 8, ring_buffer.Length, "Ring buffer.Length must never exceed capacity.")
}

func TestFunctional(t *testing.T) {
	rb := NewRingBuffer(config)
	rb.Expand()
	assert.Equal(t, rb.Length, 2)
	rb.Add(Position{X: 0, Y: 0})
	rb.Add(Position{X: 1, Y: 1})
}
