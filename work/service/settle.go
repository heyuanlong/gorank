package service

import (
	"fmt"
	"math/rand"
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
		d := NewDataStruct(index, rand.Intn(999))
		rank.Add(d)

	}
	rank.LookAll()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
