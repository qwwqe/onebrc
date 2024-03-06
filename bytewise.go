package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BytewiseStrategy struct{}

func (s BytewiseStrategy) Process(filename string) {
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

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data); i++ {
			if data[i] == ';' || data[i] == '\n' {
				return i + 1, data[0:i], nil
			}
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner.Split(split)

	station := ""

	for atStationToken := true; scanner.Scan(); atStationToken = !atStationToken {
		token := scanner.Text()

		if atStationToken {
			station = token
			continue
		}

		temp, err := strconv.ParseFloat(token, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing float from '%s': %v\n", token, err)
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

var _ Strategy = (*BytewiseStrategy)(nil)
