package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isNumberPrime(num int) bool {
	i := 2
	for i < num {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true

}

func isNumberTwoSidedPrime(num int) bool {
	const base = 10
	res := num
	isLeftPrime := false
	isRightPrime := false
	for res > 0 {
		isRightPrime = isNumberPrime(res)
		if !isRightPrime {
			break
		}
		res /= base
	}
	power := 0
	res = num
	for res > 0 {
		power++
		res /= base
	}
	power--
	res = num
	for res > 0 {
		isLeftPrime = isNumberPrime(res)
		if !isLeftPrime {
			break
		}
		divisor := (int)(math.Pow(float64(base), float64(power)))
		res = res % ((res / divisor) * divisor)
		power--

	}

	return isLeftPrime && isRightPrime
}

func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strNum := vars["num"]
	num, err := strconv.ParseInt(strNum, 10, 16)
	if err != nil {
		fmt.Fprintln(w, "invalid integer")
	} else {
		fmt.Fprintln(w, isNumberPrime(int(num)))
	}

}

func TwoSidedPrimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strNum := vars["num"]
	num, err := strconv.ParseInt(strNum, 10, 32)
	if err != nil {
		fmt.Fprintln(w, "invalid integer")
	} else {
		fmt.Fprintln(w, isNumberTwoSidedPrime(int(num)))
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/isNumberPrime/{num}", PrimeHandler)
	router.HandleFunc("/isNumberTwoSidedPrime/{num}", TwoSidedPrimeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
