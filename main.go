package wordcount

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputWords := os.Args[1]

	if len(inputWords) <= 0 {
		fmt.Println(0)
		return
	}

	wordSlice := strings.Split(inputWords, " ")

	fmt.Println(len(wordSlice))
	return
}