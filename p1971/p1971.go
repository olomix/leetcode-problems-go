package p1971

func validPath(n int, edges [][]int, source int, destination int) bool {
	queue := []int{source}
	seen := map[int]bool{source: true}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == destination {
			return true
		}
		for _, n := range neighbors(edges, node) {
			if !seen[n] {
				seen[n] = true
				queue = append(queue, n)
			}
		}
	}

	return false
}

func neighbors(edges [][]int, node int) []int {
	var ns []int
	for i := range edges {
		if edges[i][0] == node {
			ns = append(ns, edges[i][1])
		}
		if edges[i][1] == node {
			ns = append(ns, edges[i][0])
		}
	}
	return ns
}
