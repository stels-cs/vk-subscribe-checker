package main

import (
	"testing"
)

func TestReadFile(t *testing.T)  {
	r, err := GetReadet("testIds.txt")
	if err != nil {
		t.Error(err)
	}

	id, err := r.GetNexId()
	for err == nil {
		if id <= 0 {
			t.Error("Expected id > 0, got", id)
		}
		id, err = r.GetNexId()
	}

	if err.Error() != "EOF" {
		t.Error(err)
	}
}