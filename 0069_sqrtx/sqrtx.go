/*
69. Sqrt(x)
https://leetcode.com/problems/sqrtx/
Implement int sqrt(int x).
Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
*/
package sqrtx

func Sqrt(x int) int {
	if x < 2 {
		return x
	}
	var result, start int
	// start := 0
	end := x
	for start <= end {
		mid := int((start + end) / 2)
		sq := mid * mid
		if sq == x {
			result = mid
			break
		}
		if sq > x {
			end = mid - 1
		} else {
			result = mid
			start = mid + 1
		}
	}
	return result
}
