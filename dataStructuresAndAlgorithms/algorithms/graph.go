package algorithms

import (
	"fmt"
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

// 最短路径：迪杰斯特拉算法
/*
1.数据结构：邻接矩阵
2.时间复杂度：O(n^2)
3.运行结果：函数结束时，dist[]数组中存放了v0点到其余顶点的最短路径长度，path[]中存放了v0点到其余顶点的最短路径
*/
func Dijkstra(g *ds.MGraph, v0 int32) ([]int32, []int32) {
	// 记录当前已找到的从v0到每个终点vi的最短路径长度
	// 其初始状态为：若v0到vi有边，则dist[i]为边上的权值，否则置为无穷大
	var dist []int32 = make([]int32, 0)
	// 顶点标记数组，vSet[i]=0表示顶点i已经被并入最短路径，vSet[i]=1则表示顶点i还未被并入最短路径
	var vSet []int32 = make([]int32, 0)
	// 保存从v0到vi最短路径上vi的前一个顶点，其初始状态为：若v0到vi有边，则path[vi]=v0，否则path[vi]=-1
	var path []int32 = make([]int32, 0)

	// 1.数据初始化
	for i := int32(0); i < g.N-1; i++ {
		if g.Edges[v0][i] == 0 {
			dist[i] = INF
			path[i] = -1
		} else {
			dist[i] = g.Edges[v0][i]
			path[i] = v0
		}
		vSet[i] = 0
	}

	// 2.以v0为起始点，开始寻找各个顶点的最短路径
	vSet[v0] = 1 // 顶点v0作为起始点，第一个被并入最短路径
	for i := int32(0); i < g.N; i++ {
		min := int32(INF) // 权值最小的下一条路径上的权值
		var v = int32(0)  // 下一个将被并入最段路径的顶点
		// 3.该循环每次从剩余顶点中选出一个顶点，通往这个顶点的路径在通往所有剩余顶点的路径中是长度最短的
		for i := int32(0); i < g.N; i++ {
			if vSet[i] == 0 && dist[i] < min {
				min = dist[i] // 记录最小权值
				v = i         // 记录权值最小的顶点的编号
			}
		}

		// 4.将选出的顶点并入最短路径中，同时更新各个辅助数组的数据
		vSet[v] = 1
		// 该循环以刚并入的顶点作为中间点，对所有通往剩余顶点的路径进行检测
		for i := int32(0); i < g.N; i++ {
			// 判断顶点v的加入是否会出现通往顶点i的更短路径，如果出现，则改变原来路径及其长度
			if vSet[i] == 0 && dist[i] > dist[v]+g.Edges[v][i] {
				dist[i] = dist[v] + g.Edges[v][i]
				path[i] = v
			}
		}
	}
	// 5.返回最短路径及路径数组
	return dist, path
}

// path[]数组最短路径打印
func PrintPath(path [MAXSIZE]int32, a int32) {
	// 借助栈来实现逆向输出
	var stack [MAXSIZE]int32
	var top = int32(-1)
	// 该循环以由叶子结点到根结点的顺序将其入栈
	for path[a] != -1 {
		top++
		stack[top] = a
		a = path[a]
	}
	top++
	stack[top] = a
	// 打印
	for top != -1 {
		v := stack[top]
		top--
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
}

// 最短路径：佛洛伊德算法
/*
1.数据结构：邻接矩阵
2.时间复杂度：O(n^3)
*/
func Floyd(g *ds.MGraph) ([][]int32, [][]int32) {
	// arr用来记录当前已经求得的任意两个顶点最短路径的长度
	var arr [][]int32 = make([][]int32, 0)
	// path用来记录当前两顶点间最短路径上要经过的中间顶点
	var path [][]int32 = make([][]int32, 0)

	// 1.数据初始化
	for i := int32(0); i < g.N; i++ {
		for j := int32(0); j < g.N; j++ {
			arr[i][j] = g.Edges[i][j]
			path[i][j] = -1
		}
	}

	// 2.以k为中间点，对所有的顶点对<i,j>进行检测
	for k := int32(0); k < g.N; k++ {
		for i := int32(0); i < g.N; i++ {
			for j := int32(0); j < g.N; j++ {
				if arr[i][j] > arr[i][k]+arr[k][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
					path[i][j] = k
				}
			}
		}
	}

	// 3.返回最短路径及路径
	return arr, path
}

// 拓扑排序
func TopSort(g *ds.AGraph) int32 {
	// 定义并初始化栈
	var stack [MAXSIZE]int32
	var top = int32(-1)
	// 计数器：记录当前已经被排序的顶点个数
	var num = int32(0)

	// 1.将图中入度为0的顶点入栈
	for i := int32(0); i < g.N; i++ {
		if g.AdjList[i].Count == 0 {
			top++
			stack[top] = i
		}
	}
	// 2.关键操作
	for top != -1 {
		// 顶点出栈
		v := stack[top]
		top--
		// 计数器+1，统计当前顶点
		num++
		// 输出当前顶点
		fmt.Printf("%d, ", v)
		// 该循环将所有由此顶点引出的边所指向的顶点的入度都-1，并将这个过程中入度变为0的顶点入栈
		p := g.AdjList[num].FirstArc // p指向被排序顶点的第一条边
		k := p.AdjVex                // 该边连接的另一个顶点
		for p != nil {
			g.AdjList[k].Count--
			if g.AdjList[k].Count == 0 {
				top++
				stack[top] = k
			}
			// p指向该顶点连接的下一条边
			p = p.NextArc
		}
	}
	// 3.判断该有向图是否可以进行拓扑排序
	if num == g.N {
		return 1
	} else {
		return 0
	}
}
