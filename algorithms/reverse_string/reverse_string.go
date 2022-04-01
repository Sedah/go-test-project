package main

// func Reverse(word string) string {
// 	runes := []rune(word)
// 	//make i and j, i=0 j= len(runes)-1
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		//reverse it
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)

// }

func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res

}
