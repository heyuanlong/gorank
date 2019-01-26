package service

import (
	"fmt"
)

type SettleStruct struct {
	tm int64
}

func NewSettleStruct() *SettleStruct {
	return &SettleStruct{}
}

func (ts *SettleStruct) Init() {

}
func (ts *SettleStruct) Run(nums int) {
	rank := NewRankStruct()
	for index := 0; index < nums; index++ {
		//d := NewDataStruct(index, rand.Intn(999))
		d := NewDataStruct(index, index)
		rank.Add(d)

	}
	for index := 0; index < nums; index++ {
		//d := NewDataStruct(index, rand.Intn(999))
		d := NewDataStruct(index+100000, index)
		rank.Add(d)

	}

	rank.LookAll()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(rank.Update(NewDataStruct(60, 59)))

	rank.LookAll()
	fmt.Println()

	objs := rank.GetPage(1, 15)
	for _, v := range objs {
		fmt.Print(v.GetValue(), ",")
	}
	fmt.Println()
	objs2 := rank.GetPage(2, 15)
	for _, v := range objs2 {
		fmt.Print(v.GetValue(), ",")
	}
	fmt.Println()
	fmt.Println()
}
