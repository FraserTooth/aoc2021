package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
        file, err := os.Open("input.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        instructionsChannels := [2]chan string{}
        for i := range instructionsChannels {
                instructionsChannels[i] = make(chan string)
        }
        solutionChannels := [2]chan string{}
        for i := range instructionsChannels {
                solutionChannels[i] = make(chan string)
        }
        go partOne(instructionsChannels[0], solutionChannels[0])
        go partTwo(instructionsChannels[1], solutionChannels[1])

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                instruction := scanner.Text()
                for _, instructions := range instructionsChannels {
                        instructions <- instruction
                }
        }
        for _, instructions := range instructionsChannels {
                close(instructions)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        for _, solutions := range solutionChannels {
                for solution := range solutions {
                        fmt.Printf("%s\n", solution)
                }
        }
}



func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var instructions []string
        var coords [][]int
        maxX := 0
        maxY := 0
	for line := range instructionChannel {
                if len(line)==0{
                        continue
                }
                firstChar := line[0:1]
                // Instruction
                if firstChar == "f"{
                        instructions = append(instructions, strings.Split(line," ")[2])
                // Dot
                } else {
                        coordString := strings.Split(line,",")
                        y, _ := strconv.Atoi(coordString[1])
                        x, _ := strconv.Atoi(coordString[0])
                        if x > maxX{
                                maxX = x
                        }
                        if y > maxY {
                                maxY = y
                        }
                        coord := []int{x,y}
                        coords = append(coords, coord)
                }
                
	}
        // Make a sized grid of zeros
        paper := make([][]int, maxY+1)
        for i := range paper{
                paper[i] = make([]int, maxX+1)
        }
        // Fill the dots
        for _, coord := range coords{
                paper[coord[1]][coord[0]] = 1
        }

        // Do one fold
        instructionParts := strings.Split(instructions[0], "=")
        axis := instructionParts[0]
        foldPoint, _ :=  strconv.Atoi(instructionParts[1])
        if axis == "y"{
                for yAfterFold, row := range paper[foldPoint + 1:] {
                        actualY := foldPoint + 1 + yAfterFold
                        equivalentY := foldPoint - (yAfterFold + 1 * 2) + 1
                        for x := range row{
                                paper[equivalentY][x] += paper[actualY][x]
                        }
                }
                paper = paper[:foldPoint]
        } else {
                for y, row := range paper {
                        for xAfterFold := range row[foldPoint + 1:]{
                                actualX := foldPoint + 1 + xAfterFold
                                equivalentX := foldPoint - (xAfterFold + 1 * 2) + 1
                                paper[y][equivalentX] += paper[y][actualX]
                        }
                        paper[y] = paper[y][:foldPoint]
                }
        }

        // Count Dots
        dotsLeft := 0
        for _, row := range paper{
                for _, val := range row{
                        if val > 0{
                                dotsLeft++
                        }
                }
        }

	solutions <- fmt.Sprintf(
                "Part 1:\n Dots Left: %d\n",
                dotsLeft,
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var instructions []string
        var coords [][]int
        maxX := 0
        maxY := 0
	for line := range instructionChannel {
                if len(line)==0{
                        continue
                }
                firstChar := line[0:1]
                // Instruction
                if firstChar == "f"{
                        instructions = append(instructions, strings.Split(line," ")[2])
                // Dot
                } else {
                        coordString := strings.Split(line,",")
                        y, _ := strconv.Atoi(coordString[1])
                        x, _ := strconv.Atoi(coordString[0])
                        if x > maxX{
                                maxX = x
                        }
                        if y > maxY {
                                maxY = y
                        }
                        coord := []int{x,y}
                        coords = append(coords, coord)
                }
                
	}
        // Make a sized grid of zeros
        paper := make([][]int, maxY+1)
        for i := range paper{
                paper[i] = make([]int, maxX+1)
        }
        // Fill the dots
        for _, coord := range coords{
                paper[coord[1]][coord[0]] = 1
        }

        // Do one fold
        for _, instruction :=  range instructions {
                instructionParts := strings.Split(instruction, "=")
                axis := instructionParts[0]
                foldPoint, _ :=  strconv.Atoi(instructionParts[1])
                if axis == "y"{
                        for yAfterFold, row := range paper[foldPoint + 1:] {
                                actualY := foldPoint + 1 + yAfterFold
                                equivalentY := foldPoint - (yAfterFold + 1 * 2) + 1
                                for x := range row{
                                        paper[equivalentY][x] += paper[actualY][x]
                                }
                        }
                        paper = paper[:foldPoint]
                } else {
                        for y, row := range paper {
                                for xAfterFold := range row[foldPoint + 1:]{
                                        actualX := foldPoint + 1 + xAfterFold
                                        equivalentX := foldPoint - (xAfterFold + 1 * 2) + 1
                                        paper[y][equivalentX] += paper[y][actualX]
                                }
                                paper[y] = paper[y][:foldPoint]
                        }
                }
        }

        // Count Dots and Print Paper Layout
        dotsLeft := 0
        for _, row := range paper{
                for _, val := range row{
                        if val > 0{
                                dotsLeft++
                                fmt.Printf("#")
                        } else {
                                fmt.Printf(".")
                        }
                }
                fmt.Printf("\n")
        }

	solutions <- fmt.Sprintf(
                "Part 2:\n Dots Left: %d\n",
                dotsLeft,
	)
	close(solutions)
}
