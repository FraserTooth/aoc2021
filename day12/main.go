package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Cave struct {  
        name string
        big bool
        connections []string
}

func findPaths(currentCave *Cave, caves map[string]*Cave, cavesVisited []string, paths [][]string) [][]string{
        cavesVisited = append(cavesVisited, currentCave.name)
        for _, targetCaveName := range currentCave.connections{
                targetCave := caves[targetCaveName]

                if targetCaveName == "end"{
                        paths = append(paths, append(cavesVisited, "end"))
                } else if targetCave.big || ( 
                !funk.Contains(cavesVisited, targetCaveName) &&
                targetCaveName != "start"){
                        paths = findPaths(targetCave, caves, cavesVisited, paths)
                }
        }
        return paths
}


func partOne(instructionChannel <-chan string, solutions chan<- string) {
        var caves = make(map[string]*Cave)

        // Build cave graph
	for line := range instructionChannel {
                pathLocations := strings.Split(line, "-")
                // Get From and To
                fromName := pathLocations[0]
                toName := pathLocations[1]
                _, hasFromCave := caves[fromName]
                _, hasToCave := caves[toName]

                // Create if They don't exist yet
                if !hasFromCave{
                        caves[fromName] = &Cave{
                                name: fromName,
                                big: strings.ToUpper(fromName) == fromName,
                                connections: make([]string, 0, 10),
                        }
                }
                if !hasToCave{
                        caves[toName] = &Cave{
                                name: toName,
                                big: strings.ToUpper(toName) == toName,
                                connections: make([]string, 0, 10),
                        } 
                }
                
                // Make Connections
                caves[fromName].connections = append(caves[fromName].connections, toName) 
                caves[toName].connections = append(caves[toName].connections, fromName) 
	}

        // Find Paths
        var paths [][]string
        paths = findPaths(caves["start"], caves, make([]string, 0), paths)

	solutions <- fmt.Sprintf(
                "Part 1:\n Paths thru cave: %d\n",
                len(paths),
	)
	close(solutions)
}

func findPathsv2(currentCave *Cave, caves map[string]*Cave, cavesVisited []string, paths [][]string) [][]string{
        // Add visit
        cavesVisited = append(cavesVisited, currentCave.name)
        if currentCave.name == "end"{
                return append(paths, cavesVisited)
        }

        
        for _, targetCaveName := range currentCave.connections{
                if targetCaveName == "start"{
                        continue
                }

                targetCave := caves[targetCaveName]

                smallCaveVisits := funk.Filter(cavesVisited, func (caveName string) bool{
                        return !caves[caveName].big
                }).([]string)
                hasVisitedAnySmallCaveTwice := len(smallCaveVisits) > len(funk.Uniq(smallCaveVisits).([]string))
                canVisitThisSmallCave := !funk.Contains(cavesVisited, targetCaveName) || !hasVisitedAnySmallCaveTwice

                if (targetCave.big || canVisitThisSmallCave) && !funk.Contains(cavesVisited, "end"){
                        paths = findPathsv2(targetCave, caves, cavesVisited, paths)
                }
        }
        return paths
}

func partTwo(instructionChannel <-chan string, solutions chan<- string) {
        var caves = make(map[string]*Cave)

        // Build cave graph
	for line := range instructionChannel {
                pathLocations := strings.Split(line, "-")
                // Get From and To
                fromName := pathLocations[0]
                toName := pathLocations[1]
                _, hasFromCave := caves[fromName]
                _, hasToCave := caves[toName]

                // Create if They don't exist yet
                if !hasFromCave{
                        caves[fromName] = &Cave{
                                name: fromName,
                                big: strings.ToUpper(fromName) == fromName,
                                connections: make([]string, 0, 10),
                        }
                }
                if !hasToCave{
                        caves[toName] = &Cave{
                                name: toName,
                                big: strings.ToUpper(toName) == toName,
                                connections: make([]string, 0, 10),
                        } 
                }
                
                // Make Connections
                caves[fromName].connections = append(caves[fromName].connections, toName) 
                caves[toName].connections = append(caves[toName].connections, fromName) 
	}


        // Find Paths
        var paths [][]string
        paths = findPathsv2(caves["start"], caves, make([]string, 0), paths)

	solutions <- fmt.Sprintf(
                "Part 2:\n Paths thru cave: %d\n",
                len(paths),
	)
	close(solutions)
}
