package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// what ever u write x u can read on y (but not other way around)
func pipe() {

	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("from pipe"))
		pw.Close()
	}()
	io.Copy(os.Stdout, pr)

}

// what ever u read from x write to y
func tee() {
	var r io.Reader = strings.NewReader("from tee")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}
}

func multiread() {

	f, _ := os.Open("")

	header := strings.NewReader("<msg>")
	body_file, _ := os.Open("api_response.file")
	footer := strings.NewReader("</msg>")

	mr := io.MultiReader(header, body_file, footer)
	io.Copy(f, mr)

}
func multiwrite() {
	f1, _ := os.Create("file1")
	f2, _ := os.Create("file2")
	f3, _ := os.Create("file3")

	mw := io.MultiWriter(f1, f2, f3)

	fmt.Fprintln(mw, "line1")
	fmt.Fprintln(mw, "line2")

}

// use read from for smaller content so that no for{buffering} is needed
// interface itself says read total till EOF
// same as io.Copy(dst_writer,dst_reader)
func readfrom() {

	r1, _ := os.Open("small file")
	r2 := strings.NewReader("small string whose length is not so big")
	r3 := bytes.NewBuffer([]byte("some api response not so big"))

	w1, _ := os.Open("destination file")

	w1.ReadFrom(io.MultiReader(r1, r2, r3))
	io.Copy(w1, io.MultiReader(r1, r2, r3))

}

func writeto() {
	destF, _ := os.Open("dest.file")
	sb := new(strings.Builder)
	byb := bytes.NewBuffer([]byte{})

	content := "copy this to all the above writers"
	strings.NewReader(content).WriteTo(io.MultiWriter(destF, sb, byb))
}
