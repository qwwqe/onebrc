package main

type SimpleStrategy struct{}

func (s SimpleStrategy) Process(filename string) []string {
	return []string{}
}

var _ Strategy = (*SimpleStrategy)(nil)
