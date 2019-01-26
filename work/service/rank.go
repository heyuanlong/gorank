package service

import (
	"errors"
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

func (ts *RankStruct) Del(d basedataInterface) error {
	bdata, ok := ts.datam[d.GetKey()]
	if ok == false {
		return errors.New("not this key")
	}
	fmt.Println(bdata.GetValue())

	bucket := ts.bucket
	var startBucket *BucketStruct = nil //从这个bucket开始查找
	for {
		pos := bucket.CanAdd(bdata, true)
		if pos == MID_POS { //一定在某个bucket范围内
			startBucket = bucket
			break
		}
		bucket = bucket.next

		if bucket == nil {
			break
		}
	}

	for {
		if startBucket == nil {
			break
		}
		err := startBucket.Del(bdata)
		if err == nil {
			return nil
		}
		startBucket = startBucket.next
	}

	return errors.New("not this key!")

}
func (ts *RankStruct) GetRank(d basedataInterface) (int, error) {
	bdata, ok := ts.datam[d.GetKey()]
	if ok == false {
		return 0, errors.New("not this key")
	}
	fmt.Println(bdata.GetValue())

	bucket := ts.bucket
	var startBucket *BucketStruct = nil //从这个bucket开始查找
	var rankNums int = 0
	for {
		pos := bucket.CanAdd(bdata, true)
		if pos == MID_POS { //一定在某个bucket范围内
			startBucket = bucket
			break
		}
		rankNums += bucket.GetNums()
		bucket = bucket.next

		if bucket == nil {
			break
		}
	}

	for {
		if startBucket == nil {
			break
		}
		pos, err := startBucket.Find(bdata)
		if err == nil {
			rankNums += pos + 1
			return rankNums, nil
		}
		rankNums += startBucket.GetNums()
		startBucket = startBucket.next
	}

	return 0, errors.New("not this key!")
}
func (ts *RankStruct) GetPage(page int, paseSize int) []basedataInterface {

	return []basedataInterface{}
}

func (ts *RankStruct) LookAll() {

	bucket := ts.bucket
	for {

		bucket.Print()
		fmt.Print(" -> ")
		bucket = bucket.next

		if bucket == nil {
			break
		}
	}
}

//-------------------------------------------------------------------------------
func (ts *RankStruct) add(bdata basedataInterface) error {
	ts.datam[bdata.GetKey()] = bdata

	preBucket := ts.bucket
	bucket := ts.bucket
	for {

		pos := bucket.CanAdd(bdata, false)
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
