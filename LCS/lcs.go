package main

import "fmt"

func lcs(x, y string) ([][]int, [][]int) {
	x = " " + x
	y = " " + y

	m, n := len(x), len(y)
	dp := make([][]int, m)
	sp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range sp {
		sp[i] = make([]int, n)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if x[i] == y[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				sp[i][j] = 1
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					sp[i][j] = 2
					dp[i][j] = dp[i][j-1]
				} else {
					sp[i][j] = 3
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return dp, sp
}

func getLCS(i, j int, sp [][]int, x string) string {
	if (i == 0) || (j == 0) {
		return ""
	}

	if sp[i][j] == 1 {
		return getLCS(i-1, j-1, sp, x) + string(x[i])
	} else if sp[i][j] == 2 {
		return getLCS(i, j-1, sp, x)
	} else if sp[i][j] == 3 {
		return getLCS(i-1, j, sp, x)
	}

	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := "ABCDE"
	b := "BDEF"
	res, res2 := lcs(a, b)

	for i := range res {
		fmt.Println(res[i])
	}

	fmt.Println()

	for i := range res2 {
		fmt.Println(res2[i])
	}

	fmt.Println()

	res3 := getLCS(len(a), len(b), res2, " "+a)
	fmt.Println(res3)
}
