package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("error opening csv file: %s", *csvFileName))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("error reading file")
	}
	problems := parseLines(lines)

	correct := 0
	for i, problem := range problems {
		isCorrect := askQuestion(i+1, problem)
		if isCorrect {
			correct++
			fmt.Println("You are correct!")
			continue
		}
		fmt.Println("You are not correct.")
	}

	fmt.Printf("You answered %d out of %d questions correct\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func askQuestion(num int, question problem) bool {
	fmt.Printf("Problem #%d - %s = ", num, question.q)
	var answer string
	fmt.Scanf("%s\n", &answer)
	return answer == question.a
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
