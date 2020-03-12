package main

import "math"

//建立一维动态数组 dp：

//边界条件：dp[1] = dp[2] = 1，表示长度为 2 的绳子最大乘积为 1；
//状态转移方程：dp[i] = max(dp[i], max((i - j) * j, j * dp[i - j]))

func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1 // 首项
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k])) // 递推公式
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

//贪心算法
//分解为 1,2,3的乘积 优先级3 > 2 > 1
func cuttingRopeGreedyAlgorithm(n int) int {
	if n <= 3 {
		return n - 1
	}

	a := n / 3
	b := n % 3

	//能被三整除 如3 result = 3 的 a 次方
	if b == 0 {
		return int(math.Pow(float64(3), float64(a)))
	}

	//余1 如4 result = 3的a-1次方 乘 2 乘 2
	if b == 1 {
		return int(math.Pow(float64(3), float64(a-1)) * 4)
	}

	//余2 如5 result = 3的a次方
	return int(math.Pow(float64(3), float64(a)) * 2)
}

func main() {
	println(cuttingRope(4))
}
