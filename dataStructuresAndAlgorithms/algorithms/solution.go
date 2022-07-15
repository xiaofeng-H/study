package algorithms

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口经典案例
	// 空串处理
	if s == "" {
		return 0
	}

	// 初始化辅助变量
	var window map[byte]int = make(map[byte]int) // 记录当前窗口中某个字符的出现次数
	left, right, res := 0, 0, 0

	// 窗口滑动
	for right < len(s) {
		c := s[right]
		// 窗口右移，更新数据
		right++
		window[c]++

		// 如果出现重复字符，窗口左移并更新数据
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		// 更新最长子串长度
		if right-left > res {
			res = right - left
		}
	}

	return res
}
