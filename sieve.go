package sieve

// Sieve defines the methods manipulate as standard sieve
type Sieve interface {

	// Set sets the @index position within sieve.
	// throws @PositionOutOfRangeError if @index >= Size()
	Set(index int) error

	// Clear clears the @index position within sieve.
	// throws @PositionOutOfRangeError if @index >= Size()
	Clear(index int) error

	// Toggle toggles the @index position value within sieve.
	// throws @PositionOutOfRangeError if @index >= Size()
	Toggle(index int) error

	// SetAll sets all positions within sieve.
	SetAll()

	// ClearAll clears all positions within sieve.
	ClearAll()

	// ToggleAll toggles all positions within sieve.
	ToggleAll()

	// IsSet returns true if @index position with sieve is set.
	IsSet(index int) bool

	// Count returns the count of all set positions within sieve.
	Count() int

	// Capacity returns the total available capacity within sieve.
	Capacity() int

	// Size returns the current size of sieve.
	Size() int

	// Resize changes the size of sieve but *only* within the capacity.
	// Throws ResizeOutOfRangeError if @size >= Capacity()
	// Resize will clear the sieve.
	Resize(size int) error

	// ToNums returns indexes of sieve that is set as an array.
	ToNums() []int

	// SubsetOf returns true if the sieve is a subset of @super.
	// SubsetOf(super Sieve) bool

	// SetSeries sets all positions that satisfy the arithmetic series upto
	// the size of sieve.
	SetSeries(startIndex, stepIndex int)

	// ClearSeries clears all positions that satisfy the arithmetic series upto
	// the size of sieve.
	ClearSeries(startIndex, stepIndex int)
}
