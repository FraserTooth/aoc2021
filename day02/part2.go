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
		"aim": 0,
	}

	f, _ := os.Open("./day02/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		direction, quantityString := parts[0], parts[1]
		quantity, _ := strconv.Atoi(quantityString)
		if direction == "forward"{
			location["x"] += quantity
			location["y"] += quantity * location["aim"]
		} else if direction == "down" {
			location["aim"] += quantity
		} else if direction == "up" {
			location["aim"] -= quantity 
		}
	}

	fmt.Println("Location: x:" + strconv.Itoa(location["x"]) + ", y:" + strconv.Itoa(location["y"]))
	fmt.Println("Answer: " + strconv.Itoa(location["x"] * location["y"]))
}