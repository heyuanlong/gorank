package service

import "fmt"

type dataStruct struct {
	userId int
	value  int
}

func NewDataStruct(userId int, value int) *dataStruct {
	return &dataStruct{
		userId: userId,
		value:  value,
	}
}

func (ts *dataStruct) GetKey() int {
	return ts.userId
}
func (ts *dataStruct) Compare(b basedataInterface) bool {
	return ts.GetValue() < b.GetValue()
}
func (ts *dataStruct) GetValue() int {
	return ts.value
}
func (ts *dataStruct) Print() {
	fmt.Print(ts.value, ",")
}
