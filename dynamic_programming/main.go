package main

import (
	"fmt"
	"slices"
)

// https://blog.csdn.net/XGLLHZ/article/details/134533484
// 01背包问题
// dp[i][j] = max(dp[i-1, j], dp[i-1, j-v[i]] + w[i])
func pack01() {
	v := [5]int{-1, 1, 2, 3, 4}
	w := [5]int{-1, 4, 5, 1, 2}

	var dp [100][100]int

	for i := 1; i <= 4; i++ {
		for j := 0; j <= 10; j++ {
			if j-v[i] >= 0 {
				dp[i][j] = slices.Max([]int{dp[i-1][j], dp[i-1][j-v[i]] + w[i]})
			}
		}
	}
	fmt.Println(dp[4][9])
}

// 多重背包问题
// dp[i][j] = max(dp[i-1, j], dp[i-1, j-k*v[i]] + w[i])
func pack02() {
	v := [5]int{-1, 1, 2, 3, 4}
	w := [5]int{-1, 4, 5, 1, 2}
	num := [5]int{-1, 2, 2, 2, 2}

	var dp [100][100]int
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 10; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= num[i]; k++ {
				if j-k*v[i] >= 0 {
					dp[i][j] = slices.Max([]int{dp[i][j], dp[i-1][j-k*v[i]] + w[i]})
				}
			}
		}
	}

	fmt.Println(dp[4][10])
}

// 完全背包问题
// dp[i][j] = max(dp[i-1, j], dp[i-1, j-k*v[i]] + w[i])
func pack03() {
	v := [5]int{-1, 1, 2, 3, 4}
	w := [5]int{-1, 4, 5, 1, 2}

	var dp [100][100]int
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 10; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; j-k*v[i] >= 0; k++ {
				dp[i][j] = slices.Max([]int{dp[i][j], dp[i-1][j-k*v[i]] + w[i]})
			}
		}
	}

	fmt.Println(dp[4][10])
}

// 最长公共子序列
// dp[i][j] = dp[i-1][j-1] + 1 if s1[i-1] == s2[j-1]
// dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
func lcs() {
	s1 := "abdc"
	s2 := "axbdc"

	s1Len := len(s1)
	s2Len := len(s2)
	var dp [100][100]int
	for i := 1; i <= s1Len; i++ {
		for j := 1; j <= s2Len; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = slices.Max([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]})
			}
		}
	}

	fmt.Println(dp[s1Len][s2Len])
}

// 最长公共子串
// dp[i][j]：以i,j结尾的最长公共子串
// dp[i][j] = dp[i-1][j-1] + 1 if s1[i-1] == s2[j-1]
// dp[i][j] = 0
func lsc() {
	s1 := "abc"
	s2 := "abdc"

	s1Len := len(s1)
	s2Len := len(s2)
	var dp [100][100]int
	result := 0

	for i := 1; i <= s1Len; i++ {
		for j := 1; j <= s2Len; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			result = slices.Max([]int{result, dp[i][j]})
		}
	}

	fmt.Println(result)
}

// 最长回文串 dp[i][j]：是否为回文串
// dp[i][j] = dp[i+1][j-1] if s[i-1] == s[j-1]
// dp[i][j] = false if if s[i-1] != s[j-1]
func longReverseString() {
	s1 := "ababbb"

	var dp [100][100]bool
	stringLength := len(s1)

	// 初始化空串为回文串 i <= j
	for i := 0; i < stringLength+10; i++ {
		for j := 0; j < stringLength+10; j++ {
			if !(i <= j) {
				dp[i][j] = true // 初始化空串为回文串
			}
		}
	}

	result := 0
	for sLen := 1; sLen <= stringLength; sLen++ {
		for i := 1; i <= stringLength; i++ {
			j := i + sLen - 1
			if j > stringLength {
				continue
			}
			if s1[i-1] == s1[j-1] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}
			if dp[i][j] == true {
				result = slices.Max([]int{result, j - i + 1})
			}
		}
	}

	fmt.Println(result)
}

func longestSeq() {
	numList := []int{1, 2, 10, 4, 5}
	var dp [100]int

	result := 0
	for i := 0; i < len(numList); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if numList[j] < numList[i] {
				dp[i] = slices.Max([]int{dp[i], dp[j] + 1})
			}
		}
		result = slices.Max([]int{result, dp[i]})
	}
	fmt.Println(result)
}

// 最大连续和
// dp[i] = max(a[i], dp[i-1] + a[i])
func maxSum() {
	numList := []int{1, 2, -5, 4, 5}
	var dp [100]int

	result := 0
	for i := 1; i <= len(numList); i++ {
		dp[i] = slices.Max([]int{numList[i-1], dp[i-1] + numList[i-1]})
		result = slices.Max([]int{result, dp[i]})
	}
	fmt.Println(result)
}

func main() {
	//pack01()
	//pack02()
	//pack03()
	//lcs()
	//lsc()
	//longReverseString()
	//longestSeq()
	//maxSum()
}
