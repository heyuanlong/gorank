package rank

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
func (ts *dataStruct) Comparep(b basedataInterface) bool {
	return ts.GetValue() < b.GetValue() //一定得>= 或者 <= ，否则Find将失败
}
func (ts *dataStruct) Compare(b basedataInterface) bool {
	return ts.Comparep(b) || ts.GetValue() == b.GetValue() //一定得>= 或者 <= ，否则Find将失败
}
func (ts *dataStruct) Equal(b basedataInterface) bool {
	return ts.GetValue() == b.GetValue() //
}

func (ts *dataStruct) GetValue() int {
	return ts.value
}
func (ts *dataStruct) SetValue(v int) {
	ts.value = v
}
 
func (ts *dataStruct) Print() {
	fmt.Println(ts.userId, ":", ts.value)
}
