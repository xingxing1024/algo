package lc_hot100

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func found(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func foundR(s string) int {
	leftP := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			leftP++
		}
		if s[i] == ']' {
			leftP--
			if leftP == 0 {
				return i
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func parse(s string) string {
	if s == "" {
		return ""
	}

	if s[0] >= '0' && s[0] <= '9' {
		// 解析k[encoding_string结构]xxx
		leftIdx := found(s, '[')
		rightIdx := foundR(s)

		num, _ := strconv.Atoi(s[0:leftIdx])
		repeatedString := s[leftIdx+1 : rightIdx]
		prefixString := strings.Repeat(parse(repeatedString), num)
		return prefixString + parse(s[min(rightIdx+1, len(s)):])
	} else {
		// 解析纯字符串
		curIdx := 0
		for curIdx < len(s) && s[curIdx] >= 'a' && s[curIdx] <= 'z' {
			curIdx++
		}
		return s[0:curIdx] + parse(s[curIdx:])
	}
}

func decodeString(s string) string {
	return parse(s)
}

func TestA(t *testing.T) {
	result := decodeString("3[a]2[bc]")
	fmt.Println(result)
}
