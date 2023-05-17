package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 求解二叉树最大深度
// 二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用递归求解
func maxDepth(root *TreeNode) int {
	// 如果是叶子结点，直接返回1（递归终止条件）
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 最大深度为左右子树最大深度+1
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 求两个数（int）的较大者
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 求解日志统计
func solve() []string {
	// 记录结果 key=url value=url出错的次数
	var res map[string]int = make(map[string]int)
	// 带缓存方式的读取文件
	// 假设文件日志路径
	filePath := "/test/file/read/url.log"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("read file err", err)
	}
	// 关闭文件句柄
	defer file.Close()
	// 带缓存方式读取文件
	reader := bufio.NewReader(file)
	// 循环的读取文件内容
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 读取到文件的末尾
		}
		if analyze(content) {
			// 该url在七天之内出错过
			array := strings.Split(" ", content)
			// 该url出现的次数加一
			res[array[len(array)-1]]++
		}
	}

	// 使用堆排序来统计出错率最高的三位
	maxThird := make([]string, 3)
	arr := make([]int, 0)
	for _, v := range res {
		arr = append(arr, v)
	}
	heapSort(arr)
	for k, v := range res {
		if v == arr[0] || v == arr[1] || v == arr[2] {
			maxThird = append(maxThird, k)
		}
	}
	return maxThird
}

func heapSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	// 建立初始堆
	for i := len(a) / 2; i >= 0; i-- {
		shift(a, i, len(a)-1)
	}
	// 进行len(a）-1次循环，完成堆排序
	var temp int // 临时变量
	for i := len(a) - 1; i >= 1; i-- {
		// 换出根结点中的关键字，将其放入最终位置
		temp = a[0]
		a[0] = a[i]
		a[i] = temp
		// 在减少了一个关键字的无序序列中进行调整
		shift(a, 0, i-1)
	}
}

// 在数组a[low]到a[high]的范围内对在位置low上的结点进行调整
func shift(a []int, low int, high int) {
	i := low     // 父结点下标
	j := 2*i + 1 // a[j]是a[i]的左孩子结点（注意：若数组从下标0开始储存数据，则i结点对应的左孩子的下标为2*i+1；若数组从下标1开始储存数据，则为2*i）
	temp := a[i] // temp指向父结点
	for j <= high {
		// 若右孩子较大，则把j指向右孩子
		if j < high && a[j] < a[j+1] {
			j++ // j变为2*i+2
		}
		// 若父结点小于孩子结点的值，说明当前堆不满足大顶堆定义，进行调整
		if temp < a[j] {
			// 将a[j]调整到双亲结点的位置上，同时修改i和j的值，以便继续向下调整
			a[i] = a[j]
			i = j
			j = 2*i + 1
		} else {
			break // 调整结束
		}
		// 被调整结点的值放入最终位置
		a[i] = temp
	}
}

func analyze(str string) bool {
	timeStr := str[:19]
	bytes := []byte(timeStr)
	// 将时间格式化标准
	for k := range bytes {
		if bytes[k] == '/' {
			bytes[k] = '-'
		}
	}
	timeStr = string(timeStr)

	// 转化为时间戳，并判读是否在七天以内
	timeLayout := "2006-01-02 15:04:05"                          //转化所需模板
	loc, _ := time.LoadLocation("Local")                         //获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //使用模板在对应时区转化为time.time类型
	timestamp := theTime.Unix()
	// 前7天
	timeZone := time.FixedZone("CST", 8*3600) // 东八区
	nowTime := time.Now().In(timeZone)
	beforeTime := nowTime.AddDate(0, 0, -7)
	// 时间转换格式
	beforeTimeS := beforeTime.Unix() // 秒时间戳

	// 如果是七天之前，则排除
	if timestamp < beforeTimeS {
		return false
	} else {
		// 判断是否出错
		tmp := str[19:]
		// 将字符串以空格分割
		array := strings.Split(" ", tmp)
		// 判断是否出错
		if array[0] == "error" {
			return true
		} else {
			return false
		}
	}
}
