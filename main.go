package wordcount

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var inputWords []string
	for scanner.Scan() {
		inputWords = append(inputWords, scanner.Text())
	}

	if len(inputWords) <= 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(len(inputWords))
	return
}