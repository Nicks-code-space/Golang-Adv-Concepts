package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	Name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.Name))
	return err
}

func main() {
	f1, err := os.Create("hello.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f1.Close()

	_, err = f1.Write([]byte("Hello gophers!"))
	if err != nil {
		log.Fatalf("error %s", err)
	}

	f2, err := os.Create("success.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f2.Close()

	_, err = f2.Write([]byte("You have successfully created two text files and wrote some data to those."))
	if err != nil {
		log.Fatalf("error %s", err)
	}

	p := person{Name: "James"}
	var b bytes.Buffer
	p.writeOut(&b)
	fmt.Println(b.String())
}
