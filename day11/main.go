package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
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


func flash(x int, y int, octopi [][]int){
        diffs := []int{-1,0,1}

        for _, xDiff := range diffs{  
                for _, yDiff := range diffs{
                                targetX := x+xDiff
                                targetY := y+yDiff
                                // If in bounds
                                if targetX >= 0 &&
                                targetX < len(octopi[0]) &&
                                targetY >=0 &&
                                targetY < len(octopi) {
                                        octopi[targetY][targetX]++
                                        if octopi[targetY][targetX] == 10 {
                                                flash(targetX, targetY, octopi)
                                        }
                                }
                        } 
        }
}


func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var octopi [][]int
        flashes := 0
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		octopi = append(octopi, rowInts)
	}

        for i := 0; i < 100; i++ {
                // Increment
                for y, row := range octopi{
                        for x := range row{
                                octopi[y][x]++  
                                // If Flashed
                                if octopi[y][x] == 10{
                                        flash(x,y,octopi)
                                }
                                }   
                }                
                

                // Reset Flashed
                for y, row := range octopi{
                        for x, octopus := range row{
                                if octopus > 9{
                                        flashes++
                                        octopi[y][x] = 0  
                                }
                                }  
                }
        }

	solutions <- fmt.Sprintf(
                "Part 1:\n Flashes after 100 steps: %d\n",
                flashes,
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var octopi [][]int
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		octopi = append(octopi, rowInts)
	}

        totalToFlash := len(octopi)*len(octopi[0])
        stepWhereAllFlashed := 0

        for i := 0; i < 1000; i++ {
                // Increment
                for y, row := range octopi{
                        for x := range row{
                                octopi[y][x]++  
                                // If Flashed
                                if octopi[y][x] == 10{
                                        flash(x,y,octopi)
                                }
                                }   
                }                
                

                // Reset Flashed
                totalFlashedThisRound := 0 
                for y, row := range octopi{
                        for x, octopus := range row{
                                if octopus > 9{
                                        totalFlashedThisRound++
                                        octopi[y][x] = 0  
                                }
                                }  
                }

                if totalFlashedThisRound == totalToFlash{
                        stepWhereAllFlashed = i+1
                        break
                }
        }

	solutions <- fmt.Sprintf(
                "Part 2:\n All Flash at step: %d\n",
                stepWhereAllFlashed,
	)
	close(solutions)
}
