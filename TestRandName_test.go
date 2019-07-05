package main

import (
	"../interview/randName"
	"errors"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRandName(t *testing.T)  {
	Convey("TestGetRandName", t, func() {
		m := make(map[string]bool)
		Convey("生成随机字符串，重复，即返回一个错误", func() {
			So(func() error{
				for i:= 0; i < 1000000000; i ++{
					x:= randName.GetRandName(13)
					if _, ok := m[x]; ok{
						return  errors.New("error :repeated string")
					}else {
						m[x] = true
					}
				}
				return nil
			}(), ShouldBeNil)
		})
		Convey("确认生成的字符串是否是13位", func() {
			So(func() error{
				for k, _ := range m {
					if len(k) != 13 {
						return  errors.New("error :string len not is 13")
					}
				}
				return nil
			}(), ShouldBeNil)
		})
	})
}
