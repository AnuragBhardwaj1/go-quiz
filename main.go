package main

import (
	"bufio"
	"fmt"
	"gopher-quiz/quiz"
	"gopher-quiz/services"
	"os"
)

func main() {
	startQuiz()
}

const FileName = "problems.csv"

func startQuiz() {
	questionService, err := services.NewQuestionService().WithReader(FileName)

	if err != nil {
		fmt.Println("error in loading file")
		fmt.Println("--- exiting ---")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	inputReader := services.NewInputService(reader)

	quiz := quiz.NewQuiz(questionService, inputReader)
	quiz.Begin()
}
