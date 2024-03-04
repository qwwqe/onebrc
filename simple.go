package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SimpleStrategy struct{}

func (s SimpleStrategy) Process(filename string) {
	tempCount := map[string]int{}
	tempSum := map[string]float64{}
	tempMin := map[string]float64{}
	tempMax := map[string]float64{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, ";")
		station := results[0]
		temp, err := strconv.ParseFloat(results[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing float from '%s': %v\n", results[1], err)
			os.Exit(1)
		}

		if count, ok := tempCount[station]; !ok {
			tempCount[station] = 1
		} else {
			tempCount[station] = count + 1
		}

		if sum, ok := tempSum[station]; !ok {
			tempSum[station] = temp
		} else {
			tempSum[station] = temp + sum
		}

		if min, ok := tempMin[station]; !ok || temp < min {
			tempMin[station] = temp
		}

		if max, ok := tempMax[station]; !ok || temp > max {
			tempMax[station] = temp
		}
	}

	fmt.Print("{")
	items := 0
	for station, count := range tempCount {
		fmt.Printf("%s=%.1f/%.1f/%.1f", station, tempMin[station], tempSum[station]/float64(count), tempMax[station])
		if items != len(tempCount)-1 {
			fmt.Printf(", ")
		}
		items++
	}
	fmt.Print("}\n")
}

var _ Strategy = (*SimpleStrategy)(nil)
