package main

import (
	. "github.com/agiledragon/gomonkey"
	"github.com/liuhongdi/unittest02/model"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)



//测试double,为add函数打桩1
func TestDoubleRight(t *testing.T) {
	patch := ApplyFunc(Add, func(a,b int) int {
		return a * 2
	})
	defer patch.Reset()
	//fmt.Println(GetDouble(2))
	Convey("test 2 x 2", t, func() {
		So(GetDouble(2), ShouldEqual,4)
	})
}

//add函数的桩子
func addstub(a,b int) int {
	return a*3
}

//测试double,为add函数打桩2
func TestDoubleError(t *testing.T) {
	patch := ApplyFunc(Add, addstub)
	defer patch.Reset()
	//fmt.Println(GetDouble(2))
	Convey("test 2 x 2", t, func() {
		So(GetDouble(2), ShouldEqual,4)
	})
}

//测试给方法打桩1,返回正确
func TestMethodRight(t *testing.T) {
	var temp *model.MyUser
	patch := ApplyMethod(reflect.TypeOf(temp), "GetUserName", func(_ *model.MyUser) string {
		return "hello,world!"
	})
	defer patch.Reset()
	Convey("GetUserName将返回:hello,world!", t, func() {
		var user *model.MyUser
		user = new(model.MyUser)
		So(user.GetUserName(), ShouldEqual, "hello,world!")
	})
}

//测试给方法打桩2,返回错误
func TestMethodError(t *testing.T) {
	var temp *model.MyUser
	patch := ApplyMethod(reflect.TypeOf(temp), "GetUserName", func(_ *model.MyUser) string {
		return "hello,老刘!"
	})
	defer patch.Reset()
	Convey("GetUserName将返回:hello,world!", t, func() {
		var user *model.MyUser
		user = new(model.MyUser)
		So(user.GetUserName(), ShouldEqual, "hello,world!")
	})
}
