/*
Given a string s, return the longest
palindromic

substring

	in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/
package paindromicstring

func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		odd := expandAroundCenter(s, i, i)
		even := expandAroundCenter(s, i, i+1)
		if len(odd) > len(result) {
			result = odd
		}
		if len(even) > len(result) {
			result = even
		}
	}
	return result
}
func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
