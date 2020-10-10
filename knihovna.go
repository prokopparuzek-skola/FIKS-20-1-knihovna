package main

import (
	"fmt"
)

type room struct {
	C  int
	T  int
	V  int
	ID []int
	G  float64 // poměr svitků s časem
}

type element struct {
	left  *element
	right *element
	value room
}

func (e *element) append(v room) *element {
	e.right = &element{e, nil, v}
	return e.right
}

func (e *element) delete() {
	if e.right != nil {
		e.right.left = e.left
	}
	if e.left != nil {
		e.left.right = e.right
	}
}

func (l *element) String() (out string) {
	if l == nil {
		return
	}
	out += fmt.Sprint(l.value)
	out += ","
	out += l.right.String()
	return
}

func solve(rooms *element) (save int, how []int) {
	var previous map[int]bool
	var this map[int]bool
	var uniq []int
	var tik int
	var best *element = rooms
	var pBest *element
	var r *element

	previous = make(map[int]bool)

	for tik <= rooms.value.T {
		for r = rooms; r != nil; r = r.right { // pročištění
			uniq = make([]int, 0)
			this = make(map[int]bool)
			for _, paper := range r.value.ID { // projde zatím nesmazané svitky
				if previous[paper] == false && this[paper] == false {
					uniq = append(uniq, paper)
					this[paper] = true
				}
			}
			r.value.ID = uniq
			r.value.G = float64(len(r.value.ID)) / float64(r.value.V)
			if r.value.G > best.value.G && (pBest == nil || r.value.G < pBest.value.G) {
				best = r
			}
		}
		fmt.Println(rooms)
		// upravení výsledků
		if pBest == best {
			return
		}
		if (tik + best.value.V) <= best.value.T {
			fmt.Println(*best)
			for _, paper := range best.value.ID {
				previous[paper] = true
			}
			pBest = nil
			save += len(best.value.ID)
			how = append(how, best.value.C)
			tik += best.value.V
			best.delete()
			best = rooms
		} else {
			pBest = best
		}
	}
	return
}

func main() {
	var U int

	fmt.Scan(&U)
	for i := 0; i < U; i++ {
		var N, K int
		var C, T, V, M int
		var rooms, e *element
		var ID []int
		rooms = &element{}
		e = rooms

		fmt.Scan(&N, &K)
		fmt.Scan(&C, &T, &V, &M)
		e.value = room{C, T, V, ID, float64(len(ID)) / float64(V)}
		ID = make([]int, M)
		for k := 0; k < M; k++ {
			fmt.Scan(&ID[k])
		}
		for j := 1; j < N; j++ {
			fmt.Scan(&C, &T, &V, &M)
			ID = make([]int, M)
			for k := 0; k < M; k++ {
				fmt.Scan(&ID[k])
			}
			e = e.append(room{C, T, V, ID, 0.0})
		}
		fmt.Println(rooms)
		fmt.Println(solve(rooms))
	}
}
