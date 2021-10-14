package dataStructures

// 优先级队列简单实现
// 一个很实用的功能是：插入或删除元素的时候，元素会自动排序。

// 二叉堆：其实就是一种特殊的二叉树（完全二叉树），只不过存储在数组里。分为最大堆和最小堆。
// 最大堆的性质是：每个节点都大于等于它的两个子节点。

// 一般的链表二叉树，我们操作节点的指针，而在数组里面，我们把数组索引作为指针。

// 自定义类型别名
type Key uint32

// 优先级队列（最大堆实现）
type MaxPQ struct {
    // 存储元素的数组
    pq []Key
    // 当前 Priority Queue 中元素的个数
    N int32
}

// 优先级队列初始化
func NewMaxPQ(cap int32) *MaxPQ {
    return &MaxPQ{
        pq: make([]Key, 0, cap+1),
        N:  0,
    }
}

// 返回当前队列中最大元素
func (mpq *MaxPQ) max() Key {
    return mpq.pq[1]
}

// 插入元素 e
func insert(e Key) {

}

// 删除并返回当前队列中最大元素
func delMax() Key {
    return 0
}

// 上浮第 k 个元素，以维护最大堆性质
func swim(k Key) {

}

// 下沉第 k 个元素，以维护最大堆性质
func sink(k Key) {

}

// 交换数组的两个元素
func (mpq *MaxPQ) exch(i, j int32) {
    temp := mpq.pq[i]
    mpq.pq[i] = mpq.pq[j]
    mpq.pq[j] = temp
}

// 判断 pq[i] 是否比 pq[j] 小?
func (mpq *MaxPQ) less(i, j int32) bool {
    return mpq.pq[i] < mpq.pq[j]
}

// 父节点的索引
func parent(root int32) int32 {
    return root / 2
}

// 左孩子的索引
func left(root int32) int32 {
    return root * 2
}

// 右孩子的索引
func right(root int32) int32 {
    return root*2 + 1
}
