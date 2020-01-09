package main

import (
    "bufio"
    "gopher-quiz/quiz"
    "gopher-quiz/services"
    "os"
)

func main() {
    startQuiz()
}

func startQuiz()()  {
    questionService := services.NewQuestionService()
    questionService.WithReader("problems.csv")
    reader := bufio.NewReader(os.Stdin)
    inputReader := services.NewInputService(reader)

    quiz := quiz.NewQuiz(questionService, inputReader)
    quiz.Begin()
}
