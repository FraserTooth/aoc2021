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
        targetNumsCount := 0
	for line := range instructionChannel {
                messageSections := strings.Split(line, " | ")
		output := strings.Split(messageSections[1], " ")

                for _, number := range output{
                        if len(number) == 2 || len(number) == 4 || len(number) == 3 || len(number) == 7{
                                targetNumsCount++
                        }
                }
	}


	solutions <- fmt.Sprintf(
                "Part 1:\nTimes seeing 1,4,7 or 8: %d",
                targetNumsCount,
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var outputCodes []string
	for line := range instructionChannel {
                messageSections := strings.Split(line, " | ")
		segments := strings.Split(messageSections[0], " ")
		output := strings.Split(messageSections[1], " ")

                var lettersInOneArray []string
                // var lettersInFour []string
                var possible069s []string
                var possible235s []string
                
                // Arrange Letters
                for _, seg := range segments{
                        switch len(seg){
                        case 2: 
                                lettersInOneArray = strings.Split(seg, "")
                        case 6:
                                possible069s = append(possible069s, seg)
                        case 5:
                                possible235s = append(possible235s, seg)
                        }
                }

                // Figure out 3
                lettersIn3 := funk.Filter(possible235s, func(seg string) bool {
                        uniqueAfterCombined := funk.Uniq(append(strings.Split(seg, ""), lettersInOneArray...)).([]string)
                        return len(uniqueAfterCombined) == 5
                }).([]string)[0]
                // Order letters
                lettersIn3Array := strings.Split(lettersIn3,"")
                sort.Strings(lettersIn3Array)
                lettersIn3Sorted := strings.Join(lettersIn3Array, "")
                
                // Figure out 6
                lettersIn6 := funk.Filter(possible069s, func(seg string) bool {
                        uniqueAfterCombined := funk.Uniq(append(strings.Split(seg, ""), lettersInOneArray...)).([]string)
                        return len(uniqueAfterCombined) == 7
                }).([]string)[0]
                // Order letters
                lettersIn6Array := strings.Split(lettersIn6,"")
                sort.Strings(lettersIn6Array)
                lettersIn6Sorted := strings.Join(lettersIn6Array, "")
                
                // Figure out 9
                lettersIn9 := funk.Filter(possible069s, func(seg string) bool {
                        uniqueAfterCombined := funk.Uniq(append(strings.Split(seg, ""), lettersIn3Array...)).([]string)
                        return len(uniqueAfterCombined) == 6
                }).([]string)[0]
                // Order letters
                lettersIn9Array := strings.Split(lettersIn9,"")
                sort.Strings(lettersIn9Array)
                lettersIn9Sorted := strings.Join(lettersIn9Array, "")

                // We now know 0
                lettersIn0 := funk.Filter(possible069s, func(seg string) bool {
                        return !(seg == lettersIn6 || seg == lettersIn9)
                }).([]string)[0]
                // Order letters
                lettersIn0Array := strings.Split(lettersIn0,"")
                sort.Strings(lettersIn0Array)
                lettersIn0Sorted := strings.Join(lettersIn0Array, "")

                // Figure out 5
                lettersIn5 := funk.Filter(possible235s, func(seg string) bool {
                        uniqueAfterCombined := funk.Uniq(append(strings.Split(seg, ""), lettersIn6Array...)).([]string)
                        return len(uniqueAfterCombined) == 6
                }).([]string)[0]
                // Order letters
                lettersIn5Array := strings.Split(lettersIn5,"")
                sort.Strings(lettersIn5Array)
                lettersIn5Sorted := strings.Join(lettersIn5Array, "")

                // We now know 2 
                lettersIn2 := funk.Filter(possible235s, func(seg string) bool {
                        return !(seg == lettersIn3 || seg == lettersIn5)
                }).([]string)[0]
                // Order letters
                lettersIn2Array := strings.Split(lettersIn2,"")
                sort.Strings(lettersIn2Array)
                lettersIn2Sorted := strings.Join(lettersIn2Array, "")

                // Discover

                for i, seg := range output{
                        switch len(seg){
                        case 7:
                                output[i] = "8"
                        case 3: 
                                output[i] = "7"
                        case 4: 
                                output[i] = "4"
                        case 2: 
                                output[i] = "1"
                        }
                        // Order letters
                        segStringArray := strings.Split(seg,"")
                        sort.Strings(segStringArray)
                        segSorted := strings.Join(segStringArray, "")
                        
                        if segSorted == lettersIn3Sorted{
                                output[i] = "3"
                        } else if segSorted == lettersIn6Sorted{
                                output[i] = "6"
                        } else if segSorted == lettersIn9Sorted{
                                output[i] = "9"
                        } else if segSorted == lettersIn0Sorted{
                                output[i] = "0"
                        } else if segSorted == lettersIn5Sorted{
                                output[i] = "5"
                        } else if segSorted == lettersIn2Sorted{
                                output[i] = "2"
                        }
                }

                outputCodes = append(outputCodes, strings.Join(output, ""))
	}

        outputInts := funk.Map(outputCodes, func(x string) int {
                num, _ := strconv.Atoi(x)
                return num
        }).([]int) 

	solutions <- fmt.Sprintf(
                "Part 2:\nTotal of Outputs: %d",
                funk.SumInt(outputInts),
	)
	close(solutions)
}
