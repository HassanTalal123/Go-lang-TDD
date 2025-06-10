package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	Convey("When calling Hello", t, func() {
		got := Hello()
		want := "Hello, world"

		Convey("It should return 'Hello, world'", func() {
			So(got, ShouldEqual, want)
		})
	})
}
