package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	// Handle Input
	var instructions []string
	for line := range instructionChannel {
		instructions = append(instructions, line)
	}

	// Separate Data
	numbersCalled := strings.Split(instructions[0],",")
	boardsInput := instructions[2:]
	var bingoBoards [][][]string

	for i := 0; i<len(boardsInput); i+=6{
		var board [][]string
		for j := 0; j<5; j++{
			inputIndex := i+j
			board = append(board, strings.Fields(boardsInput[inputIndex]))	
		}
		bingoBoards = append(bingoBoards, board)
	}

	isWinner := false
	var winningBoard int
	numberIndex := 0
	finalNumberCalled := 0
	for isWinner == false {
		numberCalled := numbersCalled[numberIndex]

		for boardNum, board := range bingoBoards{
			colPossible := false
			colMark := 0
			colCount := 0
			for rowI, row := range board {
				rowCount := 0
				for i, numberWritten := range row{
					// Mark Called
					if numberWritten == numberCalled{
						row[i] = "X"
					}
					// Count Xs
					if row[i] == "X"{
						rowCount++
						if rowI == 0 {
							colPossible = true
							colMark = i
							colCount++
						} else if colPossible && colMark == i {
							colCount++
						}
					}
					// End Game
					if rowCount == 5 || colCount == 5{
						winningBoard = boardNum
						finalNumberCalled, _ = strconv.Atoi(numberCalled)
						isWinner = true
					}
				}

			}
			
		}

		numberIndex++
	}

	// Calculate Score
	sum := 0
	for _, row := range bingoBoards[winningBoard] {
		for _, numberWritten := range row{
			if numberWritten != "X"{
				numberInt, _ := strconv.Atoi(numberWritten)
				sum += numberInt
			} 
		}
	}
	score := finalNumberCalled * sum 

	solutions <- fmt.Sprintf(
					"Part 1:\nBoard Num: %d\nWinning Number: %d\nScore: %d\nWinning Board: %s\n",
					winningBoard,
					finalNumberCalled,
					score,
					bingoBoards[winningBoard],
	)
	close(solutions)
}


func partTwo(instructionChannel <-chan string, solutions chan<- string) {
	// Handle Input
	var instructions []string
	for line := range instructionChannel {
		instructions = append(instructions, line)
	}

	// Separate Data
	numbersCalled := strings.Split(instructions[0],",")
	boardsInput := instructions[2:]
	var bingoBoards [][][]string

	for i := 0; i<len(boardsInput); i+=6{
		var board [][]string
		for j := 0; j<5; j++{
			inputIndex := i+j
			board = append(board, strings.Fields(boardsInput[inputIndex]))	
		}
		bingoBoards = append(bingoBoards, board)
	}

	winners := make(map[int]bool)
	finalWinningBoard := 0
	numberIndex := 0
	finalNumberCalled := 0
	for len(winners) < len(bingoBoards) {
		numberCalled := numbersCalled[numberIndex]

		for boardNum, board := range bingoBoards{
			colCounts := []int{0,0,0,0,0}
			for _, row := range board {
				rowCount := 0
				for i, numberWritten := range row{
					// Mark Called
					if numberWritten == numberCalled{
						row[i] = "X"
					}
					// Count Xs
					if row[i] == "X"{
						rowCount++
						colCounts[i]++
					}
					isWinner := winners[boardNum]
					// End Game
					if (rowCount == 5 || containsFive(colCounts)) && !isWinner{
						winners[boardNum] = true
						finalWinningBoard = boardNum
						finalNumberCalled, _ = strconv.Atoi(numberCalled)
					}
				}

			}
		}
		numberIndex++
	}

	// Calculate Score
	sum := 0
	for _, row := range bingoBoards[finalWinningBoard] {
		for _, numberWritten := range row{
			if numberWritten != "X"{
				numberInt, _ := strconv.Atoi(numberWritten)
				sum += numberInt
			} 
		}
	}
	score := finalNumberCalled * sum 

	solutions <- fmt.Sprintf(
					"Part 2:\nBoard Num: %d\nWinning Number: %d\nScore: %d\nWinning Board: %s\n",
					finalWinningBoard,
					finalNumberCalled,
					score,
					bingoBoards[finalWinningBoard],
	)
	close(solutions)
}

func containsFive(s []int) bool {
	for _, v := range s {
		if v == 5 {
			return true
		}
	}

	return false
}