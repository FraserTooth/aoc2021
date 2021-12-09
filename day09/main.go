package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
        lowPoints := 0
        totalRiskLevel := 0
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		heightMap = append(heightMap, rowInts)
	}

        for y, row := range heightMap {
                for x := range row {
                        pointHeight := heightMap[y][x]

                        isLowest := true
                        // Check Above
                        if (y > 0 && pointHeight >= heightMap[y-1][x]) || 
                        // Check Left
                        (x > 0 && pointHeight >= heightMap[y][x-1]) ||
                        // Check Below
                        (y < len(heightMap) - 1 && pointHeight >= heightMap[y+1][x]) ||
                        // Check Right
                        (x < len(heightMap[0]) - 1 && pointHeight >= heightMap[y][x+1]) {
                                isLowest = false
                        }

                        if isLowest{
                                lowPoints++
                                totalRiskLevel += pointHeight+1
                        }
                }
        }

	solutions <- fmt.Sprintf(
                "Part 1:\nLow Points: %d\nTotal Risk Level: %d",
                lowPoints,
                totalRiskLevel,
	)
	close(solutions)
}

func visitPoint(x int, y int, heightMap [][]int, pointsVisited []string) []string{
        pointHeight := heightMap[y][x]

        pointsVisited = append(pointsVisited, fmt.Sprintf("%d,%d", x, y))
        // Go Up
        if !funk.Contains(pointsVisited, fmt.Sprintf("%d,%d", x, y-1)) && 
        y > 0 && 
        pointHeight < heightMap[y-1][x] && 
        heightMap[y-1][x] != 9{
                pointsVisited = visitPoint(x, y-1, heightMap, pointsVisited)
        } 

        // Go Left
        if !funk.Contains(pointsVisited, fmt.Sprintf("%d,%d", x-1, y)) && 
        x > 0 && 
        pointHeight < heightMap[y][x-1] && 
        heightMap[y][x-1] != 9{
                pointsVisited = visitPoint(x-1, y, heightMap, pointsVisited)
        }

        // Go Down
        if !funk.Contains(pointsVisited, fmt.Sprintf("%d,%d", x, y+1)) &&
        y < len(heightMap) - 1 &&
        pointHeight < heightMap[y+1][x] &&
        heightMap[y+1][x] != 9{
                pointsVisited = visitPoint(x, y+1, heightMap, pointsVisited)
        }

        // Go Right
        if !funk.Contains(pointsVisited, fmt.Sprintf("%d,%d", x+1, y)) &&
        x < len(heightMap[0]) - 1 &&
        pointHeight < heightMap[y][x+1] &&
        heightMap[y][x+1] != 9{
                pointsVisited = visitPoint(x+1, y, heightMap, pointsVisited)
        }

        return pointsVisited
}

func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var heightMap [][]int
        var basins []int
	for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		heightMap = append(heightMap, rowInts)
	}

        for y, row := range heightMap {
                for x := range row {
                        pointHeight := heightMap[y][x]

                        isLowest := true
                        // Check Above
                        if (y > 0 && pointHeight >= heightMap[y-1][x]) || 
                        // Check Left
                        (x > 0 && pointHeight >= heightMap[y][x-1]) ||
                        // Check Below
                        (y < len(heightMap) - 1 && pointHeight >= heightMap[y+1][x]) ||
                        // Check Right
                        (x < len(heightMap[0]) - 1 && pointHeight >= heightMap[y][x+1]) {
                                isLowest = false
                        }


                        if isLowest{
                                var pointsVisited []string
                                pointsVisited = visitPoint(x,y, heightMap, pointsVisited)
                                basins = append(basins, len(pointsVisited))
                        }
                }
        }

        sort.Ints(basins)

	solutions <- fmt.Sprintf(
                "Part 2:\nBasins: %d\nSum of Top Three: %d",
                len(basins),
                int(funk.Product(basins[len(basins)-3:])),
	)
	close(solutions)
}
