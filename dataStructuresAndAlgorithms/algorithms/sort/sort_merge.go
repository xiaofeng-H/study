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
 * 算法思想：将一个前半部分有序和后半部分有序的序列归并为一个完全有序的序列（不开辟新空间）（哈哈哈哈，还是挺绕的哈，晓峰厉害，yyds）
 * 稳定性：稳定排序
 * 时间复杂度：O(n*lgn)
 * 空间复杂度：O(n)
 * 注意：该算法在归并两个有序序列未开辟新的存储空间，是在原有的数组中做了归并（有点难度的），因此时间复杂度比n*lgn要大，
 * 		但空间复杂度为O(1)。书中归并两个有序序列采用了新开辟一个与整个序列同样大小的内存空间来作辅助序列，算法易于实现，
 * 		不会考虑数组插入元素需要后移其他元素的问题，总之，该归并是以时间换空间（其实大可不必---纯属炫技）。
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
归并排序，典型的分治算法；分治，典型的递归结构。
分治算法可以分三步⾛：分解 -> 解决 -> 合并
	1. 分解原问题为结构相同的⼦问题。
	2. 分解到某个容易求解的边界之后，进⾏递归求解。
	3. 将⼦问题的解合并成原问题的解。
归并排序，我们就叫这个函数 merge_sort 吧，按照我们上⾯说的，要明确该函数的职责，即对传⼊的⼀个数组排序。OK，那么这个问题能不能分解
呢？当然可以！给⼀个数组排序，不就等于给该数组的两半分别排序，然后合并就完事了。
void merge_sort(⼀个数组) {
	if (可以很容易处理) return;
	merge_sort(左半个数组);
	merge_sort(右半个数组);
	merge(左半个数组, 右半个数组);
}
好了，这个算法也就这样了，完全没有任何难度。记住之前说的，相信函数的能⼒，传给他半个数组，那么这半个数组就已经被排好了。⽽且你会发现
这不就是个⼆叉树遍历模板吗？为什么是后序遍历？因为我们分治算法的套路是：分解 -> 解决（触底） -> 合并（回溯） 啊，先左右分解，再处理合
并，回溯就是在退栈，就相当于后序遍历了。⾄于 merge 函数，参考两个有序链表的合并，简直⼀模⼀样，下⾯直接贴代码吧。
下⾯参考《算法4》的 Java 代码，很漂亮。由此可⻅，不仅算法思想很重要，编码技巧也是挺重要的吧！多思考，多模仿（建议背下来）。
*/
// 不要在 merge 函数里构造新数组了， 因为 merge 函数会被多次调用，影响性能，直接一次性构造一个足够大的数组，简洁，高效
// tmp 为辅助数组（减小算法的空间复杂度）
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
	// 初始化辅助数组，用于记录 nums 中前后两个已经各自有序的序列
	// 因为在归并排序中，原数组（nums）会被修改，故而需要开辟新空间记录原始数据
	for i := lo; i <= hi; i++ {
		tmp[i] = nums[i]
	}

	// 开始归并
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		// 卧槽，这个边界值处理有点牛逼呀---2023/7/27 17:12
		if i > mid {
			// 此时前半部分有序的数组已经归并完毕，仅剩后半部分尚未归并，直接复制即可
			nums[k] = tmp[j]
			j++
		} else if j > hi {
			// 此时后半部分有序的数组已经归并完毕，仅剩前半部分尚未归并，直接复制即可
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			// 归并：取较小者进入有序序列
			nums[k] = tmp[i]
			i++
		} else {
			// 归并：取较小者进入有序序列
			nums[k] = tmp[j]
			j++
		}
	}

	fmt.Println("二路归并排序后的结果为：")
	algorithms.PrintArray(nums)
}
