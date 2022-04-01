package main

// func Sum(numbers []int) int {
// 	result := 0
// 	for _, num := range numbers {
// 		result += num
// 	}
// 	return result

// }

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
