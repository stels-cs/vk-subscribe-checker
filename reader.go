package main

import (
	"os"
	"strconv"
	"strings"
	"errors"
)

type Reader struct {
	file *os.File
	prefix *[]byte
	buffer []byte
	ids *[]int
}

func GetReadet(fileName string) (*Reader, error) {

	input, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	r := &Reader{
		input,
		nil,
		make([]byte, 4*1024),
		nil,
	}

	return r, nil
}

func (r *Reader) GetNexId() (int, error) {
	if r.ids == nil || len(*r.ids) == 0 {
		var err error
		r.ids, err = r.readPart()
		if err != nil {
			return 0, err
		}
	}
	if len(*r.ids) == 0 {
		return 0, errors.New("EOF")
	}
	id := (*r.ids)[0]
	*r.ids = (*r.ids)[1:]
	return id, nil
}

func (r *Reader) readPart() (*[]int, error) {
	_, err := r.file.Read(r.buffer)
	if err != nil {
		return nil,err
	}
	startId := -1
	endId := -1

	ids := make([]int, 0, 100)

	for i := 0; i < len(r.buffer); i++ {
		x := r.buffer[i]

		if x >= '0' && x <= '9' {
			if startId == -1 {
				startId = i
				endId = i
			} else {
				endId = i
			}
		} else {
			if endId != -1 {
				data := r.buffer[startId:endId+1]
				if r.prefix != nil {
					data = append(*r.prefix, data...)
				}
				id, err := strconv.Atoi(strings.TrimSpace(string(data)))
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
				r.prefix = nil
			} else if r.prefix != nil {
				id, err := strconv.Atoi(strings.TrimSpace(string(*r.prefix)))
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
				r.prefix = nil
			}
			startId = -1
			endId = -1
		}
	}

	if startId != -1 {
		b := make([]byte, 0, 20)
		b = append(b, r.buffer[startId:endId+1]...)
		r.prefix = &b
	}
	return &ids, nil
}
