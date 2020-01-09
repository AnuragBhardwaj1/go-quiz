package quiz

import (
    "fmt"
    "strings"
)

type quiz struct {
    questionService QuestionService
    inputReader     ReaderService
    questions       []*Question
    score           int
}

func (q *quiz) Begin() {
    var currentQuestion *Question
    for {
        currentQuestion = q.questionService.Next()
        if currentQuestion == nil {
            q.questionService.CloseConnection()
            break
        }
        q.askQuestions(currentQuestion)
        submittedAnswer := q.promptForAnswer()
        if q.isInterrupted(submittedAnswer) {
            break
        }
        q.processSubmittedAnswer(currentQuestion, submittedAnswer)

    }
    q.evaluateAnswers()
    q.displayResults()

}

func (q *quiz) askQuestions(currentQuestion *Question) {
    q.addQuestion(currentQuestion)
    problemStatement := fmt.Sprintf("what is %s ?", currentQuestion.problemStatement)
    fmt.Println(problemStatement)
}

func (q *quiz) addQuestion(currentQuestion *Question) {
    q.questions = append(q.questions, currentQuestion)
}

func (q *quiz) evaluateAnswers() {
    for _, question := range q.questions {
        if question.submittedAnswer == question.correctAnswer {
            q.incrementScore()
        }
    }
}

func (q *quiz) incrementScore() {
    q.score++
}

func (q *quiz) displayResults() {
    fmt.Println("your score is:", q.score)
}

func (q *quiz) isInterrupted(answer string) bool {
    return answer == "quit" || answer == "" || answer == "exit" || answer == "q"
}

func (q *quiz) promptForAnswer() string {
    //Todo: move terminator from here to reader
    submittedAnswer, err := q.inputReader.Read()
    if err != nil {
        return ""
    }
    submittedAnswer = strings.TrimSpace(submittedAnswer)
    return submittedAnswer
}

func (q *quiz) processSubmittedAnswer(currentQuestion *Question, submittedAnswer string) {
    currentQuestion.submittedAnswer = submittedAnswer
    if currentQuestion.submittedAnswer != "" && currentQuestion.correctAnswer != "" {
        currentQuestion.answered = true
    }
}

func NewQuiz(questionService QuestionService, inputReader ReaderService) *quiz {
    return &quiz{
        questionService: questionService,
        inputReader:     inputReader,
        questions:       []*Question{},
        score:           0,
    }
}
