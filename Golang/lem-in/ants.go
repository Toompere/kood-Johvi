package main

import (
	"fmt"
	"math"
)

func MoveAnts() {
	done := false
	ant := 1

	for !done {
		done = true
		fmt.Println()
		for i, v := range paths {
			empty := CheckRooms(v)
			if antCount == 0 && !empty {
				done = false
			}
			if antCount > 0 {
				if MoveNextAnt(ant, i) {
					ant++
					antCount--
					done = false
				} else if i == len(paths)-1 && empty {
					tmppaths := paths[:len(paths)-1]
					paths = tmppaths
				}
			}

		}
	}
}

func CheckRooms(path []string) bool {
	empty := true
	pathlen := len(path) - 2
	for i := pathlen; i >= 0; i-- {
		antNum := rooms[path[i]].Ant
		if antNum != 0 {
			tmproom := rooms[path[i+1]]
			tmproom.Ant = antNum
			rooms[path[i+1]] = tmproom
			fmt.Printf("L%v-%v ", tmproom.Ant, tmproom.Name)
			tmproom = rooms[path[i]]
			tmproom.Ant = 0
			rooms[path[i]] = tmproom
			if i+1 <= pathlen+1 {
				empty = false
			}
		}
	}
	return empty
}

func MoveNextAnt(ant, index int) bool {
	if index > 0 {
		sum := antCount
		for i := 0; i <= index; i++ {
			sum += len(paths[i])
		}
		sum = (sum - (index+1)*2)
		res := math.Round(float64(sum) / float64(index+1))
		if int(res) < len(paths[index])-1 {
			return false
		}
	}

	tmproom := rooms[paths[index][1]]
	tmproom.Ant = ant
	rooms[paths[index][1]] = tmproom
	fmt.Printf("L%v-%v ", tmproom.Ant, tmproom.Name)
	return true

}
