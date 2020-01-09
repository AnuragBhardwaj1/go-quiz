package services

import (
    "bufio"
)

const Terminator =  '\n'

type inputService struct {
    reader *bufio.Reader
}

func (i inputService) Read() (string, error) {
    return  i.reader.ReadString(Terminator)

}

func NewInputService(reader *bufio.Reader) *inputService {
    return &inputService{
        reader: reader,
    }
}

