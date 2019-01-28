package sieve

import "fmt"

// PositionOutOfRangeError caused by trying to access a sieve past
// the end of its size
type PositionOutOfRangeError struct {
	position, size int
}

// Error returns a human readable description of the out-of-range error.
func (err PositionOutOfRangeError) Error() string {
	msg := "Sieve position out of range (Size:%d, Position:%d)"
	return fmt.Sprintf(msg, err.size, err.position)
}

// ResizeOutOfRangeError caused by trying to resize sieve above the
// available capacity
type ResizeOutOfRangeError struct {
	resize, capacity int
}

// Error returns a human readable description of the out-of-range error.
func (err ResizeOutOfRangeError) Error() string {
	msg := "Cannot resize above sieve capacity (Capacity:%d, Resize:%d)"
	return fmt.Sprintf(msg, err.capacity, err.resize)
}
