package piscine

func Rot14(str string) string {
	var result string
	for _, char := range str {
		// Check if the character is an uppercase or lowercase letter
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			var base byte
			if char >= 'a' && char <= 'z' {
				base = 'a'
			} else {
				base = 'A'
			}
			// Calculate the new character by adding 14 to its ASCII value
			newChar := ((byte(char)-base)+14)%26 + base
			result += string(newChar)
		} else {
			// For non-alphabetic characters, keep them unchanged
			result += string(char)
		}
	}
	return result
}
