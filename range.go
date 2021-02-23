// package main

// import "fmt"

// func main() {
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}

// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }

package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) map[string]int {
	s := strings.Split(str, " ")
	count := make(map[string]int)
	for i := range s {
		count[s[i]]++
	}

	return count
}

func main() {
	fmt.Println(wordCount("wow nice"))
}
