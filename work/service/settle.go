package service

import (
	"fmt"
	krank "gorank/work/rank"
	"math/rand"
	"time"
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
func (ts *SettleStruct) Run2(nums int) {
	start := time.Now()
	rank := krank.NewRankStruct()
	for index := 0; index < 20*10000; index++ {
		d := krank.NewDataStruct(index, rand.Intn(9999*10000))
		rank.Add(d)
	}
	fmt.Println("add cost=", time.Since(start))
	fmt.Println()

	start = time.Now()
	d := krank.NewDataStruct(8784, 0)
	fmt.Println(rank.GetRank(d))
	fmt.Println("get rank cost=", time.Since(start))
	fmt.Println()

	start = time.Now()
	d = krank.NewDataStruct(8784, 788999)
	fmt.Println(rank.Update(d))
	fmt.Println("update cost=", time.Since(start))
	fmt.Println()

	start = time.Now()
	d = krank.NewDataStruct(8784, 0)
	fmt.Println(rank.GetRank(d))
	fmt.Println("get rank cost=", time.Since(start))
	fmt.Println()

	start = time.Now()
	d = krank.NewDataStruct(8784, 0)
	fmt.Println(rank.Del(d))
	fmt.Println("get Del cost=", time.Since(start))
	fmt.Println()

	start = time.Now()
	objs2 := rank.GetPage(2000, 15)
	for _, v := range objs2 {
		fmt.Print(v.GetValue(), ",")
	}
	fmt.Println("page cost=", time.Since(start))
	fmt.Println()
	fmt.Println()
}
