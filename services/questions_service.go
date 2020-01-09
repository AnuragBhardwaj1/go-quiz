package services

import (
    "encoding/csv"
    "gopher-quiz/quiz"
    "io"
    "log"
    "os"
)

type questionService struct {
    buffer *csv.Reader
    file   *os.File
}

func (this *questionService) CloseConnection() error {
    return this.file.Close()
}

func NewQuestionService() *questionService {
    return &questionService{}
}

func (this *questionService) Next() *quiz.Question {
    record, err := this.buffer.Read()
    if err != nil {
        if err == io.EOF {
            this.file.Close()
            return nil
        }
        log.Fatal(err)
    }
    return quiz.NewQuestion(record[0], record[1])
}

func (this *questionService) WithReader(fileName string) (*questionService, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return this , err
    }
    this.file = file
    this.buffer = csv.NewReader(file)
    return this, nil
}
