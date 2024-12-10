package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetLines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not read input.txt")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not parse number %s", s)
	}
	return n
}
