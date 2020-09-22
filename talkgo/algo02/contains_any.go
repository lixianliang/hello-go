package algo02

import (
	"fmt"
)

/*
方法1: 使用hash将s已经存在的字符使用hashmap存储
主串s的长度记为M，子串的长度记为N
时间复杂度：O(N)
空间复杂度：O(M)
*/
func ContainsAny(s, chars string) bool {
	if len(s) == 0 || len(chars) == 0 {
		return false
	}

	chMap := make(map[rune]bool, len(s))
	for _, c := range s {
		chMap[c] = true
	}

	for _, ch := range chars {
		if v, ok := chMap[ch]; !ok {
			fmt.Println(v)
			return false
		}
	}
	return true
}
