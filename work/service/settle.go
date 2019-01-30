package service

import (
	"fmt"
	krank "gorank/work/rank"
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
	rank := krank.NewRankStruct()
	for index := 0; index < nums; index++ {
		//d := NewDataStruct(index, rand.Intn(999))
		d := krank.NewDataStruct(index, index)
		rank.Add(d)

	}
	for index := 0; index < nums; index++ {
		//d := NewDataStruct(index, rand.Intn(999))
		d := krank.NewDataStruct(index+100000, index)
		rank.Add(d)

	}

	rank.LookAll()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(rank.Update(krank.NewDataStruct(60, 59)))

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
