package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSpec(t *testing.T) {
	var path string

	Convey("Given some non-existing path", t, func() {

		path = "lala"

		Convey("checkPathExists() should return false", func() {
			So(checkPathExists(path), ShouldEqual, false)
		})
	})

	Convey("Given some existing path", t, func() {

		path = "./_lala"
		os.Mkdir(path, 0700)
		Convey("checkPathExists() should return false", func() {
			So(checkPathExists(path), ShouldEqual, false)
		})
		os.Remove(path)
	})

}
