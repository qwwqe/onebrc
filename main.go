package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type StrategyType = string

const (
	SimpleStrategyType StrategyType = "simple"
)

type Strategy interface {
	Process(filename string) []string
}

func Usage() {
	fmt.Println("Usage: onebrc [OPTIONS] <data-file>")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage

	strategyType := flag.String("strategy", SimpleStrategyType, "the strategy to use")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	var strategy Strategy

	switch *strategyType {
	case SimpleStrategyType:
		strategy = SimpleStrategy{}
	default:
		fmt.Fprintf(os.Stderr, "Unknown strategy '%s'\n", *strategyType)
		os.Exit(1)
	}

	startTime := time.Now()
	strategy.Process(filename)
	duration := time.Since(startTime)

	fmt.Printf("Executed '%s' strategy in %v seconds\n", *strategyType, duration)
}
