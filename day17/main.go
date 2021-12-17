package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

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



func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var xBoundary [2]int
        var yBoundary [2]int

        highestY := 0

        for line := range instructionChannel {
                pattern := `(-?\d+)`
                r, _ := regexp.Compile(pattern)
                numberStrings := r.FindAllString(line, -1)
                numbers := funk.Map(numberStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
                xBoundary = [2]int{numbers[0], numbers[1]}
                yBoundary = [2]int{numbers[2], numbers[3]}
	}

        // Boundary of Sim
        xInitialMin := 0
        xInitialMax := 1000
        yInitialMin := 0
        yInitialMax := 1000

        // For a variety of Initial Sims
        for xInitialVel := xInitialMin; xInitialVel < xInitialMax; xInitialVel++ {
                for yInitialVel := yInitialMin; yInitialVel < yInitialMax; yInitialVel++ {
                        highestYInSim := 0
                        simDone := false
                        // Run Simulation
                        xPos := 0
                        yPos := 0
                        xVel := xInitialVel
                        yVel := yInitialVel
                        // While Higher than the minimum y value
                        for !simDone{
                                // Move
                                xPos += xVel
                                yPos += yVel
                                if yPos > highestYInSim{
                                        highestYInSim = yPos
                                }
                                // Check If Hit
                                if (xPos >= xBoundary[0] && xPos <= xBoundary[1]) &&
                                (yPos >= yBoundary[0] && yPos <= yBoundary[1]) {
                                        simDone = true
                                        if highestYInSim > highestY{
                                                highestY = highestYInSim
                                        }
                                }
                                // Apply Drag and Gravity
                                if xVel > 0{
                                        xVel--
                                }
                                yVel--

                                // Check to Stop Infinite Fall
                                if yPos < funk.MinInt(yBoundary[:]){
                                        simDone = true
                                }
                        }
                }
        }


	solutions <- fmt.Sprintf(
                "Part 1:\n Highest Y: %d\n",
                highestY,
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var xBoundary [2]int
        var yBoundary [2]int

        highestY := 0
        validValues := 0

        for line := range instructionChannel {
                pattern := `(-?\d+)`
                r, _ := regexp.Compile(pattern)
                numberStrings := r.FindAllString(line, -1)
                numbers := funk.Map(numberStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
                xBoundary = [2]int{numbers[0], numbers[1]}
                yBoundary = [2]int{numbers[2], numbers[3]}
	}

        // Boundary of Sim
        xInitialMin := -600
        xInitialMax := 600
        yInitialMin := -600
        yInitialMax := 600

        // For a variety of Initial Sims
        for xInitialVel := xInitialMin; xInitialVel < xInitialMax; xInitialVel++ {
                for yInitialVel := yInitialMin; yInitialVel < yInitialMax; yInitialVel++ {
                        highestYInSim := 0
                        simDone := false
                        // Run Simulation
                        xPos := 0
                        yPos := 0
                        xVel := xInitialVel
                        yVel := yInitialVel
                        // While Higher than the minimum y value
                        for !simDone{
                                // Move
                                xPos += xVel
                                yPos += yVel
                                if yPos > highestYInSim{
                                        highestYInSim = yPos
                                }
                                // Check If Hit
                                if (xPos >= xBoundary[0] && xPos <= xBoundary[1]) &&
                                (yPos >= yBoundary[0] && yPos <= yBoundary[1]) {
                                        simDone = true
                                        validValues++
                                        if highestYInSim > highestY{
                                                highestY = highestYInSim
                                        }
                                }
                                // Apply Drag and Gravity
                                if xVel > 0{
                                        xVel--
                                }
                                yVel--

                                // Check to Stop Infinite Fall
                                if yPos < funk.MinInt(yBoundary[:]){
                                        simDone = true
                                }
                        }
                }
        }


	solutions <- fmt.Sprintf(
                "Part 2:\n Highest Y: %d\n Valid Values: %d",
                highestY,
                validValues,
	)
	close(solutions)
}
