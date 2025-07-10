package double_idx

import (
	"fmt"
	"testing"
)

func cover(curCount map[byte]int, tCount map[byte]int) bool {
	for k, _ := range tCount {
		if curCount[k] < tCount[k] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	// 统计t字符串
	tMap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	result := ""
	sIdx, eIdx := 0, 0
	curMap := make(map[byte]int, 0)
	for {
		for eIdx < len(s) && !cover(curMap, tMap) {
			curMap[s[eIdx]]++
			eIdx++
		}

		if cover(curMap, tMap) {
			fmt.Println(s[sIdx:eIdx])
			if result == "" {
				result = s[sIdx:eIdx]
			} else if eIdx-sIdx < len(result) {
				result = s[sIdx:eIdx]
			}
		}

		if sIdx == len(s) {
			break
		}

		curMap[s[sIdx]]--
		sIdx++
	}

	return result
}

func TestA(t *testing.T) {
	result := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println(result)
}
