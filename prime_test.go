package main

import (
	"testing"
)

func TestPrime(t *testing.T) {
	twoSidedPrimes := []int{2, 3, 5, 7, 23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397}
	for i := 0; i < 15; i++ {
		if !isNumberTwoSidedPrime(twoSidedPrimes[i]) {
			t.Error("FAIL!!! Expected", twoSidedPrimes[i], "to be two sided prime, found otherwise")
		}
	}
	nonPrimes := []int{4, 6, 9, 25}
	for i := 0; i < 4; i++ {
		if isNumberTwoSidedPrime(nonPrimes[i]) {
			t.Error("FAIL!!! Expected", nonPrimes[i], "to be non prime, found otherwise")
		}
	}

	primeButNotTwoSided := []int{79, 97}
	for i := 0; i < 2; i++ {
		if !(isNumberPrime(primeButNotTwoSided[i]) && !isNumberTwoSidedPrime(primeButNotTwoSided[i])) {
			t.Error("FAIL!!! Expected", primeButNotTwoSided[i], "to be prime but not two sided, found otherwise")
		}
	}

}
