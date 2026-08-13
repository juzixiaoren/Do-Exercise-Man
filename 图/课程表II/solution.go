package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	queue := []int{}
	count := 0
	ans := []int{}
	for _, coursePair := range prerequisites {
		graph[coursePair[1]] = append(graph[coursePair[1]], coursePair[0])
		indegrees[coursePair[0]]++
	}
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		ans = append(ans, course)
		for _, next := range graph[course] {
			indegrees[next]--
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if count != numCourses {
		return []int{}
	}
	return ans
}
