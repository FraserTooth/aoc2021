package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	var largerCount = 0
	var prevLine = 0

	f, _ := os.Open("./day01/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		if prevLine != 0 && line > prevLine {
			largerCount++
		}
		prevLine = line
	}

	fmt.Println("Larger Count: " + strconv.Itoa(largerCount))
}