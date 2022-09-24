package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	fileName = "problems.csv"
	introMsg = "Welcome do Gopher Quiz v1! In this first version, the questions of the gopher-quiz are" +
		"preset in the problems.csv file in the root of this repository, and this test is not timed." +
		"Press enter to start."
)

func main() {
	csvReader, err := csvReaderFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	stdinReader := bufio.NewReader(os.Stdin)
	err = runIntro(stdinReader, introMsg)
	if err != nil {
		log.Fatal(err)
	}
	correct, total, err := runQuestions(csvReader, stdinReader)
	if err != nil {
		log.Fatal(err)
	}
	score := calculateScore(correct, total)
	fmt.Printf("You got %d out of %d, which means %.2f%% right", correct, total, score)
}

func runIntro(r *bufio.Reader, introMsg string) error {
	fmt.Println(introMsg)
	_, err := r.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading answer: %w", err)
	}

	return nil
}

func csvReaderFromFile(fileName string) (*csv.Reader, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", fileName, err)
	}

	return csv.NewReader(bytes.NewReader(data)), nil
}

func runQuestions(csvReader *csv.Reader, stdinReader *bufio.Reader) (int, int, error) {
	var counter int
	var correctAnswers int
	for {
		record, err := nextRecord(csvReader)
		if record == nil {
			break
		}
		if err != nil {
			return 0, 0, err
		}
		question := record[0]
		answerKey := record[1]
		counter++

		userAnswer, err := askQuestion(stdinReader, counter, question)
		if err != nil {
			return 0, 0, err
		}

		if isCorrect(userAnswer, answerKey) {
			fmt.Println("Great answer!")
			correctAnswers++

			continue
		}
		fmt.Println("Dumb answer!")
	}

	return correctAnswers, counter, nil
}

func calculateScore(correct, total int) float64 {
	return float64(correct) / float64(total) * 100
}

func isCorrect(userAnswer, answerKey string) bool {
	return strings.Replace(userAnswer, "\n", "", -1) == answerKey
}

func nextRecord(r *csv.Reader) ([]string, error) {
	record, err := r.Read()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading csv: %w", err)
	}
	if len(record) != 2 {
		return nil, fmt.Errorf("invalid line in csv file. each line must contain 2 fields")
	}

	return record, nil
}

func askQuestion(r *bufio.Reader, number int, question string) (string, error) {
	fmt.Printf("Question %d\n%s:\n", number, question)
	userAnswer, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading answer: %w", err)
	}

	return userAnswer, nil
}
