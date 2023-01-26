/*
Reverse Integer
https://leetcode.com/problems/reverse-integer/
Given a 32-bit signed integer, reverse digits of an integer.
*/
package reverseinteger

// time complexity: O(log 10 x ) = O(log x)
// space complexity: O(1)
func Reverse(n int) int {
	if n == 0 {
		return n
	}
	isPositive := true
	if n < 0 {
		isPositive = false
		n *= -1
	}
	res := 0
	for n > 0 {
		res = res*10 + n%10
		n /= 10
	}
	if !isPositive {
		res *= -1
	}
	// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
	if res < -1<<31 || res > (1<<31)-1 {
		return 0
	}
	return res
}
