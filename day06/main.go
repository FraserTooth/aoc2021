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
        var fish []int
        var input []string
	for line := range instructionChannel {
		input = append(input, line)
	}
        fish = funk.Map(strings.Split(input[0], ","), func(x string) int {
                num, _ := strconv.Atoi(x)
                return num
        }).([]int)  
	
        for day := 0; day < 80; day++ {
                var newFish []int
                for i := range fish {
                        fish[i]--
                        if fish[i] < 0{
                                fish[i]=6
                                newFish = append(newFish, 8)
                        }
                }
                fish = append(fish, newFish...)
        }


	solutions <- fmt.Sprintf(
                "Part 1:\nNumber of Fish: %d\n",
                len(fish),
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        fishMap := make(map[int]int)
        var input []string
	for line := range instructionChannel {
		input = append(input, line)
	}
        for _, num := range strings.Split(input[0], ","){
                age, _ := strconv.Atoi(num)
                fishMap[age]++
        }

	
        for day := 0; day < 256; day++ {
                newFishMap := make(map[int]int)
                for age := range fishMap {
                        if age == 0{
                                newFishMap[8]+= fishMap[0]
                                newFishMap[6]+= fishMap[0]
                        } else {
                                newFishMap[age-1] += fishMap[age]
                        }
                }
                fishMap = newFishMap
        }

        // Sum Up
        totalFish := 0
        for _, age := range fishMap {
                totalFish+=age
        }


	solutions <- fmt.Sprintf(
                "Part 2:\nNumber of Fish: %d\n",
                totalFish,
	)
	close(solutions)
}
