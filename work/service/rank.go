package service

import (
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

	bucket := ts.bucket
	for {
		bucket
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
