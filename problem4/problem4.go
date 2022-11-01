package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	var hasil []int
	for baris := 0; baris < len(arrayA); baris++ {
		var jawab bool = false
		for kolom := 0; kolom < len(arrayB); kolom++ {
			if arrayA[baris] == arrayB[kolom] {
				jawab = true
			}
		}
		if !jawab {
			hasil = append(hasil, arrayA[baris])
		}
	}
	if len(hasil) == 0 {
		return []int{}
	}
	return hasil
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))   //[2,4]
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})) //[20, 30, 40]
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))              //[7]
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))                    //[3]
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))              //[]
}
