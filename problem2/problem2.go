package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {
	var result2 float64
	var i float64
	var result1 float64
	var min float64

	for i = 0; i < float64(len(arrayInput)); i++ {
		result1 += arrayInput[int(i)]
	}
	for i = 1; i <= float64(len(arrayInput)); i++ {
		result2 = i
	}
	min = result1 / result2
	var median float64
	l := len(arrayInput)
	if l%2 == 0 {
		median = (arrayInput[l/2-1] + arrayInput[l/2]) / 2
	} else {
		median = arrayInput[l/2]
	}
	return min, median

}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
