package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 || len(prerequisites[0]) == 0 {
		return true
	}
	var result []int
	visited := make(map[int]bool)
	tmp := make(map[int]bool)
	graph := make(map[int][]int)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	for k := range graph {
		if dfs(k, visited, graph, tmp, &result) {
			return false
		}
	}
	return len(result) == numCourses

}
func dfs(i int, visited map[int]bool, graph map[int][]int, tmp map[int]bool, result *[]int) bool {
	if tmp[i] {
		return true
	}
	if visited[i] {
		return false
	}
	tmp[i] = true
	for _, v := range graph[i] {
		if dfs(v, visited, graph, tmp, result) {
			return true
		}
	}
	tmp[i] = false
	visited[i] = true
	*result = append(*result, i)
	return false
}
func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
