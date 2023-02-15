/*
66. Plus One
https://leetcode.com/problems/plus-one/
Given a non-empty array of digits representing a non-negative integer,
plus one to the integer.
The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contain a single digit.
You may assume the integer does not contain any leading zero, except the number 0 itself.
*/

package plusone

// import "fmt"

func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits

}
