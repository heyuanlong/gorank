package main

import (
	"fmt"
	krank "gorank/rank"
)

type dataStruct struct {
	userId int
	value  int64
}

func NewDataStruct(userId int, value int64) *dataStruct {
	return &dataStruct{
		userId: userId,
		value:  value,
	}
}

// key
func (ts *dataStruct) GetKey() int {
	return ts.userId
}

// 比较 > 或者 <
func (ts *dataStruct) Comparep(b krank.BasedataInterface) bool {
	bobj := b.(*dataStruct)
	return ts.value < bobj.value
}

// 比较 >= 或者 <=
func (ts *dataStruct) Compare(b krank.BasedataInterface) bool {
	return ts.Comparep(b) || ts.Equal(b) //一定得>= 或者 <= ，否则Find将失败
}

// 比较 =
func (ts *dataStruct) Equal(b krank.BasedataInterface) bool {
	bobj := b.(*dataStruct)
	return ts.value == bobj.value
}

// 设置 比较值
func (ts *dataStruct) SetValue(b krank.BasedataInterface) {
	bobj := b.(*dataStruct)
	ts.value = bobj.value
}

//
func (ts *dataStruct) Write() {
	fmt.Println(ts.userId, ":", ts.value)
}
