package main

import (
	. "dependency/src/github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	Convey("TestRemoveDuplicate", t, func() {
		Convey("no duplicate", func() {
			in := []int{1, 2, 3}
			err := RemoveDuplicate(&in)
			So(err, ShouldBeNil)
			So(in, ShouldResemble, []int{1, 2, 3})
		})
		Convey("int slice", func() {
			in := []int{1, 1, 2, 2, 3}
			err := RemoveDuplicate(&in)
			So(err, ShouldBeNil)
			So(in, ShouldResemble, []int{1, 2, 3})
		})
		Convey("string slice", func() {
			in := []string{"1", "1", "2", "2", "3"}
			err := RemoveDuplicate(&in)
			So(err, ShouldBeNil)
			So(in, ShouldResemble, []string{"1", "2", "3"})
		})
		Convey("float slice", func() {
			in := []float64{1.1, 1.1, 1.111111, 2.2, 2.2}
			err := RemoveDuplicate(&in)
			So(err, ShouldBeNil)
			So(in, ShouldResemble, []float64{1.1, 1.111111, 2.2})
		})
		Convey("slice in slice", func() {
			in := [][]int{{1, 2, 3}, {1, 2, 3}}
			err := RemoveDuplicate(&in)
			So(err, ShouldNotBeNil)

		})
		Convey("map in slice", func() {
			in := []map[string]string{{"1": "1"}, {"1": "1"}}
			err := RemoveDuplicate(&in)
			So(err, ShouldNotBeNil)
		})
		Convey("not pass an address", func() {
			in := []int{1, 1, 2, 3}
			err := RemoveDuplicate(in)
			So(err, ShouldNotBeNil)
		})
		Convey("comparable struct slice", func() {
			type s struct {
				a int
			}
			in := []s{{1}, {1}, {2}, {2}, {3}}
			err := RemoveDuplicate(&in)
			So(err, ShouldBeNil)
			So(in, ShouldResemble, []s{{1}, {2}, {3}})
		})
		Convey("incomparable struct slice", func() {
			type s struct {
				a []int
			}
			in := []s{{[]int{1}}, {[]int{1}}, {[]int{2}}, {[]int{2}}, {[]int{3}}}
			err := RemoveDuplicate(&in)
			So(err, ShouldNotBeNil)
		})
		Convey("array", func() {
			in := [3]int{1, 2, 3}
			err := RemoveDuplicate(&in)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestInArray(t *testing.T) {
	Convey("TestInArray", t, func() {
		Convey("str in", func() {
			arr := []string{"1", "1", "xyz"}
			in, err := InSlice("1", arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeTrue)
		})
		Convey("str not in", func() {
			arr := []string{"1", "1", "xyz"}
			in, err := InSlice("2", arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeFalse)
		})
		Convey("int in", func() {
			arr := []int{1, 2, 1, 1, 4}
			in, err := InSlice(1, arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeTrue)
		})
		Convey("struct in", func() {
			type s struct {
				a []int
			}
			arr := []s{{a: []int{1, 2}}, {a: []int{1, 2, 3, 4}}}
			in, err := InSlice(s{a: []int{1, 2}}, arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeTrue)
		})
		Convey("struct not in", func() {
			type s struct {
				a []int
			}
			arr := []s{{a: []int{1, 2}}, {a: []int{1, 2, 3, 4}}}
			in, err := InSlice(s{a: []int{1, 2, 3}}, arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeFalse)
		})
		Convey("struct in array", func() {
			type s struct {
				a []int
			}
			arr := [2]s{{a: []int{1, 2}}, {a: []int{1, 2, 3, 4}}}
			in, err := InSlice(s{a: []int{1, 2}}, arr)
			So(err, ShouldBeNil)
			So(in, ShouldBeTrue)
		})
		Convey("not array/slice", func() {

			arr := map[int]int{1: 1, 2: 2}
			_, err := InSlice(1, arr)
			So(err, ShouldNotBeNil)
		})
	})
}
