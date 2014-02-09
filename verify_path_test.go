package main

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"regexp"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("#checkPathExists", t, func() {
		var path string
		cs := make(chan string)
		Convey("given some non-existing path", func() {
			path = "lala"
			Convey("checkPathExists() should write missing message to channel", func() {
				go checkPathExists(cs, path)
				So(<-cs, ShouldEqual, "Path lala does not exists!")
			})
		})

		Convey("given some existing path", func() {
			path = "./_lala"
			os.Mkdir(path, 0700)
			Convey("checkPathExists() should return false", func() {
				go checkPathExists(cs, path)
				So(<-cs, ShouldEqual, "")
				os.Remove(path)
			})
		})
	})
	Convey("#readMatchingLines", t, func() {
		input, _ := os.Open("./test_input.txt")
		regex, _ := regexp.Compile(`\/var\/www.*?\s`)
		Convey("readMatchingLines() should return the correct number of lines mathing the regexp", func() {
			lines := readMatchingLines(input, regex)
			So(len(lines), ShouldEqual, 2)
		})
	})
	Convey("#main", t, func() {
		oldStdIn := os.Stdin
		oldStdOut := os.Stdout
		input, _ := os.Open("./test_input.txt")
		r, w, _ := os.Pipe()
		var buf bytes.Buffer
		outC := make(chan string)
		Convey(" should write two lines (128 chars) to stdout", func() {
			os.Stdin = input
			os.Stdout = w
			go func() {
				io.Copy(&buf, r)
				outC <- buf.String()
			}()
			main()
			w.Close()
			lala := <-outC
			os.Stdin = oldStdIn
			os.Stdout = oldStdOut
			So(len(lala), ShouldEqual, 128)
		})
	})
}
