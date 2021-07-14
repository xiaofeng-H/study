package Algorithms

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
)

/*
算法描述：数组A总共有m+n个元素，前m个元素递增有序，后n个元素也递增有序
算法要求：设计一个算法，是的整个顺序表有序
算法规模：时间复杂度：O(mn); 空间复杂度：O(1)
*/
func ReSort(A []int, m, n int) bool {
	if len(A) != m+n {
		fmt.Printf("The paraments are error! A`length is %d, but m=%d, n=%d\n", len(A), m, n)
		return false
	}

	for i := m; i < m+n; i++ {
		temp := A[i]
		index := -1
		for j := i - 1; j >= 0; j-- {
			if temp < A[j] {
				A[j+1] = A[j]
				index = j
			} else {
				break
			}
		}
		if index != -1 {
			A[index] = temp
		}
	}

	return true
}

/*
算法描述：两个递增有序集合A和B，求A-B
算法要求：差集保存在A中，并保持元素的递增有序性
算法规模：时间复杂度：O(m+n); 空间复杂度：O(1)
*/
func SubLinkList(A, B *dataStructures.LNode) bool {
	if A.Next == nil {
		fmt.Println("NUll LinkList!")
		return false
	}
	/*
		// 解法一：自己想的，时间复杂度：O(mn)
		p := A
		q := B.Next
		for {
			if q == nil {
				break
			}
			for {
				if p.Next.Data > q.Data || p.Next == nil {
					break
				}
				if p.Next.Data == q.Data {
					p.Next = p.Next.Next
					A.Data--
				} else {
					p = p.Next
				}
			}
			q = q.Next
		}
	*/

	// 解法二：别人的，时间复杂度：O(m+n)
	p := A
	q := B.Next
	for {
		if p.Next == nil || q == nil {
			break
		}

		if p.Next.Data < q.Data {
			p = p.Next
		} else if p.Next.Data > q.Data {
			q = q.Next
		} else {
			p.Next = p.Next.Next
			A.Data--
		}
	}

	return true
}

// “小师妹，我们来日方长呀！”
// “方长是谁？”
// “你大爷！”
// ————《江湖少年》(2021/7/14 14:47)

// @desc  	数组逆置
// @param   A
func ArrayReverse(A []int) {
	if len(A) <= 0 {
		fmt.Println("The array is null")
	}

	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		temp := A[i]
		A[i] = A[j]
		A[j] = temp
	}
}

// @desc    单链表的逆置（要求：不能建立新的结点）
// @param   ln
func LinkListReverse(ln *dataStructures.LNode) {
	if ln.Next == nil {
		fmt.Println("The linkList is null")
	}

	// 单链表使用头插法建立时，得到的表示一个顺序与原先相反的单链表，故此处使用头插法思想
	p := ln.Next
	ln.Next = nil
	var q *dataStructures.LNode // 头插法建表时使用的中间变量
	for p != nil {
		// 思考良久，如不借助第三个指针变量，无法完成头插法构建单链表操作，看了参考解法亦然（2021/7/14 16:53）
		q = p.Next
		p.Next = ln.Next
		ln.Next = p
		p = q
	}

}

// @desc    一次快速排序得到的结果
// @param   A
func OnceQuickSort(A []int) {
	if len(A) <= 0 {
		fmt.Println("The array is null")
	}

	temp := A[0]
	flag := true
	for i, j := 0, len(A)-1; i <= j; {
		if flag {
			if temp > A[j] {
				A[i] = A[j]
				A[j] = temp
				i++
				flag = false
			} else {
				j--
			}
		} else {
			if temp < A[i] {
				A[j] = A[i]
				A[i] = temp
				j--
				flag = true
			} else {
				i++
			}
		}
	}
}

/*
算法描述：有N个个位正整数存放在int型数组A[0,...,N-1]中，N为已经定义的常量且N<=9，数组A的长度为N，另给一个int型变量i。
算法要求：要求只用上述变量，求出这N个数中的最小者并不能破坏A[]中的数据。
算法规模：时间复杂度：O(N); 空间复杂度：O(1)
*/
func GetMinByI(A []int) int {
	if len(A) <= 0 {
		fmt.Println("The array is nil")
		return -1
	}

	i := A[0] // 只能使用这一个变量哦
	for ; i < 10*len(A); i += 10 {
		if A[i/10] < i%10 {
			i -= i % 10
			i += A[i/10]
		}
	}

	return i % 10
}

/*
算法描述：主元素：一个序列中，若一个数出现的次数m大于序列长度n的一半，即m>(n/2)。
算法要求：求给定数组中的主元素，没有则返回-1。
算法思路：从前向后扫描数组元素，标记出一个可能成为主元素的元素Num，然后重新计数，确认Num是否为主元素。
算法规模：时间复杂度：O(N); 空间复杂度：O(1)
*/
func IsMajority(A []int) int {
	if len(A) <= 0 {
		fmt.Println("The array is nil")
		return -1
	}

	var c int = A[0]  // 用来保存候选主元素
	var count int = 1 // 用来计数

	for i := 0; i < len(A); i++ {
		if A[i] == c {
			count++
		} else {
			if count > 0 {
				count--
			} else {
				c = A[i]
				count = 1
			}
		}

	}

	if count > 0 {
		count = 0
		for i := 0; i < len(A); i++ {
			if A[i] == c {
				count++
			}
		}
	}

	// 存在主元素
	if count > len(A)/2 {
		return c
	}
	// 不存在主元素
	return -1
}
