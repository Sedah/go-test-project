package base_to_base

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	charSets := "0123456789ABCDEF"
	multiplier := 1
	res := 0

	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charSets {
			if char == rune(value[i]) {
				index = j
				break
			}

		}
		res = res + index*multiplier
		multiplier = multiplier * base
	}

	return res
}
