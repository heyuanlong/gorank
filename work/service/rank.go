package service

import (
	"fmt"
	"sync"
)

type RankStruct struct {
	tm     int64
	rwLock sync.RWMutex
	bucket *BucketStruct
	datam  map[int]basedataInterface
}

func NewRankStruct() *RankStruct {
	r := &RankStruct{
		datam:  make(map[int]basedataInterface),
		bucket: NewBucketStruct(),
	}
	r.bucket.prev = nil
	r.bucket.next = nil

	return r
}

func (ts *RankStruct) Add(bdata basedataInterface) error {
	if _, ok := ts.datam[bdata.GetKey()]; ok {
		//更新
		return ts.update(bdata)
	}

	//添加
	return ts.add(bdata)
}
func (ts *RankStruct) add(bdata basedataInterface) error {
	ts.datam[bdata.GetKey()] = bdata

	preBucket := ts.bucket
	bucket := ts.bucket
	for {

		pos := bucket.CanAdd(bdata)
		if pos == PRE_POS {
			preBucket.Add(bdata)
			break
		}
		if pos == MID_POS {
			bucket.Add(bdata)
			break
		}
		preBucket = bucket
		bucket = bucket.next

		if bucket == nil {
			preBucket.Add(bdata)
			break
		}
	}

	return nil
}
func (ts *RankStruct) update(bdata basedataInterface) error {
	return nil
}

func (ts *RankStruct) Del(basedataInterface) error {
	return nil
}
func (ts *RankStruct) GetRank(basedataInterface) (int, error) {

	return 0, nil
}
func (ts *RankStruct) GetPage(page int, paseSize int) []basedataInterface {

	return []basedataInterface{}
}

func (ts *RankStruct) LookAll() {

	bucket := ts.bucket
	for {
		fmt.Println("-----------------------------")
		bucket.Print()
		bucket = bucket.next

		if bucket == nil {
			break
		}
	}
}
