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
	Edges [MaxSize][MaxSize]int32 // 邻接矩阵定义
	N     int32                   // 顶点数
	E     int32                   // 边数
	Vex   [MaxSize]VertexType     // 存放顶点信息
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
	FirstArc *ArcNode // 指向第一条边的指针
}

// 使用邻接表作为存储结构的图的数据结构
type AGraph struct {
	AdjList [MaxSize]VNode // 邻接表
	N       int32          // 顶点数
	E       int32          // 边数
}

/* 一下两种遍历都是连通图的遍历，若图是非连通的，只需循环调用这两种遍历即可 */
// 以邻接表作为存储结构的图的深度优先搜索遍历（DFS）的递归算法
var visitDFS [MaxSize]int32 // 顶点的访问标记（0：未访问；1：已访问）
func DFS(G *AGraph, v int32) {
	// 1.标记当前顶点状态为已访问
	visitDFS[v] = 1
	visitGraph(v)
	// 2.若当前顶点v所指向的第一条边不为空，则以这一条边所指向的顶点为起始点进行DFS
	p := G.AdjList[v].FirstArc // p指向顶点v的第一条边
	for p != nil {
		// 3.当且仅当该顶点尚未访问时才递归进行DFS
		if visitDFS[p.AdjVex] == 0 {
			DFS(G, p.AdjVex)
		}
		// 4.与当前顶点相连的每一条边都要进行DFS
		p = p.NextArc // p指向与顶点v的下一条边
	}
}

// 以邻接表作为存储结构的图的广度优先搜索遍历（BFS）算法
/*
注意：选择在入队时进行访问而不是出队时进行访问的好处是可防止顶点重复进队，如果图是个环，可能引起死循环
*/
var visitBFS [MaxSize]int32 // 顶点的访问标记（0：未访问；1：已访问）
func BFS(G *AGraph, v int32) {
	// 1.初始化一个队列，用来辅助完成图的BFS（类似于二叉树的层次遍历）
	var que [MaxSize]int32 // 使用循环队列
	var front int32 = 0    // 队头
	var rear int32 = 0     // 队尾
	// 2.先让当前顶点v入队并访问之
	rear = (rear + 1) % MaxSize
	que[rear] = v
	visitBFS[v] = 1
	visitGraph(v)
	// 3.在队列不为空的时候进行遍历
	for front != rear {
		// 4.队头顶点出队
		front = (front + 1) % MaxSize
		j := que[front]
		// 5.将刚出队顶点的所有边所指向的尚未被访问的顶点依次入队
		p := G.AdjList[j].FirstArc // p指向该顶点的第一条边
		for p != nil {
			// 6.如果该顶点尚未访问，则访问之且有资格入队
			if visitBFS[p.AdjVex] == 0 {
				visitBFS[p.AdjVex] = 1
				visitGraph(p.AdjVex)
				rear = (rear + 1) % MaxSize
				que[rear] = p.AdjVex
			}
			// 7.p依次指向与顶点相连的所有边
			p = p.NextArc
		}
	}
}

// 操作遍历顶点（此处只做简单打印）
func visitGraph(data interface{}) {
	fmt.Printf("%v, ", data)
}
