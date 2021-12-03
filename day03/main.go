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

func partOne(instructions <-chan string, solutions chan<- string) {
	var bitCount []int
	for instruction := range instructions {
					bits := strings.Split(instruction, "")
					for i, bit := range bits {
						if len(bitCount) < i + 1{
							bitCount = append(bitCount, 0)
						}
						if bit == "1"{
							bitCount[i]++
						} else {
							bitCount[i]--
						}
					}
	}

	// Parse the counts
	var gamma  []string
	var epsilon  []string

	for _, count := range bitCount{
		if count > 0{
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}

	gammaBinString := strings.Join(gamma, "")
	gammaInt, _ := strconv.ParseInt(gammaBinString, 2, 64)
	epsilonBinString := strings.Join(epsilon, "")
	epsilonInt, _ := strconv.ParseInt(epsilonBinString, 2, 64)

	fmt.Println()

	solutions <- fmt.Sprintf(
					"Part 1:\nGamma: %s\nEpsilon: %s\nSolution: %d\n",
					gammaBinString,
					epsilonBinString,
					gammaInt * epsilonInt,
	)
	close(solutions)
}

func StrArrayEquals(a []string, b []string) bool {
	return strings.Join(a, "") == strings.Join(b, "")
}

func partTwo(instructionChannel <-chan string, solutions chan<- string) {
	var oxygen  []string
	var c02  []string
	var instructions []string

	for line := range instructionChannel {
		instructions = append(instructions, line)
	}

	for len(oxygen) < 12{
		oxygenCount := 0
		var oxygenValids [][]string
		c02Count := 0
		var c02Valids [][]string
		targetIndex := len(oxygen)
		
		for _, instruction := range instructions {
			bits := strings.Split(instruction, "")
			targetBit := bits[targetIndex]

			if targetBit == "1"{
				if StrArrayEquals(oxygen[:targetIndex], bits[:targetIndex]){
					oxygenValids = append(oxygenValids, bits)
					oxygenCount++
				}
				if StrArrayEquals(c02[:targetIndex], bits[:targetIndex]){
					c02Valids = append(c02Valids, bits)
					c02Count++
				}
			} else {
				if StrArrayEquals(oxygen[:targetIndex], bits[:targetIndex]){
					oxygenValids = append(oxygenValids, bits)
					oxygenCount--
				}
				if StrArrayEquals(c02[:targetIndex], bits[:targetIndex]){
					c02Valids = append(c02Valids, bits)
					c02Count--
				}
			}
		}
		

		if oxygenCount >= 0{
			oxygen = append(oxygen, "1")
		} else {
			oxygen = append(oxygen, "0")
		}
		
		if c02Count >= 0{
			c02 = append(c02, "0")
		} else {
			c02 = append(c02, "1")
		}
		if len(oxygenValids) == 1{
			oxygen = oxygenValids[0]
		}
		if len(c02Valids) == 1{
			c02 = c02Valids[0]
		}
	}
	

	oxygenBinString := strings.Join(oxygen, "")
	oxygenInt, _ := strconv.ParseInt(oxygenBinString, 2, 64)
	c02BinString := strings.Join(c02, "")
	c02Int, _ := strconv.ParseInt(c02BinString, 2, 64)

	fmt.Println()

	solutions <- fmt.Sprintf(
					"Part 2:\nOxygen: %s\nC02: %s\nSolution: %d\n",
					oxygenBinString,
					c02BinString,
					oxygenInt * c02Int,
	)
	close(solutions)
}