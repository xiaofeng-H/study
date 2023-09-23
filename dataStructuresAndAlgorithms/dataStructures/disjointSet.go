package dataStructures

/*======================================Union-Find 算法 start============================================*/
/*
关键点：
1、⽤  parent  数组记录每个节点的⽗节点，相当于指向⽗节点的指针，所以  parent  数组内实际存储着⼀个森林（若⼲棵多叉树）。
2、⽤  size  数组记录着每棵树的重量，⽬的是让  union  后树依然拥有平衡性，⽽不会退化成链表，影响操作效率。
3、在  find  函数中进⾏路径压缩，保证任意树的⾼度保持在常数，使得 union  和  connected  API 时间复杂度为 O(1)。
*/
// 并查集结构体
type UnionFindSet struct {
	// 记录连通分量
	counts int
	// 结点x的父结点是parent[x]
	parents []int
	// 记录树的“重量”（树的结点数量）
	size []int
}

/*
构造函数
时间复杂度：O(N)
空间复杂度：O(N)
*/
func NewUF(n int) *UnionFindSet {
	// 数据初始化
	tmp1, tmp2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		// 父结点指针初始化指向自己
		tmp1[i] = i
		// 最初每棵树只有一个结点，重量应该初始化为1
		tmp2[i] = 1
	}

	return &UnionFindSet{
		// 一开始互不连通
		counts:  n,
		parents: tmp1,
		size:    tmp2,
	}
}

/*
将结点p、q连接使其连通
时间复杂度：O(logN) || 进行路径压缩后优化为O(1)
*/
func (u *UnionFindSet) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	// p、q本就是连通的，则什么都不做
	if rootP == rootQ {
		return
	}

	// 将两棵树合并为一颗
	// 连接两颗树不做优化可能会使树极度不平衡，从而增加Find方法的时间复杂度
	/* u.parents[rootP] = rootQ
	   u.parents[rootQ] = rootP 也一样 */
	// 小树接到大树下面，较平衡
	if u.size[rootP] > u.size[rootQ] {
		u.parents[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	} else {
		u.parents[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	}

	// 两个分量合二为一
	u.counts--
}

/*
返回某个结点x的根结点
时间复杂度：O(logN) || 进行路径压缩后优化为O(1)
路径压缩结果：可⻅，调⽤ find 函数每次向树根遍历的同时，顺⼿将树⾼缩短了，最终所有树⾼都不会超过 3（ union 的时候树⾼可能达到 3）。
*/
func (u *UnionFindSet) Find(x int) int {
	// 根结点的parent[x] == x
	for u.parents[x] != x {
		// 优化操作：进行路径压缩，使得相关方法的时间复杂度优化为O(1)
		u.parents[x] = u.parents[u.parents[x]]
		x = u.parents[x]
	}
	return x
}

// 返回当前的连通分量个数
func (u *UnionFindSet) Counts() int {
	return u.counts
}

/*
判断结点p、q是否连通
时间复杂度：O(logN)/进行路径压缩后优化为O(1)
*/
func (u *UnionFindSet) Connected(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	return rootP == rootQ
}

// 返回并查集
func (u *UnionFindSet) GetUnionFindSet() []int {
	return u.parents
}

/*======================================Union-Find 算法 end============================================*/
