package main

func driver(input LCInput) []int {
	return searchRange(input.arr, input.target)
}

func searchRange(nums []int, target int) []int {
	left, exists := findLeftIfExist(nums, target)

	if !exists {
		return []int{-1, -1}
	}
	right := findRight(nums, target)
	return []int{left, right}
}

func findLeftIfExist(nums []int, target int) (int, bool) {
	i := 0
	j := len(nums) - 1
	var m int

	for i <= j {
		m = (i + j) / 2
		switch {
		case nums[m] == target && isLeft(m, nums):
			return m, true
		case nums[m] > target || (nums[m] == target && nums[m-1] == target):
			j = m - 1
		case nums[m] < target:
			i = m + 1
		}
	}
	return -1, false
}

func findRight(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	var m int

	for i <= j {
		m = (i + j) / 2
		switch {
		case nums[m] == target && isRight(m, nums):
			return m
		case nums[m] > target:
			j = m - 1
		case nums[m] < target || (nums[m] == target && nums[m+1] == target):
			i = m + 1
		}
	}
	return m
}

func isRight(m int, nums []int) bool {
	return m == len(nums)-1 || nums[m+1] != nums[m]
}

func isLeft(m int, nums []int) bool {
	return m == 0 || nums[m-1] != nums[m]
}

type LCInput struct {
	arr    []int
	target int
}
