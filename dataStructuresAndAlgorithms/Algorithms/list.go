package Algorithms

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
)

/*
算法描述：数组A总共有m+n个元素，前m个元素递增有序，后n个元素也递增有序。
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
算法描述：两个递增有序集合A和B，求A-B。
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
