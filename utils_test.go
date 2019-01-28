package sieve

import "testing"

func TestPrimeSieveUpto(t *testing.T) {

	tests := []struct {
		upto int
		sol  int
	}{
		// note " - 1 " because 2 is not saved in the sieve
		// {1e2, 25-1},
		// {1e3, 168-1},
		// {1e4, 1229 - 1},
		// {1e5, 9592-1},
		// {1e6, 78498-1},
		// {1e7, 664579 - 1}, // 0.6s
		{1e8, 5761455 - 1}, // 1.2s
		// {1e9, 50847534-1},
	}

	for _, test := range tests {
		sieve := PrimeSieveUpto(test.upto)
		if cnt := sieve.Count(); cnt != test.sol {
			t.Errorf("Sieve.Count ERROR ;; Expected:%d, Got:%d \n", test.sol, cnt)
		}
	}

}

func TestPrimesUpto(t *testing.T) {

	tests := []struct {
		upto int
		sol  int
	}{
		{1e2, 25},
		// {1000, 168},
		// {10000, 1229},
		// {100000, 9592},
		// {1000000, 78498},
		// {10000000, 664579}, // 0.6s
		// {100000000, 5761455}, // 1.2s
		// {1000000000, 50847534},
	}

	for _, test := range tests {
		ints := PrimesUpto(test.upto)
		if len(ints) != test.sol {
			t.Errorf("Sieve.PrimesUpto ERROR ;; Expected:%d, Got:%d \n  %v\n", test.sol, len(ints), ints)
		}
	}

}

func TestFactorsUpto(t *testing.T) {

	tests := []struct {
		upto int
		sol  int
	}{
		{100, 4},
		{270, 6},
		{10001, 4},
		{1024, 2},
		{30030, 12},
	}

	factors := FactorsUpto(100000)

	for _, test := range tests {
		ints := factors(test.upto)
		if len(ints) != test.sol {
			t.Errorf("Sieve.FactorsUpto ERROR ;; Expected:%d, Got:%d \n  %v\n", test.sol, len(ints), ints)
		}
	}

}
