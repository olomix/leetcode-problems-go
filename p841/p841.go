package p841

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]uint64, (len(rooms)+63)/64)
	isSeen := func(i int) bool {
		return seen[i/64]&(1<<uint(i%64)) != 0
	}
	setSeen := func(i int) {
		seen[i/64] |= 1 << uint(i%64)
	}
	queue := make([]int, 1, len(rooms))
	queue[0] = 0
	setSeen(0)
	cnt := 1
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for _, j := range rooms[i] {
			if !isSeen(j) {
				setSeen(j)
				cnt++
				queue = append(queue, j)
			}
		}
	}
	return cnt == len(rooms)
}
