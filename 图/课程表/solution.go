package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	indegrees := make([]int, numCourses)
	for _, course := range prerequisites {
		graph[course[1]] = append(graph[course[1]], course[0])
		indegrees[course[0]]++
	}
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		for _, next := range graph[course] {
			indegrees[next]--
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		if indegrees[i] != 0 {
			return false
		}
	}
	return true
}
