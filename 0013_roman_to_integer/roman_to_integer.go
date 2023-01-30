package romantointeger

// Process the string from left to right up to last character.
// Use accumulator variable for summing total.
// Compare the integer value of the current character to the value of the character at the next index.
// If character at next index has larger value, subtract the current character's value from the total.
// Else, add the current character's value to the total.
// Add the value of the last character to the total.
func RomanToInteger(s string) int {
	if len(s) == 0 {
		return 0
	}
	//  Dictionary with values for Roman numerals
	digits := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := digits[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		curr := digits[s[i]]
		pre := digits[s[i-1]]
		if curr > pre {
			sum -= pre
		} else {
			sum += pre
		}
	}
	return sum
}
