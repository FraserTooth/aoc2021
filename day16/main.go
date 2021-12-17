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
        var bits []string

        versionSum := 0

        for line := range instructionChannel {
                chars := strings.Split(line, "")
                for _, char := range chars{
                        integer, _ := strconv.ParseInt(char,16,0)
                        bitString := fmt.Sprintf("%04b", integer)
                        bits = append(bits, strings.Split(bitString, "")...)
                }
	}

        var decodePacket func(packetBits []string) []string
        decodePacket =  func(packetBits []string) []string{
                // Add to version sum
                versionBits := packetBits[0:3]
                versionNum, _ := strconv.ParseInt(strings.Join(versionBits,""),2,0)
                versionSum += int(versionNum)

                // Get Type
                typeBits := packetBits[3:6]
                typeNum, _ := strconv.ParseInt(strings.Join(typeBits,""),2,0)

                remainingBits := packetBits[6:]

                if typeNum == 4{
                        // Literal Numbers
                        finalPoint := 0
                        foundFinalNum := false
                        var literalNumBits []string
                        for i := 5; !foundFinalNum; i+=5 {
                                bitSection := remainingBits[i-5:i]
                                literalNumBits = append(literalNumBits, bitSection[1:]...)
                                finalNumBit := bitSection[0]
                                if finalNumBit == "0"{
                                        foundFinalNum = true
                                }
                                finalPoint = i
                        }
                        return remainingBits[finalPoint:]
                } else {
                        // Operators
                        lengthID := remainingBits[0]
                        // Length of Sub Packets 
                        if lengthID == "0"{
                                totalLengthBits := remainingBits[1:16]
                                totalLengthNum, _ := strconv.ParseInt(strings.Join(totalLengthBits,""),2,0)
                                subPacketsBits := remainingBits[16:16+totalLengthNum]
                                for len(subPacketsBits) > 0{
                                        subPacketsBits = decodePacket(subPacketsBits)
                                }
                                remainingBits = remainingBits[16+totalLengthNum:]
                        // Number of Sub Packets
                        } else {
                                numSubPacketsBits := remainingBits[1:12]
                                numSubPacketsNum, _ := strconv.ParseInt(strings.Join(numSubPacketsBits,""),2,0)
                                subPacketsBits := remainingBits[12:]
                                for i := 0; i < int(numSubPacketsNum); i++ {
                                        subPacketsBits = decodePacket(subPacketsBits)
                                }
                                remainingBits = subPacketsBits
                        }
                }
                return remainingBits
        }

        decodePacket(bits)

	solutions <- fmt.Sprintf(
                "Part 1:\n Version Sum: %d\n",
                versionSum,
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var bits []string

        for line := range instructionChannel {
                chars := strings.Split(line, "")
                for _, char := range chars{
                        integer, _ := strconv.ParseInt(char,16,0)
                        bitString := fmt.Sprintf("%04b", integer)
                        bits = append(bits, strings.Split(bitString, "")...)
                }
	}

        var decodePacket func(packetBits []string)(int, []string)
        decodePacket =  func(packetBits []string)(int, []string){
                // Get Type
                typeBits := packetBits[3:6]
                typeNum, _ := strconv.ParseInt(strings.Join(typeBits,""),2,0)

                remainingBits := packetBits[6:]

                valueSum := 0

                if typeNum == 4{
                        // Literal Numbers
                        finalPoint := 0
                        foundFinalNum := false
                        var literalNumBits []string
                        for i := 5; !foundFinalNum; i+=5 {
                                bitSection := remainingBits[i-5:i]
                                literalNumBits = append(literalNumBits, bitSection[1:]...)
                                finalNumBit := bitSection[0]
                                if finalNumBit == "0"{
                                        foundFinalNum = true
                                }
                                finalPoint = i
                        }
                        literalNum, _ := strconv.ParseInt(strings.Join(literalNumBits,""),2,0)
                        return int(literalNum), remainingBits[finalPoint:]
                } else {
                        var values []int
                        // Operators
                        lengthID := remainingBits[0]
                        // Length of Sub Packets 
                        if lengthID == "0"{
                                totalLengthBits := remainingBits[1:16]
                                totalLengthNum, _ := strconv.ParseInt(strings.Join(totalLengthBits,""),2,0)
                                subPacketsBits := remainingBits[16:16+totalLengthNum]
                                val := 0
                                for len(subPacketsBits) > 0{
                                        val, subPacketsBits = decodePacket(subPacketsBits)
                                        values = append(values, val)
                                }
                                remainingBits = remainingBits[16+totalLengthNum:]
                        // Number of Sub Packets
                        } else {
                                numSubPacketsBits := remainingBits[1:12]
                                numSubPacketsNum, _ := strconv.ParseInt(strings.Join(numSubPacketsBits,""),2,0)
                                subPacketsBits := remainingBits[12:]
                                val := 0
                                for i := 0; i < int(numSubPacketsNum); i++ {
                                        val, subPacketsBits = decodePacket(subPacketsBits)
                                        values = append(values, val)
                                }
                                remainingBits = subPacketsBits
                        }

                        switch typeNum {
                        case 0:
                                valueSum = int(funk.Sum(values))
                        case 1:
                                valueSum = int(funk.Product(values))
                        case 2:
                                valueSum = funk.MinInt(values)
                        case 3:
                                valueSum = funk.MaxInt(values)
                        case 5:
                                if values[0] > values[1]{
                                        valueSum = 1
                                } else {
                                        valueSum = 0
                                }
                        case 6:
                                if values[0] < values[1]{
                                        valueSum = 1
                                } else {
                                        valueSum = 0
                                }
                        case 7:
                                if values[0] == values[1]{
                                        valueSum = 1
                                } else {
                                        valueSum = 0
                                }
                        }
                        
                }
                return valueSum, remainingBits
        }

        val, _ := decodePacket(bits)

	solutions <- fmt.Sprintf(
                "Part 1:\n Value Sum: %d\n",
                val,
	)
	close(solutions)
}
