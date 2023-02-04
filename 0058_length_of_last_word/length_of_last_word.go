/*
58. Length of Last Word
https://leetcode.com/problems/length-of-last-word/
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word in the string.
If the last word does not exist, return 0.
Note: A word is defined as a character sequence consists of non-space characters only.

Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/
package lengthoflastword

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	los := 0
	str := strings.Trim(s, " ")
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			los = 0
		} else {
			los++
		}
	}
	return los
}
