package rank

import (
	"errors"
	"sync"
)

type RankStruct struct {
	tm     int64
	rwLock sync.RWMutex
	bucket *BucketStruct             //桶链表
	datam  map[int]BasedataInterface //数据map
}

func NewRankStruct() *RankStruct {
	r := &RankStruct{
		datam:  make(map[int]BasedataInterface),
		bucket: NewBucketStruct(),
	}
	r.bucket.prev = nil
	r.bucket.next = nil

	return r
}

//添加元素
func (ts *RankStruct) Add(bdata BasedataInterface) error {
	return ts.Update(bdata)
}

//更新元素
func (ts *RankStruct) Update(bdata BasedataInterface) error {
	if old, ok := ts.datam[bdata.GetKey()]; ok {
		//更新
		return ts.update(old, bdata)
	}

	//添加
	return ts.add(bdata)
}

//删除元素
func (ts *RankStruct) Del(d BasedataInterface) error {
	bdata, ok := ts.datam[d.GetKey()]
	if ok == false {
		return errors.New("not this key")
	}
	//fmt.Println(bdata.GetValue())

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

	for { //因为连续的bucket里面的数据可能是一样的，所以要继续循环查找
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

//获得元素排名
func (ts *RankStruct) GetRank(d BasedataInterface) (int, error) {
	bdata, ok := ts.datam[d.GetKey()]
	if ok == false {
		return 0, errors.New("not this key")
	}
	//fmt.Println(bdata.GetValue())

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

	for { //因为连续的bucket里面的数据可能是一样的，所以要继续循环查找
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

//获得某段排名列表
func (ts *RankStruct) GetPage(page int, paseSize int) []BasedataInterface {
	obj := make([]BasedataInterface, 0, paseSize)
	if page < 1 {
		page = 1
	}
	if paseSize < 5 {
		paseSize = 5
	}

	bucket := ts.bucket
	startNums := (page - 1) * paseSize
	endNums := startNums + paseSize
	orders := 0
	for {

		n, datas := bucket.Datas()
		if (orders + n) >= startNums { //这个bucket的尾部大于startNums的位置
			for i := 0; i < n; i++ {
				if orders >= startNums {
					obj = append(obj, datas[i])
				}

				orders += 1
				if orders == endNums {
					return obj
				}
			}
		} else {
			orders += n
		}

		bucket = bucket.next
		if bucket == nil {
			break
		}
	}
	return obj
}

//打印列表
func (ts *RankStruct) LookAll() {

	bucket := ts.bucket
	for {

		bucket.Write()
		bucket = bucket.next

		if bucket == nil {
			break
		}
	}
}

//-------------------------------------------------------------------------------
func (ts *RankStruct) add(bdata BasedataInterface) error {
	ts.datam[bdata.GetKey()] = bdata

	preBucket := ts.bucket
	bucket := ts.bucket
	for {

		pos := bucket.CanAdd(bdata, false)
		if pos == PRE_POS {
			preBucket.Add(bdata) //落在这个bucket之前
			break
		}
		if pos == MID_POS {
			bucket.Add(bdata) //落在这个bucket之内
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
func (ts *RankStruct) update(old, newd BasedataInterface) error {
	if old.Equal(newd) == true {
		return nil
	}

	bucket := ts.bucket
	var startBucket *BucketStruct = nil //从这个bucket开始查找
	for {
		pos := bucket.CanAdd(old, true)
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
		pos, err := startBucket.Find(old) //找到确定的bucket
		if err == nil {
			err := startBucket.UpdateInThisBucket(pos, newd) //新的值是否也是落在这个bucket里
			if err != nil {                                  //新的值没有落在原来的bucket里
				startBucket.Del(old) //删除，这个是代码确定成功
				return ts.add(newd)
			} else {
				ts.datam[old.GetKey()].SetValue(newd)
			}
			return nil
		}
		startBucket = startBucket.next
	}

	return errors.New("not find old key!")
}
