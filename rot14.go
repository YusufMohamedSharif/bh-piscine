package piscine

func rot14(str string) string {
	var result string
	for _, char := range str {
		// Check if the character is an uppercase or lowercase letter
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Calculate the new character by adding 14 to its ASCII value
			newChar := char + 14
			// Handle wrap-around for lowercase letters
			if char >= 'a' && newChar > 'z' {
				newChar -= 26
			}
			// Handle wrap-around for uppercase letters
			if char >= 'A' && newChar > 'Z' {
				newChar -= 26
			}
			result += string(newChar)
		} else {
			// For non-alphabetic characters, keep them unchanged
			result += string(char)
		}
	}
	return result
}
