package algorithms

/*
LRU算法详解
1.为什么必须要用双向链表？
	答：因为我们需要删除操作。删除一个结点不光要得到该结点本身的指针，也需要操作其
前驱结点的指针，而双向链表才能支持直接查找前驱，保证操作的时间复杂度O(1)。
2.为什么要在链表中同时存储key和val,而不是只存储val？
	答：当缓存容量已满，我们不仅仅要删除最后一个Node结点，还要把map中映射到该结点的key同时
删除，而这个key只能由Node得到。如果Node结构中只存储了val，那么我们就无法得知key是什么，就无法删除
map中的键，造成错误。
*/
/*
双链表结点
*/
type Node struct {
	key, val  int
	pre, next *Node
}

func NewNode(k, v int) *Node {
	return &Node{
		key:  k,
		val:  v,
		pre:  nil,
		next: nil,
	}
}

// 获取Node的key
func (n *Node) GetKey() int {
	return n.key
}

/*
双链表
*/
type DoubleList struct {
	// 头结点
	Root Node
	// 链表长度
	len int
}

func (dl *DoubleList) GeetRoot() *Node {
	return &dl.Root
}

// 初始化方法
func (dl *DoubleList) Init() *DoubleList {
	dl.Root.next = &dl.Root
	dl.Root.pre = &dl.Root
	dl.len = 0
	return dl
}

func NewDoubleList() *DoubleList {
	return new(DoubleList).Init()
}

// 在链表头部添加结点x，时间复杂度O(1)
func (dl *DoubleList) AddFirst(x *Node) {
	x.next = dl.Root.next
	dl.Root.next.pre = x
	x.pre = &dl.Root
	dl.Root.next = x
	dl.len++
}

// 删除链表中的x结点（x一定存在）
// 由于是双链表且给的是Node结点，时间复杂度O(1)
func (dl *DoubleList) Remove(x *Node) {
	x.pre.next = x.next
	x.next.pre = x.pre
	x.next = nil // 防止内存泄露
	x.pre = nil  // 防止内存泄露
	dl.len--
}

// 删除链表中的最后一个结点，并返回该结点，时间复杂度O(1)
func (dl *DoubleList) RemoveLast() *Node {
	x := dl.Root.pre
	x.pre.next = &dl.Root
	dl.Root.pre = x.pre
	x.next = nil // 防止内存泄漏
	x.pre = nil  // 防止内存泄漏
	return x
}

// 返回链表的长度，时间复杂度O(1)
func (dl *DoubleList) Len() int {
	return dl.len
}

type LRUCache struct {
	// key映射到Node(key,val)
	m map[int]*Node
	// Node(k1,v1)<->Node(k2,v2)...
	cache *DoubleList
	// 最大容量
	cap int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		m:     make(map[int]*Node),
		cache: NewDoubleList(),
		cap:   capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; !ok { // key不存在
		return -1
	} else {
		val := node.val
		// 将数据(key,val)提到开头
		lru.Put(key, val)
		return val
	}
}

func (lru *LRUCache) Put(key, val int) {
	// 先把新结点x做出来
	x := NewNode(key, val)
	if node, ok := lru.m[key]; ok { // key已经存在
		// 删除旧的结点，新的插到头部
		lru.cache.Remove(node)
		lru.cache.AddFirst(x)
		// 更新map中对应的数据
		lru.m[key] = x
	} else {
		if lru.cap == lru.cache.Len() { // cache已满
			// 删除链表的最后一个数据腾位置
			last := lru.cache.RemoveLast
			// 删除map中的映射到该数据的键
			delete(lru.m, last().key)
		}
		// 将新结点x插入到开头
		lru.cache.AddFirst(x)
		// map中新建key对新结点x的映射
		lru.m[key] = x
	}
}
