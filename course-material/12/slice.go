package main

import (
	"log"
	"sort"
)

func main() {

	numbers := []int{1, 12, 3, 4, 5, 6, 7, 8, 9, 10}

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	log.Println(numbers)

}
