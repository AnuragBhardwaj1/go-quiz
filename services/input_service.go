package services

import (
    "bufio"
)

type inputService struct {
    reader *bufio.Reader
}

func (i inputService) Read(terminator byte) (string, error) {
    //panic("implement me")
    return  i.reader.ReadString(terminator)

}

func NewInputService(reader *bufio.Reader) *inputService {
    return &inputService{
        reader: reader,
    }
}

