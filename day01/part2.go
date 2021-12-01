package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func containsZero(s []int) bool {
	for _, v := range s {
		if v == 0 {
			return true
		}
	}

	return false
}


func main(){
	var largerCount = 0
	var prev3Values = make([]int, 3,3)

	f, _ := os.Open("./day01/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		value, _ := strconv.Atoi(scanner.Text())

		if !containsZero(prev3Values){
			currentWindowTotal := value + prev3Values[2] + prev3Values[1]
			prevWindowTotal := prev3Values[2] + prev3Values[1] + prev3Values[0]

			if currentWindowTotal > prevWindowTotal{
				largerCount++
			}
		}

		// Setup for next round
		prev3Values = append(prev3Values[1:], value)
	}
	if scanner.Err() == bufio.ErrTooLong {
		log.Fatal(scanner.Err())
	}

	fmt.Println("Larger Count: " + strconv.Itoa(largerCount))
}