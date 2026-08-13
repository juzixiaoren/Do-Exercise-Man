package main

func minMutation(startGene string, endGene string, bank []string) int {
	type Edge struct {
		Gene string
		Step int
	}
	graph := map[string][]string{}
	queue := []Edge{}
	visit := map[string]bool{}
	for i := 0; i < len(bank); i++ {
		buildgraph(graph, startGene, bank[i], true)
		for j := 1; j < len(bank); j++ {
			buildgraph(graph, bank[i], bank[j], false)
		}
	}
	queue = append(queue, Edge{
		Gene: startGene,
		Step: 0,
	})
	visit[startGene] = true
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range graph[cur.Gene] {
			if visit[next] {
				continue
			} else {
				if next == endGene {
					return cur.Step + 1
				} else {
					visit[next] = true
					queue = append(queue, Edge{
						Gene: next,
						Step: cur.Step + 1,
					})
				}
			}
		}
	}
	return -1
}
func buildgraph(graph map[string][]string, startGene, endGene string, hasStart bool) {
	if len([]byte(startGene)) != len([]byte(endGene)) {
		return
	}
	count := 0
	for i := 0; i < len([]byte(startGene)); i++ {
		if startGene[i] == endGene[i] {
			continue
		} else {
			count++
			if count == 2 {
				return
			}
		}
	}
	if count == 1 {
		if hasStart {
			graph[startGene] = append(graph[startGene], endGene)
		} else {
			graph[startGene] = append(graph[startGene], endGene)
			graph[endGene] = append(graph[endGene], startGene)
		}
	}
}

/*
题目：最小基因变化
题解：使用BFS，从startGene开始，每次变化一个基因，直到到达endGene，返回变化的次数。其中 bank 中的是有效基因
思路：
1. 构建图，每个基因节点连接到其他基因节点，如果两个基因只有一个字符不同，则它们之间有一条边
2. 使用BFS，从startGene开始，每次变化一个基因，直到到达endGene，返回变化的次数。其中 bank 中的是有效基因
实际，建图中的 cost 很大，因为需要两两比对，复杂度是 O(n^2*L)
实际上，每一个基因的变化是有限的，只有 L*4 次变化，所以只需要动态构建可能的路径，并在 bank 中查
这样复杂度就降为O(n*L*4)了
*/
func minMutation_fix(startGene string, endGene string, bank []string) int {
	type Edge struct {
		Gene string
		Step int
	}

	if startGene == endGene {
		return 0
	}

	bankSet := map[string]struct{}{}
	visit := map[string]struct{}{}
	queue := []Edge{}
	geneList := []byte{'A', 'C', 'G', 'T'}

	for _, gene := range bank {
		bankSet[gene] = struct{}{}
	}

	_, ok := bankSet[endGene]
	if !ok {
		return -1
	}

	queue = append(queue, Edge{
		Gene: startGene,
		Step: 0,
	})
	visit[startGene] = struct{}{}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		curGene := []byte(cur.Gene)

		for i := 0; i < len(curGene); i++ {
			oldGene := curGene[i]

			for _, gene := range geneList {
				if gene == oldGene {
					continue
				}

				curGene[i] = gene
				nextGene := string(curGene)

				if _, ok := bankSet[nextGene]; !ok {
					continue
				}

				if _, ok := visit[nextGene]; ok {
					continue
				}

				if nextGene == endGene {
					return cur.Step + 1
				}

				visit[nextGene] = struct{}{}
				queue = append(queue, Edge{
					Gene: nextGene,
					Step: cur.Step + 1,
				})
			}

			curGene[i] = oldGene
		}
	}

	return -1
}
