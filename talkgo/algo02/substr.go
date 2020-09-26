package main

import (
	"fmt"
)

func main() {
	plist := []string{"abc", "de", "af", "deabf"}
	//plist := []string{"abc"}
	s := "dabcedeabfe"
	Bf(s, plist)
	Bm(s, plist)
}

func Bf(s string, strs []string) {
	for _, ss := range strs {
		index := BfSubstr(s, ss)
		if index != -1 {
			fmt.Printf("%s->[%d]\n", ss, index)
		}
	}
}

/*
方法1: 使用朴素匹配算法
主串的长度记为n，子串的长度记为m
时间复杂度：O(n*m)
空间复杂度：O(1)
*/
func BfSubstr(s, substr string) int {
	if len(s) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}

	// 主串从起始位置[0 - n-m+1]范围内查找字串是否匹配
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}

	return -1
}

func Bm(s string, strs []string) {
	for _, ss := range strs {
		index := BmSubstr(s, ss)
		if index != -1 {
			fmt.Printf("%s->[%d]\n", ss, index)
		}
	}
}

/*
方法2: 使用BM匹配算法
主串的长度记为N，子串的长度记为M
时间复杂度：O(M)
空间复杂度：O(3*N)	计算比较复杂
*/
func BmSubstr(s, substr string) int {
	if len(s) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}

	// 字符集散列表，用于不匹配字符串的快速查找（坏字符）
	const CHAR_SIZE = 256
	var bc [CHAR_SIZE]int
	for i := 0; i < len(substr); i++ {
		ascii := int(substr[i])
		bc[ascii] = i
	}
	// fmt.Printf("bc: %v\n", bc)

	suffix, prefix := generateGS(substr)
	// fmt.Printf("suffix:%v prefix:%v\n", suffix, prefix)

	n := len(s)
	m := len(substr)
	i := 0
	for i <= n-m {
		j := 0
		for j = m - 1; j >= 0; j-- { // 从模式匹配串后往前匹配
			if s[i+j] != substr[j] { // 找到不匹配的的字符串
				break
			}
		}
		if j < 0 {
			return i
		}

		step1 := j - bc[int(s[i+j])] // 通过坏字符串计算滑动的位数
		step2 := 0
		if j < m-1 {
			step2 = moveByGS(j, m, suffix, prefix)
		}
		if step1 > step2 {
			i = i + step1
		} else {
			i = i + step2
		}
	}

	return -1
}

func generateGS(pattern string) ([]int, []bool) {
	m := len(pattern)
	suffix := make([]int, m)
	prefix := make([]bool, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ { // 循环计算
		j := i
		k := 0 // 公共后缀字串的长度
		for j >= 0 && (pattern[j] == pattern[m-1-k]) {
			// 计算[0-j]的前缀字符串和后缀的公共字串长度
			j--
			k++
			suffix[k] = j + 1 // 后面找到的字符串位置会覆盖之前的，靠后的公共字串用于好后缀移位
		}

		if j == -1 {
			prefix[k] = true // prefix[k]也为前缀子串
		}
	}

	return suffix, prefix
}

func moveByGS(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j // 好后缀长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	// 部分前缀需要移动的位置
	for i := j + 2; i <= m-1; i++ {
		if prefix[m-i] {
			return i
		}
	}
	return m
}
