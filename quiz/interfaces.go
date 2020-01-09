package quiz

type QuestionService interface {
    Next() *Question
    CloseConnection() error
}

type ReaderService interface {
    Read() (string, error)
}
