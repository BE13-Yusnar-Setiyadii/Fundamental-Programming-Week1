package main

import (
	"fmt"
	"strings"
)

func MunculSekali(angka string) []int {
	var jawaban []int
	for i := 0; i < len(angka); i++ {
		if strings.Count(angka, string(angka[i])) <= 1 {
			jawaban = append(jawaban, int(angka[i]-'0'))
		}
	}
	if len(jawaban) == 0 {
		return []int{}
	}
	return jawaban
}

func main() {
	fmt.Println(MunculSekali("1234123"))    // [4]
	fmt.Println(MunculSekali("76523752"))   // [6, 3]
	fmt.Println(MunculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // []
	fmt.Println(MunculSekali("0872504"))    // [8 7 2 5 4]
}
