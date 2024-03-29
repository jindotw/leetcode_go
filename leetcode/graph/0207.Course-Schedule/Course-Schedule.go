package main

import "fmt"

func canFinish1(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for _, pre := range prerequisites {
		course := pre[0]
		prerequisite := pre[1]
		preMap[course] = append(preMap[course], prerequisite)
	}

	visited := make(map[int]struct{})
	var dfs func(int) bool
	dfs = func(course int) bool {
		if _, present := visited[course]; present {
			return false
		}
		if len(preMap[course]) == 0 {
			return true
		}
		visited[course] = struct{}{}
		for _, pre := range preMap[course] {
			if !dfs(pre) {
				return false
			}
			delete(preMap, pre)
		}
		delete(visited, course)
		return true
	}

	for course := 0; course < numCourses; course++ {
		if !dfs(course) {
			return false
		}
	}

	return true
}

/*
Topological Sort Algorithm Notes:

1. Create an adjacency list to represent the graph.
2. Calculate the dep (number of incoming edges) for each node (course).

3. Initialize a queue and add all nodes with dep 0 (no prerequisites).

 4. Use BFS approach:
    a. Dequeue a node (course).
    b. For each neighbor of this node:
    i.  Reduce its dep by 1.
    ii. If its dep becomes 0, add it to the queue.
    c. Repeat until the queue is empty.

 5. Check if all nodes (courses) have been visited:
    a. If visited count equals total nodes, a topological order exists.
    b. If not, there's a cycle and topological order isn't possible.

The goal is to check if all courses can be taken by ensuring there's no cycle (deadlock in terms of prerequisites)
in the graph.
*/
func canFinish2(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	dep := make([]int, numCourses)

	for _, pre := range prerequisites {
		course := pre[1]
		unlock := pre[0]
		adj[course] = append(adj[course], unlock)
		dep[unlock]++
	}
	queue := make([]int, 0)
	for i, val := range dep {
		if val == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		visited++

		for _, dependsOn := range adj[course] {
			dep[dependsOn]--
			if dep[dependsOn] == 0 {
				queue = append(queue, dependsOn)
			}
		}
	}

	return visited == numCourses

}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 1}, {2, 3}}
	//numCourses := 2
	//prerequisites := [][]int{{1, 0}, {0, 1}}
	//  fmt.Println(canFinish(numCourses, prerequisites))
	fmt.Println(canFinish1(numCourses, prerequisites))
	fmt.Println(canFinish2(numCourses, prerequisites))
}

// 0 -> 1
// 1 -> 2
// 3 -> 2
