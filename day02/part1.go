package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var location = map[string]int {
		"x": 0,
		"y": 0,
	}

	f, _ := os.Open("./day02/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		direction, distanceString := parts[0], parts[1]
		distance, _ := strconv.Atoi(distanceString)
		if direction == "forward"{
			location["x"] += distance
		} else if direction == "down" {
			location["y"] += distance
		} else if direction == "up" {
			location["y"] -= distance 
		}
	}

	fmt.Println("Location: x:" + strconv.Itoa(location["x"]) + ", y:" + strconv.Itoa(location["y"]))
	fmt.Println("Answer: " + strconv.Itoa(location["x"] * location["y"]))
}