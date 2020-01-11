package quiz

import (
    "github.com/magiconair/properties/assert"
    "testing"
)

func Test_Question(t *testing.T){
    problemStatement, correctAnswer := "2+2", "4"
    question := NewQuestion(problemStatement, correctAnswer)
    assert.Equal(t, question.problemStatement, problemStatement)
    assert.Equal(t, question.correctAnswer, correctAnswer)
}
