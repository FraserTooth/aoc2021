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

func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var heightMap [][]int
        minimumRiskLevel := 1000000000
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		heightMap = append(heightMap, rowInts)
	}

        directions := [][2]int{{0,1}, {1,0}}

        var navigate func(location [2]int, visited []string, totalRiskSoFar int)
        navigate = func (location [2]int, visited []string, totalRiskSoFar int){
                visited =append(visited, locationToString(location))
                risk := heightMap[location[1]][location[0]]
                totalRiskSoFar += risk
                // End
                if location[0] == len(heightMap[0])-1 && location[1] == len(heightMap)-1{
                        if totalRiskSoFar < minimumRiskLevel {
                                minimumRiskLevel = totalRiskSoFar
                        }
                } else {
                        for _, diff := range directions{
                                target := [2]int{location[0]+diff[0], location[1]+diff[1]}
                                // If in range
                                if target[0] < len(heightMap[0]) &&
                                target[0] >= 0 &&
                                target[1] < len(heightMap) &&
                                target[1] >= 0 {
                                        targetString := locationToString(target)
                                        if !funk.Contains(visited, targetString){
                                                navigate(target, visited, totalRiskSoFar)
                                        }
                                }
                        }
                }
        }

        navigate([2]int{0,0}, make([]string,0), 0)

	solutions <- fmt.Sprintf(
                "Part 1:\n Minimum Risk Level: %d\n",
                minimumRiskLevel - heightMap[0][0],
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var heightMap [][]int
        minimumRiskLevel := 1000000000
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		heightMap = append(heightMap, rowInts)
	}

	solutions <- fmt.Sprintf(
                "Part 2:\n Minimum Risk Level: %d\n",
                minimumRiskLevel,
	)
	close(solutions)
}
