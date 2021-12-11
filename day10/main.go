package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

var bracketMap = map[string]string {
        "{": "}",
        "(": ")",
        "[": "]",
        "<": ">",
}



func partOne(instructionChannel <-chan string, solutions chan<- string) {
        errorScore := 0
        var scoreMap = map[string]int {
                "}": 1197,
                ")": 3,
                "]": 57,
                ">": 25137,
        }
	for line := range instructionChannel {
                var chunkTracker []string
                commands := strings.Split(line, "")

                for _, command := range commands{
                        // If opener
                        if funk.Contains(funk.Keys(bracketMap), command){
                                chunkTracker = append(chunkTracker, command)
                        // If closer
                        } else {
                                needsClosing := chunkTracker[len(chunkTracker)-1]
                                requiredCloser := bracketMap[needsClosing]
                                if requiredCloser == command{
                                        if len(chunkTracker) > 0{
                                                chunkTracker = chunkTracker[:len(chunkTracker)-1]
                                        }
                                        
                                } else {
                                        // Illegal Character
                                        errorScore += scoreMap[command]
                                        break
                                }
                        }
                }
                
	}

	solutions <- fmt.Sprintf(
                "Part 1:\nSyntax Error Score: %d\n",
                errorScore,
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var completionScores []int
        var scoreMap = map[string]int {
                "}": 3,
                ")": 1,
                "]": 2,
                ">": 4,
        }
	for line := range instructionChannel {
                var chunkTracker []string
                commands := strings.Split(line, "")
                lineCompletionScore := 0
                hasIllegalChar := false

                for _, command := range commands{
                        // If opener
                        if funk.Contains(funk.Keys(bracketMap), command){
                                chunkTracker = append(chunkTracker, command)
                        // If closer
                        } else {
                                needsClosing := chunkTracker[len(chunkTracker)-1]
                                requiredCloser := bracketMap[needsClosing]
                                if requiredCloser == command{
                                        if len(chunkTracker) > 0{
                                                chunkTracker = chunkTracker[:len(chunkTracker)-1]
                                        }
                                        
                                } else {
                                        // Illegal Character, ignoring these
                                        hasIllegalChar = true
                                        break
                                }
                        }
                }
                // If incomplete
                if !hasIllegalChar && len(chunkTracker) > 0{
                        completionArray := funk.Map(funk.Reverse(chunkTracker), func(x string) string {
                                return bracketMap[x]
                        }).([]string)

                        for _, closer := range completionArray{
                                lineCompletionScore = lineCompletionScore * 5
                                lineCompletionScore += scoreMap[closer]
                        }
                        completionScores = append(completionScores, lineCompletionScore)
                }
                
	}

        sort.Ints(completionScores)
        middleCompletionScore := completionScores[int(math.Floor(float64(len(completionScores)/2)))]

	solutions <- fmt.Sprintf(
                "Part 2:\n Middle Completion Score: %d\n",
                middleCompletionScore,
	)
	close(solutions)
}
