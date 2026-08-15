package main

type edge struct {
	to    string
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string][]edge{}
	for i := 0; i < len(equations); i++ {
		graph[equations[i][0]] = append(graph[equations[i][0]], edge{to: equations[i][1], value: values[i]})
		graph[equations[i][1]] = append(graph[equations[i][1]], edge{to: equations[i][0], value: 1 / values[i]})
	}
	ans := []float64{}
	for i := 0; i < len(queries); i++ {
		ans = append(ans, bfs(graph, queries[i][0], queries[i][1]))
	}
	return ans
}

func bfs(graph map[string][]edge, start string, target string) float64 {
	// 起点不存在，说明 start 是未知变量
	_, ok := graph[start]
	if !ok {
		return -1.0
	}

	// 终点不存在，说明 target 是未知变量
	_, ok = graph[target]
	if !ok {
		return -1.0
	}

	// 记录已经访问过的节点，防止在图中来回循环
	visit := map[string]bool{}

	// BFS 队列中的状态
	// node：当前走到了哪个变量
	// val：start / node 的值
	type State struct {
		node string
		val  float64
	}

	queue := []State{}

	// 从 start 开始
	// start / start = 1
	queue = append(queue, State{
		node: start,
		val:  1,
	})
	visit[start] = true

	for len(queue) != 0 {
		// 取出队首
		cur := queue[0]
		queue = queue[1:]

		// 到达目标节点
		// cur.val 就是 start / target
		if cur.node == target {
			return cur.val
		}

		// 遍历当前节点能到达的所有邻居
		for i := 0; i < len(graph[cur.node]); i++ {
			next := graph[cur.node][i]

			// 已经访问过，不再重复访问
			if visit[next.to] {
				continue
			}

			visit[next.to] = true

			// 假设：
			// start / cur.node = cur.val
			// cur.node / next.to = next.value
			//
			// 那么：
			// start / next.to
			// = (start / cur.node) * (cur.node / next.to)
			queue = append(queue, State{
				node: next.to,
				val:  cur.val * next.value,
			})
		}
	}

	// BFS 结束仍然没有找到 target
	// 说明两个变量虽然都存在，但不在同一个连通块
	return -1.0
}

func bfs2(graph map[string][]edge, start, target string) float64 {
	// 只要有一个变量不存在，就无法计算
	if _, ok := graph[start]; !ok {
		return -1
	}
	if _, ok := graph[target]; !ok {
		return -1
	}

	type State struct {
		node string
		val  float64 // start / node
	}

	queue := []State{
		{node: start, val: 1},
	}

	visited := map[string]bool{
		start: true,
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node == target {
			return cur.val
		}

		for _, next := range graph[cur.node] {
			if visited[next.to] {
				continue
			}

			visited[next.to] = true

			queue = append(queue, State{
				node: next.to,
				val:  cur.val * next.value,
			})
		}
	}

	return -1
}
