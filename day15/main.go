package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
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
        graph:=dijkstra.NewGraph()

        for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                rowInts := funk.Map(rowStrings, func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)
		heightMap = append(heightMap, rowInts)
	}

        directions := [][2]int{{0,1}, {1,0}, {0, -1}, {-1, 0}}

        // Make Vertices
        for y, row := range heightMap{
                for x := range row {
                        locationNum := x+(y*len(heightMap[0]))
                        graph.AddVertex(locationNum)
                }
        }

        // Make Arcs
        for y, row := range heightMap{
                for x := range row {
                        locationNum := x+(y*len(heightMap[0]))
                        for _, diff := range directions{
                                x2 := x+diff[0]
                                y2 := y+diff[1]
                                targetNum := x2+(y2*len(heightMap[0]))
                                // If in range
                                if x2 < len(heightMap[0]) &&
                                x2 >= 0 &&
                                y2 < len(heightMap) &&
                                y2 >= 0 {
                                        graph.AddArc(locationNum, targetNum, int64(heightMap[y2][x2]))
                                }
                        }
                }
        }


        locationOfEnd := len(heightMap[0])-1+((len(heightMap)-1)*len(heightMap[0]))
        best, err := graph.Shortest(0,locationOfEnd)
        if err!=nil{
                log.Fatal(err)
        }
        
	solutions <- fmt.Sprintf(
                "Part 1:\n Minimum Risk Level: %d\n",
                best.Distance,
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var heightMap [][]int
        graph:=dijkstra.NewGraph()

        // Do horizontal Copies
        for line := range instructionChannel {
                rowStrings := strings.Split(line, "")
                var rowInts []int
                for i := 0; i < 5; i++ {
                        scrollingNums := funk.Map(rowStrings, func(x string) int {
                                num, _ := strconv.Atoi(x)
                                shiftedNum := num + i
                                if shiftedNum > 9{
                                        return shiftedNum - 9
                                }
                                return shiftedNum
                        }).([]int)
                        rowInts = append(rowInts, scrollingNums...)
                }
                heightMap = append(heightMap, rowInts)
	}
        // Do vertical copies
        var heightMapBlocks [][]int
        for i := 1; i < 5; i++ {
                for _, row := range heightMap{
                        updatedRow := funk.Map(row, func(num int) int {
                                shiftedNum := num + i
                                if shiftedNum > 9{
                                        return shiftedNum - 9
                                }
                                return shiftedNum
                        }).([]int)
                        heightMapBlocks = append(heightMapBlocks, updatedRow)
                }
        }
        heightMap = append(heightMap, heightMapBlocks...)        

        directions := [][2]int{{0,1}, {1,0}, {0, -1}, {-1, 0}}

        // Make Vertices
        for y, row := range heightMap{
                for x := range row {
                        locationNum := x+(y*len(heightMap[0]))
                        graph.AddVertex(locationNum)
                }
        }

        // Make Arcs
        for y, row := range heightMap{
                for x := range row {
                        locationNum := x+(y*len(heightMap[0]))
                        for _, diff := range directions{
                                x2 := x+diff[0]
                                y2 := y+diff[1]
                                targetNum := x2+(y2*len(heightMap[0]))
                                // If in range
                                if x2 < len(heightMap[0]) &&
                                x2 >= 0 &&
                                y2 < len(heightMap) &&
                                y2 >= 0 {
                                        graph.AddArc(locationNum, targetNum, int64(heightMap[y2][x2]))
                                }
                        }
                }
        }


        locationOfEnd := len(heightMap[0])-1+((len(heightMap)-1)*len(heightMap[0]))
        best, err := graph.Shortest(0,locationOfEnd)
        if err!=nil{
                log.Fatal(err)
        }
        
	solutions <- fmt.Sprintf(
                "Part 2:\n Minimum Risk Level: %d\n",
                best.Distance,
	)
	close(solutions)
}
