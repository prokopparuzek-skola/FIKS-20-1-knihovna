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

func solve(rooms []room, K int) (count int, path []int) {
	var taken [][]bool
	var bag [][]thing
	var A, F []int

	path = make([]int, 0)
	taken = make([][]bool, K)
	bag = make([][]thing, len(rooms))
	A = make([]int, 0)
	F = make([]int, 0)
	sort.Slice(rooms, func(i, j int) bool { return i < j })
	for i := range taken {
		taken[i] = make([]bool, K)
	}
	for i := range bag {
		bag[i] = make([]thing, K)
	}
	A = append(A, 0)
	for i := 0; i < len(rooms)-1; i++ {
		for _, r := range A {
			// nepřidám
			if bag[i+1][r].weight == 0 || bag[i+1][r].weight > bag[i][r].weight {
				if bag[i+1][r].weight == 0 {
					F = append(F, r)
				} else { // unpush
				}
				bag[i+1][r].weight = bag[i][r].weight
				bag[i+1][r].parentY = bag[i][r].parentY
				bag[i+1][r].parentX = bag[i][r].parentX
			}
			// přidám
			if (bag[i][r].weight + rooms[i+1].V) <= rooms[i+1].T { // stihnu to?
				var newS []int = make([]int, 0) // jaké svitky ještě nebyly

				for _, svitek := range rooms[i+1].ID {
					if !taken[r][svitek] {
						newS = append(newS, svitek)
					}
				}
				if len(newS) != 0 && bag[i+1][r+len(newS)].weight == 0 || bag[i][r].weight+rooms[i+1].V < bag[i+1][r+len(newS)].weight {
					if bag[i+1][r+len(newS)].weight == 0 {
						F = append(F, r)
					} else { // unpush
					}
					bag[i+1][r+len(newS)].weight = bag[i][r].weight + rooms[i+1].V
					bag[i+1][r+len(newS)].parentX = r
					bag[i+1][r+len(newS)].parentY = i
					for j, s := range taken[r] {
						taken[r+len(newS)][j] = s
					}
					for _, s := range newS {
						taken[r+len(newS)][s] = true
					}
				}
			}
		}
		A = F
		F = make([]int, 0)
	}
	count = A[len(A)-1]
	for x, y := count, len(bag)-1; x != 0; { // najdi cestu
		path = append(path, y)
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
	}
}
