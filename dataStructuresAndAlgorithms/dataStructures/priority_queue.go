package dataStructures

/*
优先级队列简单实现
一个很实用的功能：插入或删除元素的时候，元素会自动排序。
两个主要的API：insert插入一个元素和delMax删除最大元素（如果底层是用最小堆，那么就是delMin）
底层数据结构：二叉堆：其实就是一种特殊的二叉树（完全二叉树），只不过存储在数组里。分为最大堆和最小堆。
            最大堆的性质是：每个节点都大于等于它的两个子节点。一般的链表二叉树，我们操作节点的指针，而在数组里面，我们把数组索引作为指针。
*/

// 优先级队列（最大堆实现）
type MaxPQ struct {
	// 存储元素的数组
	pq []int
	// 当前 Priority Queue 中元素的个数
	N int
}

// 构造函数
func NewMaxPQ(cap int) *MaxPQ {
	return &MaxPQ{
		// 索引0不用，所以多分配一个空间
		pq: make([]int, cap+1),
		N:  0,
	}
}

// 返回当前队列中最大元素
func (mpq *MaxPQ) max() int {
	return mpq.pq[1]
}

/*
insert方法先把要插入的元素添加到堆底的最后，然后让其上浮到正确的位置
时间复杂度：O(logK)（K为当前二叉堆（优先级队列）中元素的总数）
*/
func (mpq *MaxPQ) insert(e int) {
	mpq.N++
	// 先把新元素加到最后
	mpq.pq = append(mpq.pq, e)
	// 然后让它上浮到正确的位置
	mpq.swim(mpq.N)
}

/*
删除并返回当前队列中最大元素
delMax方法先把堆顶元素A和堆底最后的元素B对调，然后删除A，最后让B下沉到正确位置
时间复杂度：O(logK)（K为当前二叉堆（优先级队列）中元素的总数）
*/
func (mpq *MaxPQ) delMax() int {
	// 最大堆堆顶就是最大元素
	max := mpq.pq[1]
	// 把这个最大元素换到最后，删除之
	mpq.exch(1, mpq.N)
	mpq.pq = mpq.pq[:mpq.N]
	mpq.N--
	// 让pq[1]下沉到正确位置
	mpq.sink(1)
	return max
}

// 上浮第 k 个元素，以维护最大堆性质
func (mpq *MaxPQ) swim(k int) {
	// 如果浮到堆顶，就不能再上浮了
	for k > 1 && mpq.less(parent(k), k) {
		// 如果第k个元素比上层大，将k换上去
		mpq.exch(parent(k), k)
		k = parent(k)
	}
}

// 下沉第 k 个元素，以维护最大堆性质
func (mpq *MaxPQ) sink(k int) {
	// 如果沉到堆底，就沉不下去
	for left(k) <= mpq.N {
		// 先假设左边结点较大
		older := left(k)
		// 如果右边结点存在，比一下大小
		if right(k) <= mpq.N && mpq.less(older, right(k)) {
			older = right(k)
		}
		// 结点k比两孩子都大，就不必下沉了
		if mpq.less(older, k) {
			break
		}
		// 否则，不符合最大堆结构，下沉k结点
		mpq.exch(k, older)
		k = older
	}
}

// 交换数组的两个元素
func (mpq *MaxPQ) exch(i, j int) {
	temp := mpq.pq[i]
	mpq.pq[i] = mpq.pq[j]
	mpq.pq[j] = temp
}

// 判断 pq[i] 是否比 pq[j] 小?
func (mpq *MaxPQ) less(i, j int) bool {
	return mpq.pq[i] < mpq.pq[j]
}

// 父节点的索引
func parent(root int) int {
	return root / 2
}

// 左孩子的索引
func left(root int) int {
	return root * 2
}

// 右孩子的索引
func right(root int) int {
	return root*2 + 1
}
