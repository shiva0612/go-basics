package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
every writer:
	all type of write methods
	readFrom (only few has)
		-writer.readfrom(reader) //writer calls Read() on reader multiple times until EOF
								//writer <- reader (no buffering)

every reader:
	all type of read methods
	writeTo (only few has)
		-reader.writeto(writer) //writer calls Write() on reader multiple times
								//reader -> writer (no buffering)

strings bytes bufio: writer
	strings.builder
		writebyte, writestring
		get back string, as []byte
			[x-readfrom-x] not present
	bytes.buffer
		writebyte, writestring
		get back string, as []byte
			readfrom(reader)
	bufio.newWriterSize(w *io.writer,n int)
		writerbyte, writestring
		readfrom(reader)
			if w also has readfrom
				bufio.readfrom just calls w.readfrom (content is not buffered)
			else
				bufio.readfrom content is buffered

strings bytes bufio: reader
	strings.newReader("string"):
		all read methods
		writeTo(writer)
	bytes.newReader([]byte("shiva"));
		all read methods
		writeTo(writer)
	bufio.newReaderSize(r *io.reader,n int)
		all read methods
		read([]byte)
		readString(delim)
		readBytes(delim)
		writeTo(writer)
			if r has writeTo
				bufio.writeTo just calls r.writeTo (content is not buffered)
			else
				bufio.writeTo content is buffered
*/

func main() {

	// tee()
	// pipe()
	// multiread()
	// multiwrite()
	// readfrom()
	// writeto()
	// ---
	f, _ := os.Open("file.mov")
	bufr(f)
}

/*
 */
func bufr(f *os.File) {
	dest, _ := os.Open("destination file")
	bfr := bufio.NewReaderSize(f, 10) //intialize reader whic read atleast 10 bytes on each Read() call

	//it will read 3 bytes
	b := make([]byte, 3)
	for {
		n, err := bfr.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(b[:n])
	}

	b = make([]byte, 20)
	for {
		// n, err := bfr.Read(b) -> this way it will only read 10 bytes, but u wanted to 20
		n, err := io.ReadFull(bfr, b) //use this
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(b[:n])
	}

	//does not buffer since f(os.file) has readfrom method so (f -> dest) nobuffering
	bfr.WriteTo(dest)

	//same explanation
	io.Copy(dest, bfr)
}

/*
 */
func bufw(f *os.File) {
	src, _ := os.Open("source.file")
	bufw := bufio.NewWriterSize(f, 10) //intialize writer which buffers exactly 10 bytes and writes to underlying writer only if buffer is exhausted

	// since f(os.file) has readfrom method so (f <- src)
	bufw.ReadFrom(src)

	/*
		this is does not write anything to stdout -> it will be buffered and on exit it will be lost
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(2)
			exit

		content is written to stdout, since flush is called
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(2)
			bufw.flush()
			exit

		directly written to stdout without buffering since 12 > 10
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(12)
			exit

			bufio.newWriterSize(os.stdout, 10)
			bufw.write(10) //will buffer
			bufw.write(2) //prev 10 bytes will be printed and 2 bytes will be buffered
			exit

	*/

}
