package main

import "fmt"

type Graph struct {
	Nodes []string
	NodesNum int
	Edges [][]int
	EdgesNum int
	IsDag bool
	Visited []int
	ParallelRunRawMap map[string][]string
	ParallelRunFinalMap map[string][][]string
	
}

func (g *Graph) indexOfNodes(node string) int {
	for i:=0;i<len(g.Nodes);i++ {
		if node == g.Nodes[i] {
			return i
		}
	}
	return -1
}

func (g *Graph) updateEdge() {
	if len(g.Edges)  != len(g.Nodes) {
		var newEdge = make([][]int, 0)
		
		for i:=0;i<len(g.Nodes);i++ {
			eachEdge := make([]int, len(g.Nodes))
			newEdge = append(newEdge, eachEdge)
		}
		g.Edges = newEdge
	} else {
		return
	}
}

func (g *Graph) updateVisited() {
	g.Visited = make([]int, len(g.Nodes))
}


func (g *Graph) addEdge(fromNode, toNode string) bool {
	g.updateVisited()
	g.updateEdge()
	if g.indexOfNodes(fromNode) != -1 && g.indexOfNodes(toNode) != -1 {
		g.Edges[g.indexOfNodes(fromNode)][g.indexOfNodes(toNode)] = 1
		return true
	}
	return false
}

func (g *Graph) addNode(node string) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) isDag() bool {
	if !g.IsDag {
		return false
	}
	for i:=0;i<len(g.Nodes);i++{
		if g.Visited[i] == -1 {
			continue
		}
		g.dfs(i)
		if !g.IsDag {
			return false
		}
	}
	return true
}

func (g *Graph) dfs(index int) {
	g.Visited[index] = 1
	for i:=0;i<len(g.Nodes);i++ {
		if g.Edges[index][i] != 0 {
			if g.Visited[i] == 1 {
				g.IsDag = false
				break
			} else if g.Visited[i] == -1 {
				continue
			} else {
				g.dfs(i)
			}
		}
	}
	g.Visited[index] = -1
}

func (g *Graph) forEach() []string {
	if !g.isDag() {
		return nil
	}
	nodeNum, restNodeNum := len(g.Nodes), len(g.Nodes)
	result := make([]string,0)
	color := make([]int, nodeNum)
	for restNodeNum > 0 {
		removeNode := make([]int, 0)
		for i := 0; i < len(g.Nodes); i++ {
			if color[i] == -1 {
				continue
			}
			counter := 0
			for j := 0; j < len(g.Nodes); j++ {
				counter += g.Edges[j][i]
			}
			if counter == 0 {
				color[i] = -1
				removeNode = append(removeNode, i)
			}
		}
		
		for i := 0; i < len(removeNode); i++ {
			for j := 0; j < len(g.Nodes); j++ {
				g.Edges[removeNode[i]][j] = 0
			}
			result = append(result, g.Nodes[removeNode[i]])
			restNodeNum--
		}
	}
	return result
}

func(g *Graph) getParallelRunStartNodesList() {
	startNodeList := make([]string,0)
	for i:=0;i<len(g.Nodes);i++ {
		count :=0
		for j:=0;j<len(g.Nodes);j++ {
			count += g.Edges[j][i]
		}
		if count == 0 {
			startNodeList = append(startNodeList, g.Nodes[i])
		}
	}
	for _,v := range startNodeList {
		g.ParallelRunRawMap[v] = make([]string,0)
	}
}

func(g *Graph) getParallelRunRawLists() {
	for k, v := range  g.ParallelRunRawMap {
		g.getParallelRunRawListForEachNode(k,&v)
		g.ParallelRunRawMap[k] = append(g.ParallelRunRawMap[k],k)
		g.ParallelRunRawMap[k] = append(g.ParallelRunRawMap[k],v...)
	}
}

func(g *Graph) getParallelRunRawListForEachNode(node string,taskList *[]string) {
	index := g.indexOfNodes(node)
	for i:=0;i<len(g.Nodes);i++ {
		if g.Edges[index][i] != 0 {
			*taskList = append(*taskList, g.Nodes[i])
			g.getParallelRunRawListForEachNode(g.Nodes[i], taskList)
		}
	}
}

func (g *Graph) getParallFinalLists() {
	g.getParallelRunStartNodesList()
	g.getParallelRunRawLists()
	for k,v := range g.ParallelRunRawMap {
		var tempNodes = []int{}
		for i, j := range v {
			if g.Edges[g.indexOfNodes(v[0])][g.indexOfNodes(j)] == 1 {
				tempNodes  = append(tempNodes, i)
			}
		}
		var tempParallelRunLists = make([][]string,0)
		if len(tempNodes) == 1 {
			tempParallelRunLists = append(tempParallelRunLists, v)
		} else {
			for i:=0;i<len(tempNodes);i++ {
				if i == 0 {
					tempParallelRunLists = append(tempParallelRunLists, v[:tempNodes[i+1]])
				} else if i < len(tempNodes) -1 {
					var tempList = make([]string,0)
					tempList = append(tempList, v[0])
					tempList = append(tempList, v[tempNodes[i]:tempNodes[i+1]]...)
					tempParallelRunLists = append(tempParallelRunLists,tempList)
					
				} else if i == len(tempNodes) -1 {
					var tempList = make([]string,0)
					tempList = append(tempList, v[0])
					tempList = append(tempList, v[tempNodes[i]:len(v)]...)
					tempParallelRunLists = append(tempParallelRunLists,tempList)
				}
			}
		}
		g.ParallelRunFinalMap[k] = tempParallelRunLists
	}
}


func newGraph() *Graph {
	nodes := make([]string,0)
	edges := make([][]int,0)
	visited := make([]int, 0)
	parallelRunRawMap := make(map[string][]string)
	parallelRunFinalMap := make(map[string][][]string)
	
	return &Graph{
		Nodes: nodes,
		Edges: edges,
		IsDag: true,
		Visited: visited,
		ParallelRunRawMap: parallelRunRawMap,
		ParallelRunFinalMap: parallelRunFinalMap,
	}
}

func main() {
	g := newGraph()
	// 添加顶点
	g.addNode("a")
	g.addNode("b")
	g.addNode("c")
	g.addNode("d")
	g.addNode("e")
	g.addNode("f")
	g.addNode("g")
	g.addNode("h")
	g.addNode("i")
	g.addNode("j")
	// 默认初始化10个顶点，此处扩容一次
	g.addNode("k")
	// 添加边
	g.addEdge("a", "b")
	g.addEdge("b", "c")
	g.addEdge("c", "d")
	g.addEdge("e", "f")
	g.addEdge("f", "k")
	g.addEdge("k", "g")
	g.addEdge("a","g")
	g.addEdge("g","j")
	g.addEdge("h","i")
	g.addEdge("a","i")
	
	g.getParallFinalLists()
	fmt.Println(g.ParallelRunFinalMap)
	fmt.Println(g.ParallelRunRawMap)
	
	
}
