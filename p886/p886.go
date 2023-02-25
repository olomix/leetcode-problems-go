package p886

func possibleBipartition(n int, dislikes [][]int) bool {
	seen := map[int]bool{}
	for len(seen) < n {
		queue := []int{}
		red := map[int]bool{}
		for i := 1; i <= n; i++ {
			if !seen[i] {
				queue = append(queue, i)
				seen[i] = true
				red[i] = true
				break
			}
		}
		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			neighbors := []int{}
			for _, dislike := range dislikes {
				if dislike[0] == item {
					neighbors = append(neighbors, dislike[1])
				}
				if dislike[1] == item {
					neighbors = append(neighbors, dislike[0])
				}
			}

			colorFound := false
			isRed := false
			for _, neighbor := range neighbors {
				if seen[neighbor] {
					if !colorFound {
						colorFound = true
						isRed = red[neighbor]
					} else {
						if isRed != red[neighbor] {
							return false
						}
					}
				} else {
					seen[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
			red[item] = !isRed
		}
	}
	return true
}
