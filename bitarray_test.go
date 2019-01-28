package sieve

import (
	"math"
	"testing"
)

func sieve(upto, cap int) Sieve {

	// design sieve to store only odd numbers
	sieve := NewSieve(upto/2, cap/2)
	uptoRoot := int(math.Sqrt(float64(upto)))

	// start off assuming all primes
	sieve.SetAll()
	sieve.Clear(0) // index 1 as sieve designed for odd numbers only
	// clears non primes
	for n := 3; n <= uptoRoot; n += 2 {
		if sieve.IsSet(n / 2) {
			sieve.ClearSeries(3*n/2, n)
		}
	}
	return sieve

}

func TestBitArray(t *testing.T) {

	tests := []struct {
		size int
		sol1 int
		sol2 int
	}{
		{1000, 168, 832},
	}

	for _, test := range tests {
		sieve := sieve(test.size, 2*test.size)
		if cnt := sieve.Count(); cnt != test.sol1 {
			t.Errorf("Sieve.Count ERROR ;; Expected:%d, Got:%d\n", test.sol1, cnt)
			t.Errorf("%v\n", sieve.ToNums())
		}
		sieve.ToggleAll()
		if cnt := sieve.Count(); cnt != test.sol2 {
			t.Errorf("Sieve.Count ERROR ;; Expected:%d, Got:%d\n", test.sol2, cnt)
			t.Errorf("%v\n", sieve.ToNums())
		}
	}

}
