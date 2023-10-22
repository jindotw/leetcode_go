package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	// Initialize the adjacency list and indegree array.
	adjList := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pres := range prerequisites {
		course := pres[0]
		prerequisite := pres[1]
		adjList[prerequisite] = append(adjList[prerequisite], course)
		indegree[course]++
	}

	// Identify courses without prerequisites and add them to the queue.
	queue := make([]int, 0)
	for i, count := range indegree {
		if count == 0 {
			queue = append(queue, i)
		}
	}

	// Start the topological sort using BFS.
	var ordering []int
	visitedCourses := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		ordering = append(ordering, current)
		visitedCourses++

		for _, nextCourse := range adjList[current] {
			indegree[nextCourse]--
			if indegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// If all courses are visited, return the order, otherwise return an empty array.
	if visitedCourses == numCourses {
		return ordering
	}
	return []int{}
}

func main() {
	numCourses := 0
	// [[1,0],[2,0],[3,1],[3,2]]
	// 0:1,2
	// 1:3
	// 2:3
	// 0->1-2->3
	// prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	prerequisites := [][]int{}
	fmt.Println(findOrder(numCourses, prerequisites))
}
