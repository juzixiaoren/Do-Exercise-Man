package main

/*
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

class Node {
    public int val;
    public List<Node> neighbors;
}


测试用例格式：

简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。

邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。

给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。



示例 1：

*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	root := Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	cur := node
	queue := []*Node{}
	has_visit := map[int]struct{}{}
	has_visit[cur.Val] = struct{}{}
	has_create := map[int]*Node{}
	has_create[root.Val] = &root
	queue = append(queue, node)
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		cur_copy := has_create[cur.Val]
		for i := range cur.Neighbors {
			_, OK := has_visit[cur.Neighbors[i].Val]
			if !OK {
				has_visit[cur.Neighbors[i].Val] = struct{}{}
				queue = append(queue, cur.Neighbors[i])
			}
			_, OK = has_create[cur.Neighbors[i].Val]
			if !OK {
				newNode := Node{
					Val:       cur.Neighbors[i].Val,
					Neighbors: []*Node{},
				}
				has_create[newNode.Val] = &newNode
			}
			cur_copy.Neighbors = append(cur_copy.Neighbors, has_create[cur.Neighbors[i].Val])
		}
	}
	return &root
}
