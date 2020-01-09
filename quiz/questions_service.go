package quiz

type QuestionService interface {
    Next() *Question
}

type ReaderService interface {
    Read(terminater byte) (string, error)
}
