package main

func driver(input []int) bool {
	return containsDuplicate(input)
}
func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, i := range nums {
		_, exists := seen[i]
		if exists {
			return true
		}
		seen[i] = true
	}
	return false
}
