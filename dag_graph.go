package main

import "fmt"

type Graph struct {
	Nodes []string
	NodesNum int
	Edges [][]int
	EdgesNum int
	IsDag bool
	Visited []int
	TaskRunMap map[string][]string
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

func(g *Graph) taskRunStartNodesList() {
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
		g.TaskRunMap[v] = make([]string,0)
	}
}

func(g *Graph) taskRunList() {
	for k, v := range  g.TaskRunMap {
		g.getTaskRunListForEach(k,&v)
		g.TaskRunMap[k] = append(g.TaskRunMap[k],k)
		g.TaskRunMap[k] = append(g.TaskRunMap[k],v...)
		
	}
	fmt.Println(g.TaskRunMap)
}

func(g *Graph) getTaskRunListForEach(node string,taskList *[]string) {
	index := g.indexOfNodes(node)
	for i:=0;i<len(g.Nodes);i++ {
		if g.Edges[index][i] != 0 {
			*taskList = append(*taskList, g.Nodes[i])
			g.getTaskRunListForEach(g.Nodes[i], taskList)
		}
	}
	
}

func newGraph() *Graph {
	nodes := make([]string,0)
	edges := make([][]int,0)
	visited := make([]int, 0)
	taskRunMap := make(map[string][]string)
	return &Graph{
		Nodes: nodes,
		Edges: edges,
		IsDag: true,
		Visited: visited,
		TaskRunMap: taskRunMap,
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
	
	
	g.taskRunStartNodesList()
	g.taskRunList()
	fmt.Println(g.TaskRunMap)
}