package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"test/src"
	"testing"
)

func TestPalindrome(t *testing.T) {
	Convey("test func Palindrome", t, func() {
		So(src.ShouldReturnFalse("detartrated"), ShouldBeFalse)
	})
}
