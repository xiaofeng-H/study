package sort

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
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
 *		新开辟一个与整个序列同样大小的内存空间来作辅助序列，算法易于实现，不会考虑数组插入元素需要后移其他元素的问题，总之，该种归并是以时间缓空间。
 */
func MergeSort1(a []int, low int, high int) {
	// 递归结束条件
	if low >= high {
		return
	}

	// 防止溢出
	mid := low + (high-low)/2
	// 归并排序前半段
	MergeSort1(a, low, mid)
	// 归并排序后半段
	MergeSort1(a, mid+1, high)
	// 将前半段有序序列和后半段有序序列归并成整体有序序列
	merge1(a, low, mid, high)
	fmt.Println("二路归并排序后的结果为：")
	algorithms.PrintArray(a)

}

// 将两端已经有序的序列合成一条整体有序的序列（以时间换空间）
func merge1(a []int, low int, mid int, high int) {
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

/*
分治算法
归并排序，典型的分治算法；分治，典型的递归结构。
分治算法可以分三步⾛：分解 -> 解决 -> 合并
1. 分解原问题为结构相同的⼦问题。
2. 分解到某个容易求解的边界之后，进⾏第归求解。
3. 将⼦问题的解合并成原问题的解。
归并排序，我们就叫这个函数 merge_sort  吧，按照我们上⾯说的，要明确
该函数的职责，即对传⼊的⼀个数组排序。OK，那么这个问题能不能分解
呢？当然可以！给⼀个数组排序，不就等于给该数组的两半分别排序，然后
合并就完事了。
void merge_sort(⼀个数组) {
	if (可以很容易处理) return;
	merge_sort(左半个数组);
	merge_sort(右半个数组);
	merge(左半个数组, 右半个数组);
}
好了，这个算法也就这样了，完全没有任何难度。记住之前说的，相信函数
的能⼒，传给他半个数组，那么这半个数组就已经被排好了。⽽且你会发现
这不就是个⼆叉树遍历模板吗？为什么是后序遍历？因为我们分治算法的套
路是 分解 -> 解决（触底） -> 合并（回溯） 啊，先左右分解，再处理合
并，回溯就是在退栈，就相当于后序遍历了。⾄于 merge  函数，参考两个
有序链表的合并，简直⼀模⼀样，下⾯直接贴代码吧。
下⾯参考《算法4》的 Java 代码，很漂亮。由此可⻅，不仅算法思想思想重
要，编码技巧也是挺重要的吧！多思考，多模仿。
*/
// 不要在merge函数里构造新数组了， 因为merge函数会被多次调用，影响性能，直接一次性构造一个足够大的数组，简洁，高效
// tmp为辅助数组（减小算法的空间复杂度）

func MergeSort2(nums, tmp []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	MergeSort2(nums, tmp, lo, mid)
	MergeSort2(nums, tmp, mid+1, hi)
	merge2(nums, tmp, lo, mid, hi)
}

// 以空间换时间
func merge2(nums, tmp []int, lo, mid, hi int) {
	// 初始化辅助数组
	for i := lo; i <= hi; i++ {
		tmp[i] = nums[i]
	}

	// 开始归并
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = tmp[j]
			j++
		} else if j > hi {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}

	fmt.Println("二路归并排序后的结果为：")
	algorithms.PrintArray(nums)
}
