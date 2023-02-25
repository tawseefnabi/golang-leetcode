package zigzagconv

import "strings"

func conv(s string, numRows int) string {
	if len(s) < numRows || numRows < 2 {
		return s
	}
	rows := make(map[int][]string)
	i := 0
	for i < len(s) {
		// Trace vertical
		for row := 0; row < numRows && i < len(s); row++ {
			rows[row] = append(rows[row], string(s[i]))
			i++
		}
		// Trace diagonal
		for row := numRows - 2; row > 0 && i < len(s); row-- {
			rows[row] = append(rows[row], string(s[i]))
			i++
		}
	}
	var b strings.Builder
	for row := 0; row < numRows; row++ {
		b.WriteString(strings.Join(rows[row], ""))
	}
	return b.String()
}
