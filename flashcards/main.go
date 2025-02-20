package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

type Flashcard struct {
	Question     string   `json:"question"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correct"`
}

type ReviewQuestion struct {
	Question string
	Answer   string
}

func main() {
	// Read the JSON file
	file, err := os.Open("chapter1.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Create a slice to store the flashcards
	var flashcards []Flashcard

	// Unmarshal the JSON data into the flashcards slice
	err = json.Unmarshal(bytes, &flashcards)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	perms := rand.Perm(len(flashcards))[:5]
	reviewQuestions := []ReviewQuestion{}
	tally := 0
	for _, i := range perms {
		currCard := flashcards[i]
		correctAnswer := currCard.Options[currCard.CorrectIndex]
		fmt.Printf("Question: %s\n", flashcards[i].Question)
		for j, option := range flashcards[i].Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		iText, err := strconv.Atoi(text[:1])
		if err != nil {
			fmt.Println(err)
			return
		}
		guess := currCard.Options[iText-1]
		if guess != correctAnswer {
			fmt.Printf("----------------NOPE CORRECT ANSWER: %s---------------\n", correctAnswer)
			reviewQuestions = append(reviewQuestions, ReviewQuestion{Question: currCard.Question, Answer: correctAnswer})
		} else {
			fmt.Println("----------------YARP----------------")
			tally++
		}
	}

	fmt.Printf("You got %d/%d correct\nQuestions to review:\n", tally, len(perms))
	for _, question := range reviewQuestions {
		fmt.Printf("-------------------------\n%s - %s\n-------------------------\n", question.Question, question.Answer)
	}
}
