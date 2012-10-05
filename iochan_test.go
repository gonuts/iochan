package iochan

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBasic(t *testing.T) {
	b := make(Buffer, 10)
	r := bytes.NewBuffer([]byte("this is  a buffer withlongwords sometimes"))
	for s := range b.ReaderChan(r, " ") {
		fmt.Printf("%s|", s)
	}

	b = make(Buffer, 1024)
	for s := range b.FileLineChan("iochan_test.go") {
		fmt.Print(s)
	}

	b = make(Buffer, 1024)
	data, err := ioutil.ReadFile("iochan_test.go")
	if err != nil {
		t.Errorf(err.Error())
	}
	for s := range b.BytesLineChan(data) {
		fmt.Print(s)
	}
}
