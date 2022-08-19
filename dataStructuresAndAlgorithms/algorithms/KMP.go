package algorithms

// 初始化next数组
func GetNext(pat string) []int {
	var next []int = make([]int, len(pat))
	next[0] = 0

	i := 0
	j := 0
	for i < len(pat)-1 {
		// j == 0 -- 不存在相等前后缀
		if j == 0 {
			i++
			next[i] = 0
		}
		if pat[i] == pat[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}
func GetNextVal(pat string) []int{
	var nextVal []int = make([]int, len(pat))
	nextVal[0] = 0

	i := 0
	j := 0
	for i < len(pat)-1 {
		// j == 0 -- 不存在相等前后缀
		if j == 0 {
			i++
			nextVal[i] = 0
		}
		if pat[i] == pat[j] {
			i++
			j++
			if pat[i] == pat[j] {
				nextVal[i] = nextVal[j]
			} else {
				nextVal[i] = j
			}
		} else {
			j = nextVal[j]
		}
	}

	return nextVal
}
// KMP
func KMP(str, pat string) int {
	if str == "" {
		return -1
	}

	// 获取next数组
	//next := GetNext(pat)
	nextVal := GetNextVal(pat)
	i := 0
	j := 0
	for i < len(str) && j < len(pat) {
		if j == 0 || str[i] == pat[j] {
			i++
			j++
		} else {
			j = nextVal[j]
		}
	}

	if j == len(pat) {
		return i - len(pat)
	} else {
		return -1
	}
}
