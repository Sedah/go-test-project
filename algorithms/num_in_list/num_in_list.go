package main

func NumInList(list []int, num int) bool {
	for _, numbers := range list {
		if numbers == num {
			return true
		}
	}

	return false
}
