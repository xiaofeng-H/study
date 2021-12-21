package algorithms

// 在并查集中查找顶点v的根结点
func GetRoot(v int, parents []int) int {
	// 顶点v的根结点
	var root int
	root = v
	for parents[root] != -1 {
		root = parents[root]
	}
	return root
}

// 合并并查集
func Union(x int, y int, parents []int) int {
	// x, y的根结点
	xRoot := GetRoot(x, parents)
	yRoot := GetRoot(y, parents)

	if xRoot == yRoot {
		return 0
	} else {
		// 合并并查集
		parents[xRoot] = yRoot
		return 1
	}
}

// 初始化并查集
func InitialParents(parents []int) {
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}
}
