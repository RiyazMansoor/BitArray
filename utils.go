package sieve

import (
	"math"
)

// PrimeSieveUpto creates a sieve of prime numbers for *odd* numbers only.
func PrimeSieveUpto(upto int) Sieve {

	// design sieve to store only odd numbers
	sieve := NewSieve(upto/2, upto/2)
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

// PrimesUpto returns an array of prime number upto @upto
func PrimesUpto(upto int) []int {

	sieve := PrimeSieveUpto(upto)
	// note: indexes returned are for odd numbers only
	indexes := sieve.ToNums()
	// make space to add 2 at the front
	indexes = append(indexes, 0)
	for i := len(indexes) - 2; i >= 0; i-- {
		indexes[i+1] = 2*indexes[i] + 1
	}
	indexes[0] = 2
	return indexes

}

// FactorsUpto returns a function which returns the factors of the
// number supplied as an array. Returned array format
// [ prime1, prime1-repeats, prime2, prime2-repeats, ...]
func FactorsUpto(upto int) func(int) []int {

	// pre compute primes
	primes := PrimesUpto(upto)
	// init array that will contain the *first* prime divisor
	firsts := make([]int, upto)
	firsts[1] = 1
	// punch through even numbers first
	for i := 2; i < upto; i += 2 {
		firsts[i] = 2
	}
	// step thru other primes skipping even numbers
	for _, prime := range primes[1:] {
		skipeven := 2 * prime
		for pi := prime; pi < upto; pi += skipeven {
			if firsts[pi] < 2 {
				firsts[pi] = prime
			}
		}
	}

	return func(num int) []int {

		factors := make([]int, 0, 20)

		lastprime, lastcount := firsts[num], 0
		for first := firsts[num]; num > 1; first = firsts[num] {
			if lastprime != first {
				factors = append(factors, lastprime, lastcount)
				lastprime, lastcount = first, 1
			} else {
				lastcount++
			}
			num /= first
		}
		factors = append(factors, lastprime, lastcount)

		return factors

	}

}
