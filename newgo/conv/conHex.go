package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// ProcessHex checks if the current word contains "(hex)".
// If it does, it tries to convert the previous word into a hexadecimal string.
func ProcessHex(prevWord string, currentWord string) string {
	// Check if "(hex)" exists anywhere in the current word
	if strings.Contains(currentWord, "(hex)") {
		// Convert the previous word from Hex (Base 16) to Decimal (Base 10)
		// 0 means it will auto-detect bit size (usually 64)
		val, err := strconv.ParseInt(prevWord, 16, 64)
		if err != nil {
			// If it's not a valid hex number, return original word
			return prevWord
		}
		// Return the decimal as a string (e.g., "1E" becomes "30")
		return fmt.Sprintf("%d", val)
	}
	return prevWord
}