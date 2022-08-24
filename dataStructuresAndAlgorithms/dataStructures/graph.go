package dataStructures

import "fmt"

/* 邻接矩阵存储数据结构 */
// 顶点数据结构
type VertexType struct {
	NO   int32       // 顶点编号
	Info interface{} // 顶点信息（视情况而定）
}

// 使用邻接矩阵作为存储结构的图的数据结构
type MGraph struct {
	Edges [][]int32    // 邻接矩阵定义
	N     int32        // 顶点数
	E     int32        // 边数
	Vex   []VertexType // 存放顶点信息
}

/* 邻接表存储数据结构 */
// 边表结点
type ArcNode struct {
	AdjVex  int32       // 该边所指向的结点的位置
	NextArc *ArcNode    // 指向下一条边的指针
	Info    interface{} // 该边的相关信息（如权值），这一句用得不多，题目不做相关要求可以不写
}

// 表头结点
type VNode struct {
	Data     string   // 顶点信息
	Count    int32    // 顶点当前的入度（拓扑排序）
	FirstArc *ArcNode // 指向第一条边的指针
}

// 使用邻接表作为存储结构的图的数据结构
type AGraph struct {
	AdjList []VNode // 邻接表
	N       int32   // 顶点数
	E       int32   // 边数
}

/* ============================树的遍历 start ============================ */
/* 以下两种遍历都是连通图的遍历，若图是非连通的，只需循环调用这两种遍历即可 */
// 以邻接表作为存储结构的图的深度优先搜索遍历（DFS）的递归算法
var visitDFSRec []int32 // 顶点的访问标记（0：未访问；1：已访问）
func DFSRec(G *AGraph, v int32) {
	// 1.标记当前顶点状态为已访问
	visitDFSRec = make([]int32, G.N)
	visitDFSRec[v] = 1
	visitGraph(v)
	// 2.若当前顶点v所指向的第一条边不为空，则以这一条边所指向的顶点为起始点进行DFS
	p := G.AdjList[v].FirstArc // p指向顶点v的第一条边
	for p != nil {
		// 3.当且仅当该顶点尚未访问时才递归进行DFS
		if visitDFSRec[p.AdjVex] == 0 {
			DFSRec(G, p.AdjVex)
		}
		// 4.与当前顶点相连的每一条边都要进行DFS
		p = p.NextArc // p指向与顶点v的下一条边
	}
}

// 以邻接表作为存储结构的图的深度优先搜索遍历（DFS）的非递归算法
var visitDFSNoRec []int // 顶点的访问标记（0：未访问；1：已访问）
func DFSNoRec(G *AGraph, v int) {
	// 辅助栈，记录访问过程中的顶点
	var st []int = make([]int, G.N)
	var top = -1

	// 初始化数据
	visitDFSNoRec = make([]int, G.N)

	// 1.当前顶点入栈并访问
	visitGraph(v)        // 访问顶点操作
	visitDFSNoRec[v] = 1 // 标记已经访问过的顶点
	top++
	st[top] = v

	// 2.栈不空的时候进行遍历
	for top != -1 {
		// 3.栈顶顶点出栈并且让其尚未访问的到的某一个连接的顶点入栈
		tmp := st[top]                // 取栈顶元素
		p := G.AdjList[tmp].FirstArc // p指向该顶点的第一条边
		/* 下面这个循环是p沿着边行走并将图中经过的顶点入栈的过程 */
		// 找到当前顶点第一个没有访问过的邻接顶点或者p走到当前链表尾部时，循环停止
		for p != nil && visitDFSNoRec[p.AdjVex] == 1 {
			p = p.NextArc
		}
		// 如果p到达链表尾部，则说明当前顶点的所有点都访问完毕，当前顶点出栈；否则访问当前顶点并入栈
		if p == nil {
			top--
		} else {
			visitGraph(p.AdjVex)
			visitDFSNoRec[p.AdjVex] = 1
			top++
			st[top] = int(p.AdjVex)
		}
	}
}

// 以邻接表作为存储结构的图的广度优先搜索遍历（BFS）算法
/*
注意：选择在入队时进行访问而不是出队时进行访问的好处是可防止顶点重复进队，如果图是个环，可能引起死循环
*/
var visitBFS []int32 // 顶点的访问标记（0：未访问；1：已访问）
func BFS(G *AGraph, v int32) {
	// 1.初始化一个队列，用来辅助完成图的BFS（类似于二叉树的层次遍历）
	var que []int32 = make([]int32, G.N) // 使用循环队列
	var front int32 = 0                  // 队头
	var rear int32 = 0                   // 队尾
	// 2.先让当前顶点v入队并访问之
	rear++
	que[rear] = v
	visitBFS[v] = 1
	visitGraph(v)
	// 3.在队列不为空的时候进行遍历
	for front != rear {
		// 4.队头顶点出队
		front++
		j := que[front]
		// 5.将刚出队顶点的所有边所指向的尚未被访问的顶点依次入队
		p := G.AdjList[j].FirstArc // p指向该顶点的第一条边
		for p != nil {
			// 6.如果该顶点尚未访问，则访问之且有资格入队
			if visitBFS[p.AdjVex] == 0 {
				visitBFS[p.AdjVex] = 1
				visitGraph(p.AdjVex)
				rear++
				que[rear] = p.AdjVex
			}
			// 7.p依次指向与顶点相连的所有边
			p = p.NextArc
		}
	}
}

/* ============================树的遍历 end ============================ */

// 操作遍历顶点（此处只做简单打印）
func visitGraph(data interface{}) {
	fmt.Printf("%v, ", data)
}
