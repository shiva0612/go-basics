package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	simple()

	multiReader1()
	multiReader2()

	multiwriter()
}

func multiwriter() {
	byte_builder := new(bytes.Buffer)
	string_builder := new(strings.Builder)

	mw := io.MultiWriter(os.Stdout, byte_builder, string_builder)
	_, err := fmt.Fprintln(mw, "hi")
	if err != nil {
		panic(err)
	}
	fmt.Println(byte_builder.String())
	fmt.Println(string_builder.String())

}

func multiReader1() {
	head := strings.NewReader("<msg>")
	body := strings.NewReader("hi shiva")
	foot := strings.NewReader("</msg>")

	for _, r := range []io.Reader{head, body, foot} {
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func multiReader2() {
	head := strings.NewReader("<msg>")
	body := strings.NewReader("hi shiva")
	foot := strings.NewReader("</msg>")

	mr := io.MultiReader(head, body, foot)
	_, err := io.Copy(os.Stdout, mr)
	if err != nil {
		log.Println(err.Error())
	}
}

func simple() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprintf(pw, "Welcome")
		if err != nil {
			panic("error while writing: " + err.Error())
		}
	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic("error while reading: " + err.Error())
	}
}
