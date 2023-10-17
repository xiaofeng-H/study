package sort

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
)

/**
 * @description：交换类排序算法的实现
 * @author：晓峰
 * @date：2020/9/17/11:43
 */

/**
 * 冒泡排序（起泡排序）
 * 算法思想：最简单的排序算法，学不会立即推：放弃编程
 * 稳定性：稳定排序
 * 时间复杂度：O(n^2)---最坏情况：n(n-1)/2;最好情况：O(n)
 * 空间复杂度：O(1)
 * 适用场景：
 */
func BubbleSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	// 起泡算法结束的条件：在一趟排序过程中没有发生关键字的交换
	var flag bool // 标记位：该躺排序中是否发生了关键字交换
	for i := 0; i < len(a); i++ {
		flag = false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	fmt.Println("冒泡排序后的结果为：")
	algorithms.PrintArray(a)
}

// 冒泡排序进阶：双向冒泡排序
func BubbleSortDoubleDirection(arr []int) {
	var left int = 0
	var right int = len(arr) - 1
	var flag bool = true

	for flag {
		flag = false
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				algorithms.Swap(arr, i, i+1)
				flag = true
			}
		}
		right--

		for j := right; j > left; j-- {
			if arr[j] < arr[j-1] {
				algorithms.Swap(arr, j, j-1)
				flag = true
			}
		}
		left++
	}
}

/**
 * 快速排序
 * 算法思想：分治
 * 稳定性：不稳定排序
 * 时间复杂度：O(n*lgn)---最坏情况：O(n^2);最好情况：O(n*lgn)
 * 空间复杂度：O(lgn)
 * 适用场景：待排序列越接近无序，效率越高。快速排序的排序趟数和初始序列有关。与同级别时间复杂度都为O(n*lgn)的排序算法相比，
 * 			该算法的基本操作的最高次项的系数最小，效率最高，故而称为快速排序。
 */
func QuickSort(a []int, low int, high int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	// 以低位为枢轴进行快速排序（从a[low]到a[high]的关键字进行排序）
	i := low  // 低位
	j := high //高位
	if low < high {
		tmp := a[low]
		// 下面这个循环完成一趟排序，即将数组中小于tmp的关键字放在左边，大于tmp的关键字放在右边
		for i < j {
			// 从右往左扫描，找到一个小于tmp的关键字
			for i < j && tmp <= a[j] {
				j--
			}
			if i < j {
				a[i] = a[j] // 放在tmp左边
				i++         // i右移一位
			}

			// 从左往右扫描，找到一个大于tmp的关键字
			for i < j && tmp > a[i] {
				i++
			}
			if i < j {
				a[j] = a[i] // 放在tmp的右边
				j--         // j左移一位
			}
		}
		// 循环结束，i==j，该位置为当前枢轴的最终位置，将tmp放在该位置
		a[i] = tmp
		// 打印划分结果
		fmt.Println("快速排序划分后为：")
		algorithms.PrintArray(a)
		// 递归地对tmp左边的关键字进行排序
		QuickSort(a, low, i-1)
		// 递归地对tmp右边的关键字进行排序
		QuickSort(a, i+1, high)
	}
}
