package main

type RingBuffer struct {
	data   []Position
	Length int
	idx    int
}

func NewRingBuffer(config GameConfig) RingBuffer {
	var capacity int
	if config.XSize*config.YSize < config.MaxLength {
		capacity = config.XSize * config.YSize
	} else {
		capacity = config.MaxLength
	}

	data := make([]Position, capacity)
	return RingBuffer{
		data:   data,
		Length: 1,
		idx:    0,
	}
}

// Expand the ring buffer length by one until it matches capacity.
// If it can't expand any more, do not throw an error. Silently continue.
func (rb *RingBuffer) Expand() {
	if rb.Length < len(rb.data) {
		rb.Length += 1
	}
}

func (rb *RingBuffer) Add(pos Position) {
	i := (rb.idx + 1) % rb.Length
	rb.data[i] = pos
	rb.idx = i
}

func (rb *RingBuffer) Contains(pos Position) bool {
	for _, p := range rb.data {
		if p == pos {
			return true
		}
	}
	return false
}
