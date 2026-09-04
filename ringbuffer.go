package main

type RingBuffer struct {
	data   []Position
	head   int
	Length int
}

func NewRingBuffer(config GameConfig, starting_position Position) RingBuffer {
	var capacity int
	if config.XSize*config.YSize < config.MaxLength {
		capacity = config.XSize * config.YSize
	} else {
		capacity = config.MaxLength
	}

	data := make([]Position, capacity)
	data[0] = starting_position

	return RingBuffer{
		data:   data,
		head:   0,
		Length: 1,
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
	rb.data[rb.head] = pos
	rb.head = (rb.head + 1) % len(rb.data)
}

func (rb *RingBuffer) Get(idx int) Position {
	i := (rb.head - rb.Length + idx) % len(rb.data)
	return rb.data[i]
}

func (rb *RingBuffer) Contains(pos Position) bool {
	// Start at head minus the num of active data elements, accounting for wrapping.
	i := (rb.head - rb.Length) % len(rb.data)
	// End at idx, accounting for wrapping.
	for i != rb.head {
		if rb.data[i] == pos {
			return true
		}
		i = (i + 1) % len(rb.data)
	}
	return false
}

// func (rb *RingBuffer) Live() []Position {
// 	// Start at head minus the num of active data elements, accounting for wrapping.
// 	i := (rb.head - rb.Length) % len(rb.data)
// 	live_positions := make([]Position, 0, rb.Length)
// 	// End at idx, accounting for wrapping.
// 	for i != rb.head {
// 		live_positions = append(live_positions, rb.data[i])
// 	}
// 	return live_positions
// }
