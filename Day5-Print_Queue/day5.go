package main

import (
	"fmt"
	"os"
)

func main() {
	input := ReadInput("./input-day5-print_queue.txt")
	fmt.Printf("%#v\n", input)
}

func ReadInput(path string) string {
	content, _ := os.ReadFile(path)
	return string(content)
}
