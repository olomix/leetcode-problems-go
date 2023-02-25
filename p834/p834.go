package p834

import (
	"fmt"
	"math"
)

func sumOfDistancesInTree3(n int, edges [][]int) []int {
	neighbors := make(map[int][]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}

	var result = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 1 {
			fmt.Println("here")
		}
		fmt.Println(i)
		sum := 0
		seen := map[int]bool{i: true}
		queue := []int{i}
		weights := []int{1}
		for len(queue) > 0 {
			el := queue[0]
			weight := weights[0]
			queue = queue[1:]
			weights = weights[1:]
			for _, neighbor := range neighbors[el] {
				if !seen[neighbor] {
					seen[neighbor] = true
					queue = append(queue, neighbor)
					weights = append(queue, weight+1)
					fmt.Printf("neightbor: %d, weight: %d\n", neighbor, weight)
					sum += weight
				}
			}
		}
		result[i] = sum
	}
	return result
}

func sumOfDistancesInTree2(n int, edges [][]int) []int {
	var dest = make([][]int, n)
	for i := range dest {
		dest[i] = make([]int, n)
		for j := range dest[i] {
			dest[i][j] = math.MaxInt32
		}
	}

	for _, edge := range edges {
		dest[edge[0]][edge[1]] = 1
		dest[edge[1]][edge[0]] = 1
	}
	for i := 0; i < n; i++ {
		dest[i][i] = 0
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dest[i][j] > dest[i][k]+dest[k][j] {
					dest[i][j] = dest[i][k] + dest[k][j]
				}
			}
		}
	}

	var result []int
	for i := 0; i < n; i++ {
		var sum int
		for j := 0; j < n; j++ {
			sum += dest[i][j]
		}
		result = append(result, sum)
	}

	return result
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var dest = make([]uint16, n*n)
	for i := range dest {
		dest[i] = math.MaxUint16 / 6
	}

	for _, edge := range edges {
		dest[edge[0]*n+edge[1]] = 1
		dest[edge[1]*n+edge[0]] = 1
	}
	for i := 0; i < n; i++ {
		dest[i*n+i] = 0
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dest[i*n+j] > dest[i*n+k]+dest[k*n+j] {
					dest[i*n+j] = dest[i*n+k] + dest[k*n+j]
				}
			}
		}
	}

	var result []int
	for i := 0; i < n; i++ {
		var sum int
		for j := 0; j < n; j++ {
			sum += int(dest[i*n+j])
		}
		result = append(result, sum)
	}

	return result
}
