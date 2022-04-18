package base_to_base

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	res := ""
	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%X%s", rem, res)
		dec = dec / base
		//Or
		// charSets := "0123456789ABCDEF"
		// rem := dec % base
		// res = string(charSets[rem]) + res
		// dec = dec / base
	}
	return res
}
