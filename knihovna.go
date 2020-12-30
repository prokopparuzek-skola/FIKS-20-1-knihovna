package main

import (
	"fmt"
	"sort"
)

type room struct {
	C  int
	T  int
	V  int
	M  int
	ID []int
}

type thing struct {
	weight  int
	parentX int
	parentY int
}

func max(arr []int) (max int) {
	for _, i := range arr {
		if max < i {
			max = i
		}
	}
	return
}

func solve(rooms []room, K int) (count int, path []int) {
	var taken [][][]bool
	var bag [][]thing
	var A, F []int

	path = make([]int, 0)
	taken = make([][][]bool, len(rooms))
	bag = make([][]thing, len(rooms))
	A = make([]int, 0)
	F = make([]int, 0)
	sort.Slice(rooms, func(i, j int) bool { return rooms[i].T < rooms[j].T })
	for i := range taken {
		taken[i] = make([][]bool, K)
		for j := range taken[i] {
			taken[i][j] = make([]bool, K)
		}
	}
	for i := range bag {
		bag[i] = make([]thing, K)
	}
	A = append(A, 0)
	if rooms[0].V <= rooms[0].T {
		var svitky map[int]bool = make(map[int]bool)
		var c int
		for _, s := range rooms[0].ID {
			if !svitky[s] {
				c++
				svitky[s] = true
			}
		}
		A = append(A, c)
		for _, s := range rooms[0].ID {
			taken[0][rooms[0].M][s] = true
		}
		bag[0][rooms[0].M].weight = rooms[0].V
	}
	for i := 0; i < len(rooms)-1; i++ {
		for _, r := range A {
			// nepřidám
			if bag[i+1][r].weight == 0 || bag[i+1][r].weight > bag[i][r].weight {
				if bag[i+1][r].weight == 0 {
					F = append(F, r)
				}
				bag[i+1][r].weight = bag[i][r].weight
				bag[i+1][r].parentY = bag[i][r].parentY
				bag[i+1][r].parentX = bag[i][r].parentX
				for j, s := range taken[i][r] {
					taken[i+1][r][j] = s
				}
			}
			// přidám
			if (bag[i][r].weight + rooms[i+1].V) <= rooms[i+1].T { // stihnu to?
				var newS map[int]bool = make(map[int]bool) // jaké svitky ještě nebyly

				for _, svitek := range rooms[i+1].ID {
					if !taken[i][r][svitek] && !newS[svitek] {
						newS[svitek] = true
					}
				}
				if len(newS) != 0 && (bag[i+1][r+len(newS)].weight == 0 || bag[i][r].weight+rooms[i+1].V < bag[i+1][r+len(newS)].weight) {
					if bag[i+1][r+len(newS)].weight == 0 {
						F = append(F, r+len(newS))
					}
					bag[i+1][r+len(newS)].weight = bag[i][r].weight + rooms[i+1].V
					bag[i+1][r+len(newS)].parentX = r
					bag[i+1][r+len(newS)].parentY = i
					for j, s := range taken[i][r] {
						taken[i+1][r+len(newS)][j] = s
					}
					for s, _ := range newS {
						taken[i+1][r+len(newS)][s] = true
					}
				}
			}
		}
		A = F
		sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
		F = make([]int, 0)
	}
	var last int
	count = max(A)
	if bag[len(bag)-1][count].parentY != bag[len(bag)-2][count].parentY {
		last = len(bag) - 1
	} else {
		last = bag[len(bag)-1][count].parentY
	}
	for x, y := count, last; x != 0; { // najdi cestu
		path = append(path, rooms[y].C)
		x = bag[y][x].parentX
		y = bag[y][x].parentY
	}
	for i := 0; i < len(path)/2; i++ { // reverse
		swp := path[i]
		path[i] = path[len(path)-i-1]
		path[len(path)-i-1] = swp
	}
	return
}

func main() {
	var U int
	fmt.Scan(&U)
	for i := 0; i < U; i++ {
		var N, K int
		var rooms []room
		fmt.Scan(&N, &K)
		rooms = make([]room, N)
		for j, r := range rooms {
			fmt.Scan(&r.C, &r.T, &r.V, &r.M)
			r.ID = make([]int, r.M)
			for k, _ := range r.ID {
				fmt.Scan(&r.ID[k])
			}
			rooms[j] = r
		}
		c, a := solve(rooms, K)
		fmt.Println(c)
		for _, r := range a[:len(a)-1] {
			fmt.Printf("%d ", r)
		}
		fmt.Println(a[len(a)-1])
	}
}
