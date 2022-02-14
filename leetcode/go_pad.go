package main

type LCInput [][]int

func driver(input LCInput) int {
	return minimumTotal(input)
}

const MAXINT = 100000 // because LC doesn't have the math package. else I'd use math.MaxInt

func minimumTotal(triangle [][]int) int {
	dp := newDP(triangle)
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := range triangle[i] {
			dp[i][j] = triangle[i][j] + minInt(getPrevious(dp, i, j))
		}
	}

	return minInt(dp[len(triangle)-1]...)
}

func minInt(nums ...int) int {
	min := MAXINT
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func newDP(input [][]int) [][]int {
	dp := make([][]int, len(input))
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	return dp
}

func getPrevious(input [][]int, i, j int) (int, int) {
	switch {
	case j == 0:
		return input[i-1][j], MAXINT
	case j == i:
		return input[i-1][j-1], MAXINT
	default:
		return input[i-1][j], input[i-1][j-1]
	}
}
