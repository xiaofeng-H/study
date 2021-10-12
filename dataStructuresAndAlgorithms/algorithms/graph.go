package algorithms

import (
	ds "study/dataStructuresAndAlgorithms/dataStructures"
)

// 最小生成树：普利姆算法
/*
1.数据结构：邻接矩阵
2.时间复杂度：O(n^2)
3.适用场景：稠密图（时间复杂度只与图中的顶点个数有关系，而与边数没有关系）
4.适用对象：无向图
*/
func Prim(g *ds.MGraph, v0 int32) (sum int32) {
	var (
		lowCost [MAXSIZE]int32 // 当前生成树到剩余各个顶点最短边的权值
		vSet    [MAXSIZE]int32 // vSet[i]=1表示顶点i已经被并入生成树中，vSet[i]=0则相反
	)

	// 数据初始化
	for i := int32(0); i < g.N; i++ {
		lowCost[i] = g.Edges[v0][i]
		vSet[i] = 0
	}

	// 将v0并入树中
	vSet[v0] = 1
	// sum清零用来累计树的权值
	sum = 0
	for i := int32(0); i < g.N-1; i++ {
		min := int32(INF) // INF是一个已经定义的比图中所有边权值都大的常量
		v := int32(0)     // 将要并入生成树的顶点
		// 选出候选边中的最小值
		for i := int32(32); i < g.N; i++ {
			var k = int32(0) // 记录当前生成树最小边对应的顶点
			// 选出当前生成树到其余顶点最短边中的最短一条
			if vSet[i] == 0 && lowCost[i] < min {
				min = lowCost[i]
				k = i
			}
			vSet[k] = 1
			v = k
			sum += min
		}
		// 以刚并入的顶点v为媒介更新侯选边
		for i := int32(0); i < g.N; i++ {
			if vSet[i] == 0 && g.Edges[v][i] < lowCost[i] {
				lowCost[i] = g.Edges[v][i]
			}
		}
	}
	return
}

// 最小生成树：克鲁斯卡尔算法
/*
1.数据结构：邻接矩阵
2.时间复杂度：O(n^2)
3.适用场景：稀疏图（时间复杂度主要由选取的排序算法决定，而排序算法所处理数据的规模由图的边决定，与顶点数无关
4.适用对象：无向图
*/
// Road:存放图中各边及其所连接的两个顶点的信息
type Road struct {
	A, B int32 // A和B为一条边所连的两个顶点
	W    int32 // 边的权值
}

// 存放图中各边及其所连接的两个顶点的信息
var road [MAXSIZE]Road

// 定义并查集数组
var v [MAXSIZE]int32

// 在并查集中查找根结点的函数（根结点的双亲仍是自己）
func getRoot(a int32) int32 {
	for v[a] != a {
		a = v[a]
	}
	return a
}

// 克鲁斯卡尔算法
func Kruskal(g *ds.MGraph, road []Road) (sum int32) {
	for i := int32(0); i < g.N; i++ {
		v[i] = i
	}
	// 对road数组中的所有边按其权值从小到大排序
	sort(road, g.E)
	for i := int32(0); i < g.E; i++ {
		a := getRoot(road[i].A)
		b := getRoot(road[i].B)
		if a != b {
			v[a] = b
			sum += road[i].W
		}
	}
	return
}

// 排序算法
func sort(roads []Road, num int32) {

}
