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
	overlaps := 0
        // Make a 1000x1000 grid of zeros
        area := make([][]int, 1000)
        for i := range area{
                area[i] = make([]int, 1000)
        }
	for line := range instructionChannel {
                lineParts := strings.Fields(line)
                start := funk.Map(strings.Split(lineParts[0], ","), func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)      
                end := funk.Map(strings.Split(lineParts[2], ","), func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int) 
                
                // If not a horizontal or vertical line
                if start[0] != end[0] && start[1] != end[1]{
                        continue
                }

                direction := []int{0,0}

                // If Vertical
                if start[0] == end[0]{
                        // If going Down
                        if start[1] < end[1]{
                                direction = []int{0,1}
                        } else {
                                direction = []int{0,-1}
                        }
                } else {
                        // If going Right
                        if start[0] < end[0]{
                                direction = []int{1,0}
                        } else {
                                direction = []int{-1,0}
                        }
                }


                currentPosition := start
                for !(currentPosition[0] == end[0] && currentPosition[1] == end[1]){
                        area[currentPosition[1]][currentPosition[0]]++
                        if area[currentPosition[1]][currentPosition[0]] == 2{
                                overlaps++
                        }
                        currentPosition = []int{currentPosition[0]+direction[0], currentPosition[1]+direction[1]}
                }

                // Do End Position
                area[currentPosition[1]][currentPosition[0]]++
                if area[currentPosition[1]][currentPosition[0]] == 2{
                        overlaps++
                }
	}



	solutions <- fmt.Sprintf(
                "Part 1:\nOverlaps: %d\n",
                overlaps,
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        overlaps := 0
        // Make a 1000x1000 grid of zeros
        area := make([][]int, 1000)
        for i := range area{
                area[i] = make([]int, 1000)
        }
	for line := range instructionChannel {
                lineParts := strings.Fields(line)
                start := funk.Map(strings.Split(lineParts[0], ","), func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int)      
                end := funk.Map(strings.Split(lineParts[2], ","), func(x string) int {
                        num, _ := strconv.Atoi(x)
                        return num
                }).([]int) 


                direction := []int{0,0}

                // If Vertical
                if start[0] == end[0]{
                        // If going Down
                        if start[1] < end[1]{
                                direction = []int{0,1}
                        } else {
                                direction = []int{0,-1}
                        }
                // If Horizontal
                } else if start[1] == end[1] {
                        // If going Right
                        if start[0] < end[0]{
                                direction = []int{1,0}
                        } else {
                                direction = []int{-1,0}
                        }
                // If Diagonal
                } else {
                        // If going Up Right
                        if start[0] < end[0] && start[1] < end[1]{
                                direction = []int{1,1}
                        // Up Left
                        } else if start[0] > end[0] && start[1] < end[1] {
                                direction = []int{-1,1}
                        // Down Right
                        } else if start[0] < end[0] && start[1] > end[1] {
                                direction = []int{1,-1}
                        // Down Left
                        } else {
                                direction = []int{-1,-1}
                        }
                }


                currentPosition := start
                for !(currentPosition[0] == end[0] && currentPosition[1] == end[1]){
                        area[currentPosition[1]][currentPosition[0]]++
                        if area[currentPosition[1]][currentPosition[0]] == 2{
                                overlaps++
                        }
                        currentPosition = []int{currentPosition[0]+direction[0], currentPosition[1]+direction[1]}
                }

                // Do End Position
                area[currentPosition[1]][currentPosition[0]]++
                if area[currentPosition[1]][currentPosition[0]] == 2{
                        overlaps++
                }
	}


	solutions <- fmt.Sprintf(
                "Part 2:\nOverlaps: %d\n",
                overlaps,
	)
	close(solutions)
}
