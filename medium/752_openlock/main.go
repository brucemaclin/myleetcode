package main

import "fmt"

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	}
	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, "0000")
	if dead[target] {
		return -1
	}
	var level int
	for {
		if len(queue) == 0 {
			break
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			checkStr := queue[i]
			if checkStr == target {
				return level
			}
			for i := 0; i < 4; i++ {
				slice1 := []rune(checkStr)
				slice2 := []rune(checkStr)
				if slice1[i] == rune('0') {
					slice1[i] = '9'
				} else {
					slice1[i]--
				}
				if !visited[string(slice1)] && !dead[string(slice1)] {
					visited[string(slice1)] = true
					queue = append(queue, string(slice1))
				}
				if slice2[i] == rune('9') {
					slice2[i] = '0'
				} else {
					slice2[i]++
				}
				if !visited[string(slice2)] && !dead[string(slice2)] {
					visited[string(slice2)] = true
					queue = append(queue, string(slice2))
				}

			}
		}
		level++
		queue = queue[size:]
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
