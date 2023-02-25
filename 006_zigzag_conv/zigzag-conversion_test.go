package zigzagconv

// Testing in go
import (
	"fmt"
	"testing"
)

func Test_conv(t *testing.T) {
	fmt.Println("===", conv("PAYPALISHIRING", 2))
	fmt.Println("===", conv("PAYPALISHIRING", 3))
	fmt.Println("===", conv("PAYPALISHIRING", 4))
	fmt.Println("===", conv("A", 1))
}
