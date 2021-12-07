package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/montanaflynn/stats"
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
        var input []string
	for line := range instructionChannel {
		input = append(input, line)
	}
        inputInt := funk.Map(strings.Split(input[0], ","), func(x string) int {
                num, _ := strconv.Atoi(x)
                return num
        }).([]int)  
        crabLocations := stats.LoadRawData(inputInt)

        medianLocation, _ := stats.Median(crabLocations)

        absoluteDeviationFromMedian := 0.0
        for _, crabLocation := range crabLocations{
                absoluteDeviationFromMedian+=math.Abs(crabLocation-medianLocation)
        }


	solutions <- fmt.Sprintf(
                "Part 1:\nCrabs Aligned to: %d\nFuel Spent: %d",
                int(medianLocation),
                int(absoluteDeviationFromMedian),
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var input []string
	for line := range instructionChannel {
		input = append(input, line)
	}
        crabLocations := funk.Map(strings.Split(input[0], ","), func(x string) int {
                num, _ := strconv.Atoi(x)
                return num
        }).([]int)  

        var triangleDistances []int
        for i := 0; i < 10000; i++ {
                triangleDistances = append(triangleDistances, i*(i+1)/2)
        } 
        
        minimumTotalFuel := 100000000000
        minimumTotalFuelLocation := 0
        for target := funk.MinInt(crabLocations); target <= funk.MaxInt(crabLocations); target++ {
                totalFuel := 0
                for _, location := range crabLocations{
                        totalFuel += triangleDistances[int(math.Abs(float64(target-location)))]   
                }
                if totalFuel < minimumTotalFuel{
                        minimumTotalFuel = totalFuel
                        minimumTotalFuelLocation = target
                }
        }


	solutions <- fmt.Sprintf(
                "Part 2:\nCrabs Aligned to: %d\nFuel Spent: %d",
                minimumTotalFuelLocation,
                minimumTotalFuel,
	)
	close(solutions)
}
