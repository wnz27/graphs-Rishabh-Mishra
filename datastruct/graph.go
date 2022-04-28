package datastruct

import (
	"fmt"
	"math"
	"sync"
)

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
// mean O(1)
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
// add two directions
// O(1)
func (g *ItemGraph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	ed2 := Edge{
		Node:   n1,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
	g.lock.Unlock()
}

// getShortestPath dijkstra implement
// return paths, all weight
func getShortestPath(startNode *Node, endNode *Node, g *ItemGraph) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	// 存前一个点
	prev := make(map[string]string)
	//pq := make(PriorityQueue, 1)
	//heap.Init(&pq)
	q := NodeQueue{}
	pq := q.NewQ()
	startVertex := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	// 初始化无限长
	for _, node := range g.Nodes {
		dist[node.Value] = math.MaxInt64
	}
	// 起点到起点的距离为0
	dist[startNode.Value] = startVertex.Distance
	// 起点入队
	pq.Enqueue(startVertex)
	// 不断出队
	for !pq.IsEmpty() {
		// 拿出一个最近的一个向量
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			// 如果已经遍历过这个点了则看下一个
			continue
		}
		// 标记遍历该点
		visited[v.Node.Value] = true

		// 拿到这个点关联的所有边
		nearEdges := g.Edges[*v.Node]
		// 遍历这些边
		for _, edge := range nearEdges {
			// 处理没有遍历过的点
			if !visited[edge.Node.Value] {
				// 如果起点到当前点的距离 + 当前点的一条边的距离 < 起点到当前边的目标点的距离
				if dist[v.Node.Value]+edge.Weight < dist[edge.Node.Value] {
					// 起点到 edge.Node 的向量
					store := Vertex{
						// 目标点
						Node: edge.Node,
						// 最短距离
						Distance: dist[v.Node.Value] + edge.Weight,
					}
					// 距离存为最短
					dist[edge.Node.Value] = dist[v.Node.Value] + edge.Weight
					// 存下来前一个点
					prev[edge.Node.Value] = v.Node.Value
					// edge.Node 属于最短， 当前向量入队, 后面拿出来处理以它为起点的边
					pq.Enqueue(store)
				}
			}
		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	// 拿出结尾的前一个点
	pathval := prev[endNode.Value]
	// 倒置的 路径列表
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	fmt.Println(finalArr)
	// 翻转线路是我们最终需要的
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]

}

func CreateGraph(input InputGraph) *ItemGraph {
	var g ItemGraph
	nodesMap := make(map[string]*Node)

	for _, inputData := range input.Graph {

		if _, found := nodesMap[inputData.Source]; !found {
			nA := Node{inputData.Source}
			nodesMap[inputData.Source] = &nA
			g.AddNode(&nA)
		}

		if _, found := nodesMap[inputData.Destination]; !found {
			nA := Node{inputData.Destination}
			nodesMap[inputData.Destination] = &nA
			g.AddNode(&nA)
		}

		g.AddEdge(nodesMap[inputData.Source], nodesMap[inputData.Destination], inputData.Weight)
	}
	return &g
}

func GetShortestPath(from, to string, g *ItemGraph) *APIResponse {
	nA := &Node{from}
	nB := &Node{to}

	path, distance := getShortestPath(nA, nB, g)
	return &APIResponse{
		Path:     path,
		Distance: distance,
	}
}

func GetShortestPathWithInput(input InputGraph) *APIResponse {
	g := CreateGraph(input)
	nA := &Node{input.From}
	nB := &Node{input.To}

	path, distance := getShortestPath(nA, nB, g)
	return &APIResponse{
		Path:     path,
		Distance: distance,
	}
}

func BuildInputGraph() *InputGraph {
	return &InputGraph{}
}
