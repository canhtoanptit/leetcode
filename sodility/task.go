package sodility

import "strings"

// Solution write a function that returns true if the two strings are anagrams of each other, otherwise return false
func Solution(x string, y string) bool {
	j := 0 // Index for string y

	// Iterate over string x
	for i := 0; i < len(x) && j < len(y); i++ {
		// If the character in x matches the current character in y, move to the next character in y
		if x[i] == y[j] {
			j++
		}
	}

	// If we've reached the end of y, all characters were found in x in order
	return j == len(y)
}

func Solution2(connections []string, name1 string, name2 string) int {
	graph := make(map[string][]string)

	// Build the graph from the connections slice
	for _, connection := range connections {
		parts := strings.Split(connection, "-")
		if len(parts) == 2 {
			graph[parts[0]] = append(graph[parts[0]], parts[1])
			graph[parts[1]] = append(graph[parts[1]], parts[0])
		}
	}

	// Perform BFS to find the shortest path from name1 to name2
	type queueItem struct {
		name     string
		distance int
	}

	queue := []queueItem{{name: name1, distance: 0}}
	visited := make(map[string]bool)
	visited[name1] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.name == name2 {
			return item.distance
		}

		for _, neighbor := range graph[item.name] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, queueItem{name: neighbor, distance: item.distance + 1})
			}
		}
	}

	return -1 // Return -1 if there is no path between name1 and name2
}
