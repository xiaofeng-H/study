package algorithms

import (
	"fmt"
)

/**
 * @description：归并类排序算法的实现
 * @author：晓峰
 * @date：2021/10/16/21:45
 */

/**
 * 二路归并排序
 * 算法思想：注意将一个前半部分有序和后半部分有序的数组归并为一个完全有序的序列的归并算法（不开辟新空间）（哈哈哈哈，还是挺绕的哈，晓峰厉害，yyds）
 * 稳定性：稳定排序
 * 时间复杂度：O(n*log2n)
 * 空间复杂度：O(n)
 * 注意：该算法中归并两个有序序列未开辟新的存储空间，是在原有的数组中做了归并（有点难度的），因此时间复杂度比n*log2n要大，但空间复杂度为O(1)。书中归并两个有序序列采用了
 *		新开辟一个与整个序列同样大小的内存空间来存储排好序的新序列，算法易于实现，不会考虑数组插入元素需要后移其他元素的问题。
 */
func MergeSort(a []int, low int, high int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	if low < high {
		mid := (low + high) / 2
		// 归并排序前半段
		MergeSort(a, low, mid)
		// 归并排序后半段
		MergeSort(a, mid+1, high)
		// 将前半段有序序列和后半段有序序列归并成整体有序序列
		Merge(a, low, mid, high)
		fmt.Println("二路归并排序后的结果为：")
		PrintArray(a)
	}
}

// 将两端已经有序的序列合成一条整体有序的序列
func Merge(a []int, low int, mid int, high int) {
	// 该归并算法是将后半部分有序序列归并到前半部分有序序列
	var tmp int // 临时变量
	i := mid + 1
	for i <= high && low <= mid {
		if a[low] < a[i] {
			low++
		} else {
			// 将前半部分有序序列整体后移一位
			tmp = a[i]
			j := i
			for j > low {
				a[j] = a[j-1]
				j--
			}
			// 将后半部分插入到合适的位置，再进行下一个数的判断
			a[j] = tmp
			i++
			// 注意：因为前半部分有序序列插入了一个数，所以前半部分有序序列的将要进行下一次比较的下标和尾下标都要相应的向后移一位
			low++
			mid++
		}
	}
}
