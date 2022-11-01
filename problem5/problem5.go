package main

import "fmt"

func RemoveDuplicates(array []int) int {
	keys := make(map[int]bool)
	list := []int{}

	for _, items := range array {
		if _, value := keys[items]; !value {
			keys[items] = true
			list = append(list, items)

		}
	}
	return len(list)
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) //4
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) //6
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))     //4
}
