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

	fmt.Println(rank.GetRank(NewDataStruct(60, 0)))

	for index := 0; index < nums; index++ {
		//d := NewDataStruct(index, rand.Intn(999))
		d := NewDataStruct(index, index)
		rank.Del(d)

	}
	fmt.Println()
	rank.LookAll()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
