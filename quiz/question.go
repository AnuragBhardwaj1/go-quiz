package quiz

type Question struct {
 problemStatement string
 correctAnswer    string
 submittedAnswer string
 answered bool
}

func NewQuestion(problemStatement string, correctAnswer string) *Question {
    return &Question{problemStatement: problemStatement, correctAnswer: correctAnswer}
}

