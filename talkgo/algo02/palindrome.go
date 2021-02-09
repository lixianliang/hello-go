package algo02

import (
	"fmt"
)

/*
	方法：动态规划
	字符串的长度为N
	时间复杂度：O(N*N)
	空间复杂度：O(N*N)
*/
func Palindrom(s string) (int, string) {
	if len(s) < 2 {
		return len(s), s
	}

	var maxN int
	var maxStr string
	// dp为二维数组，第一纬存储子字符串起始位置，第二纬存储终点 dp[i][j]存储s[i,j]回文字符串最大长度
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	// 构造最大回文字符串map，key为"i_j"，value为[i,j]范围内最大的回文字符串
	maxStrMap := make(map[string]string, len(s))
	for i, _ := range dp {
		dp[i][i] = 1
		key := fmt.Sprintf("%d_%d", i, i)
		maxStrMap[key] = string(s[i])
	}

	for i := 1; i < len(s); i++ { // 回文动态规划范围为[2, N], i+1表示当前循环的字符串长度
		for j := 0; i+j < len(s); j++ { // 计算范围[j，i+j]的回文长度和回文字符串,j表示起始字符串
			var oldkey string
			key := fmt.Sprintf("%d_%d", j, i+j)
			if s[j] == s[i+j] { // 字符串s[j]和s[i+j]
				dp[j][i+j] = dp[j+1][i+j-1] + 2 //
				oldkey = fmt.Sprintf("%d_%d", j+1, i+j-1)
				maxStrMap[key] = string(s[j]) + maxStrMap[oldkey] + string(s[i+j])
			} else {
				if dp[j][i+j-1] > dp[j+1][i+j] {
					dp[j][i+j] = dp[j][i+j-1]
					oldkey = fmt.Sprintf("%d_%d", j, i+j-1)
					maxStrMap[key] = maxStrMap[oldkey]
				} else {
					oldkey = fmt.Sprintf("%d_%d", j+1, i+j)
					dp[j][i+j] = dp[j+1][i+j]
				}
				maxStrMap[key] = maxStrMap[oldkey]
			}
		}
	}

	maxN = dp[0][len(s)-1] // 回文最大长度
	key := fmt.Sprintf("%d_%d", 0, len(s)-1)
	maxStr, _ = maxStrMap[key] // 回文最长字符串
	return maxN, maxStr
}
