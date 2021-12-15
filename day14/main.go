package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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


func printUniqueValue( arr []string) map[string]int{
        //Create a   dictionary of values for each element
        dict:= make(map[string]int)
        for _ , num :=  range arr {
            dict[num] = dict[num]+1
        }
        return dict
    }


func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var polymer []string
        var rules = make(map[string]string)
	for line := range instructionChannel {
                // Skip empty line
                if len(line)==0{
                        continue
                }
                // Template
                if !strings.Contains(line, "->"){
                        polymer = strings.Split(line, "")
                // Dot
                } else {
                        ruleBits := strings.Split(line, " -> ")
                        rules[ruleBits[0]] = ruleBits[1]
                }
                
	}

        // Run Steps
        for step := 0; step < 10; step++ {

                for x := 0; x < len(polymer)-1; x++ {
                        charPair := polymer[x]+polymer[x+1]
                        insert, matchesRule := rules[charPair]
                        if matchesRule{
                                polymer = append(polymer[:x+1], append([]string{insert}, polymer[x+1:]...)...)
                                x++
                        }
                }
                
        }

        counts := printUniqueValue(polymer)

        biggest := 0
        smallest := 100000
        for _, count := range counts{
                if count > biggest{
                        biggest=count
                }
                if count < smallest {
                        smallest=count
                }
        }

	solutions <- fmt.Sprintf(
                "Part 1:\n Biggest Minus Smallest: %d\n",
                biggest-smallest,
	)
	close(solutions)
}



func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var template []string
        var polymer = make(map[string]int)
        var rules = make(map[string]string)
        var charCount = make(map[string]int)
	for line := range instructionChannel {
                // Skip empty line
                if len(line)==0{
                        continue
                }
                // Template
                if !strings.Contains(line, "->"){
                        template = strings.Split(line, "")
                        for x := 0; x < len(template)-1; x++ {
                                charPair := template[x]+template[x+1] 
                                polymer[charPair]++
                        }
                        // Add chars
                        for _,char := range template{
                                charCount[char]++
                        }
                // Dot
                } else {
                        ruleBits := strings.Split(line, " -> ")
                        rules[ruleBits[0]] = ruleBits[1]
                }
                
	}

        // Run Steps
        for step := 0; step < 40; step++ {
                var newPolymer = make(map[string]int)
                for pair := range polymer{
                        insert, matchesRule := rules[pair]
                        if matchesRule{
                                times := polymer[pair]
                                chars := strings.Split(pair, "")
                                // Do Left
                                leftPair := chars[0]+insert
                                newPolymer[leftPair]+=times
                                // Do right
                                rightPair := insert+chars[1]
                                newPolymer[rightPair]+=times
                                // Count New Chars
                                charCount[insert]+=times
                        } else {
                                // Otherwise, maintain the same number
                                newPolymer[pair] = polymer[pair]
                        }
                }
                polymer = newPolymer
        }

        fmt.Println(charCount)

        biggest := 0
        smallest := 10000000000000000
        for _, count := range charCount{
                if count > biggest{
                        biggest=count
                }
                if count < smallest {
                        smallest=count
                }
        }

	solutions <- fmt.Sprintf(
                "Part 2:\n Biggest Minus Smallest: %d\n",
                biggest-smallest,
	)
	close(solutions)
}
