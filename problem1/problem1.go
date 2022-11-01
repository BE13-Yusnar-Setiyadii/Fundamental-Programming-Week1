package main

import "fmt"

func DrawXYZ(n int) string {
	var pangkat int = 2
	var bantu int = 1
	var result string

	for i := 0; i < pangkat; i++ {
		bantu = bantu * n
	}
	for i := 1; i <= bantu; i++ {
		if i%3 == 0 {
			result = result + "X "
		} else if i%2 == 0 {
			result = result + "Z "
		} else {
			result = result + "Y "
		}
		if i%n == 0 {
			result = result + "\n"
		}
	}
	return result
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
