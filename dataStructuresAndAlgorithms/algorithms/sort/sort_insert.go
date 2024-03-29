package sort

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	. "study/dataStructuresAndAlgorithms/dataStructures"
)

/**
 * @description：插入类排序算法的实现
 * @author：晓峰
 * @date：2020/9/15/13:49
 */

/**
 * 直接插入排序
 * 算法思想：每趟将一个待排序的关键字按照其值的大小插入到已经排好的部分有序序列的适当位置上，直到所有待排关键字都被插入到有序序列中为止。
 * 稳定性：稳定排序
 * 时间复杂度：O(n^2)---最坏情况：n(n-1)/2;最好情况：O(n)
 * 空间复杂度：O(1)
 * 适用场景：适用于序列基本有序的情况
 */
func DirectInsertSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	// 1.因为第一个数天然有序，所以从第二位才开始进行插入排序
	for i := 1; i < len(a); i++ {
		// 2.从本算法思想可以看出，直接插入排序的核心思想是找到待排关键字在本趟排序中的恰当位置，所以无可避免的是要对原始序列
		// 元素进行移位操作。为统一化，也为了方便理解，以下代码是数组元素移位操作的基本模板代码
		// 2.1 首先记录待排关键字，因为在移位操作中，该关键字可能会被覆盖
		temp := a[i]
		// 2.2 一步步向前比较，找到合适的位置，直到数组到头
		j := i - 1
		for j >= 0 && a[j] > temp {
			// 如果当前比较位置的关键字比待排关键字大，则后移当前位置的关键字，下标指向当前位置的前一位
			a[j+1] = a[j]
			j--
		}
		// 3.经过上面的移位操作，j指向了待排关键字需要插入位置的前一位，插入即可
		a[j+1] = temp
	}

	// 打印结果
	fmt.Println("直接插入排序后的结果为：")
	algorithms.PrintArray(a)
}

// @desc 直接插入排序（单链表）
// @param *ListNode 单链表头结点的地址
func DirectInsertSortOfLinkList(head *ListNode) {
	var cur, tmp, node, pre *ListNode
	// 1.单链表不为空
	if head.Next != nil {
		// 2.默认第一个结点有序，cur指向单链表第一个结点的后继
		cur = head.Next.Next
		// 3.单链表断链
		head.Next.Next = nil
		// 4.第一个结点有直接后继时，才真正进行直接插入排序
		for cur != nil {
			// 4.1 pre指向头结点
			pre = head
			// 4.2 node指向第一个结点
			node = pre.Next
			for node != nil && node.Val < cur.Val {
				// 4.3 在有序表中找到一个结点p，使其val值刚好大于cur.val
				pre = node
				node = node.Next
			}
			// 4.4 将cur插入到已找到的位置
			tmp = cur.Next
			pre.Next = cur
			cur.Next = node
			cur = tmp
		}
	}
}

/**
 * 折半插入排序
 * 算法思想：与直接插入排序类似，只是在查找插入位置时使用二分查找以提升查找速率（并不影响算法的时间复杂度）
 * 稳定性：稳定排序
 * 时间复杂度：O(n^2)---最坏情况：O(n^2);最好情况：O(n*lgn)
 * 空间复杂度：O(1)
 * 适用场景：适用于关键字树较多的情况
 */
func HalfInsertSort(a []int) {
	// 判空检查
	if len(a) <= 0 {
		fmt.Println("待排序列为空，请重新输入！！！")
	}

	for i := 1; i < len(a); i++ {
		j := -1 // 标记位，记录一个比目标值大且最接近目标值的数的数组下标
		mid := 0
		head := 0
		tail := i - 1

		/* 以下这段代码还是挺难理解的，书上也没有折半查找的代码，现在很好奇这段代码是在哪抄的还是自己想出来的，哈哈哈！！！（2021/10/16 14:54）*/
		/* 这有啥难理解的，不很简单吗？？？但是不得不说这个标记位用得还是很妙的(2021/11/25 22:49 )*/
		// 以下代码用于查找一个最小的且比目标值大的数
		for head <= tail { // 当head==tail时，也要进行判断
			mid = (head + tail) / 2
			// 该循环是找一个比目标值大且最近的数，故只有在目标值小于某一个数的时候做标记
			if a[i] < a[mid] {
				j = mid
				tail = mid - 1
			} else {
				head = mid + 1
			}
		}

		// 以下代码用于移动元素操作
		if j >= 0 {
			tmp := a[i]
			k := i
			for k > j {
				a[k] = a[k-1]
				k--
			}
			a[k] = tmp
		}
	}

	// 打印结果
	fmt.Println("折半插入排序后的结果为：")
	algorithms.PrintArray(a)
}

/**
 * 希尔排序（缩小增量排序）
 * 算法思想：使序列变得越来越有序，从而让直接插入排序效率更高
 * 稳定性：不稳定排序
 * 时间复杂度：与增量选取有关
 * 空间复杂度：O(1)
 * 适用场景：
 */
func ShellSort(a []int) {

}
