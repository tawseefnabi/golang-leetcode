package stringtointeger

import (
	"math"
	"strings"
)

func Atoi(s string) int {
	// Remove any additional spaces before and after the given string
	s = strings.Trim(s, " ")
	n := len(s)
	// If string is empty return 0
	if n == 0 {
		return 0
	}

	// String index from where the processing should start
	var start int
	signMultipler := 1
	if s[0] == '-' {
		signMultipler = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	// ASCII of '0' = 48
	// s[i] - '0' gives the actual number as an integer
	var res int
	for i := start; i < len(s); i++ {
		// Handling for non numeric values
		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMultipler
		}
		// Calculate the result for current iteration
		res = res*10 + int(s[i]-'0')
		// Check if result exceeds MinInt32 or MaxInt32
		if res*signMultipler <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultipler >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * signMultipler
}
