/*
Palindrome Number
https://leetcode.com/problems/palindrome-number/
Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.
*/

package palindromenumber

func IsPalindrome(num int) bool {
	var (
		y int
		z = num
	)
	for num > 0 {
		y = y*10 + num%10
		num /= 10
	}

	return y == z
}
