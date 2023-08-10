package sort

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
)

/**
 * @description：选择类排序算法的实现
 * @author：晓峰
 * @date：2021/10/16/16:38
 */

/**
 * 简单选择排序
 * 算法思想：过于简单，无需赘述
 * 稳定性：不稳定排序
 * 时间复杂度：O(n^2)
 * 空间复杂度：O(1)
 * 适用场景：
 */
func SimpleSelectSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	var index int // 最小值的数组下标
	var tmp int   // 临时变量
	for i := 0; i < len(a); i++ {
		index = i
		// 本算法精髓：从无序序列中挑出一个最小的关键字
		for j := i + 1; j < len(a); j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		tmp = a[i]
		a[i] = a[index]
		a[index] = tmp
	}

	fmt.Println("简单选择排序后的结果为：")
	algorithms.PrintArray(a)
}

/**
 * 堆排序
 * 算法思想：利用堆这种数据结构本身的特点来进行排序操作。
 * 稳定性：不稳定排序
 * 时间复杂度：O(n*log2n)。堆排序在最坏情况下的时间复杂度也是O(n*log2n)，这是它相对于快速排序最大的优点。
 * 空间复杂度：O(1)。堆排序的空间复杂度为O(1)，在所有时间复杂度为O(n*log2n)的排序中是最小的，这也是其一大优点。
 * 适用场景：适用的场景是关键字很多的情况，典型的例子是从10 000个关键字中选出前十个最小的，这种情况用堆排序最好。
 *			如果关键字个数较少，则不提倡适用堆排序。
 */
func HeapSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	// 建立初始堆，从第一个非叶子结点开始调整
	// 注：1.若数组下标从1开始，则第一个非叶结点为 n/2 取下整；
	//     2.若数组下标从0开始，则第一个非叶结点为 (n-2)/2 取下整。
	for i := len(a)/2 - 1; i >= 0; i-- {
		Shift(a, i, len(a)-1)
	}
	// 进行len(a）-1次循环，完成堆排序
	var temp int // 临时变量
	for i := len(a) - 1; i >= 1; i-- {
		// 换出根结点中的关键字，将其放入最终位置
		temp = a[0]
		a[0] = a[i]
		a[i] = temp
		// 在减少了一个关键字的无序序列中进行调整
		Shift(a, 0, i-1)
	}

	fmt.Println("堆排序后的结果为：")
	algorithms.PrintArray(a)
}

// 在数组a[low]到a[high]的范围内对在位置low上的结点进行调整
func Shift(a []int, low int, high int) {
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
