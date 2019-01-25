package service

import (
	kinits "gorank/initialize"
	"sort"
)

const (
	PRE_POS  = -1
	MID_POS  = 0
	NEXT_POS = 1
)

type datasStruct []basedataInterface

func (s datasStruct) Len() int {
	return len(s)
}
func (s datasStruct) Less(i, j int) bool {
	return s[i].Compare(s[j])
}
func (s datasStruct) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type BucketStruct struct {
	prev  *BucketStruct
	next  *BucketStruct
	nums  int         //当前数量
	maxs  int         //合理最大数量
	caps  int         //最大容量
	datas datasStruct //数据

	headData basedataInterface
	tailData basedataInterface
}

func NewBucketStruct() *BucketStruct {
	return &BucketStruct{
		maxs:  20,
		caps:  30,
		datas: make(datasStruct, 500),
	}
}

func (ts *BucketStruct) CanAdd(bdata basedataInterface) int { //-1:前面 0:中间 1:后面
	if ts.nums == 0 {
		return MID_POS
	}

	if bdata.Compare(ts.headData) == false {
		return PRE_POS
	}
	if bdata.Compare(ts.tailData) == true {
		return MID_POS
	}
	return NEXT_POS
}

func (ts *BucketStruct) Add(bdata basedataInterface) {

	pos := sort.Search(ts.nums, func(i int) bool { return ts.datas[i].Compare(bdata) })
	kinits.LogInfo.Println("pos:", pos, "value:", bdata.GetValue())
	for index := ts.nums; index >= pos; index-- {
		ts.datas[index+1] = ts.datas[index]
	}
	ts.datas[pos] = bdata

	ts.nums += 1
	ts.headData = ts.datas[0]
	ts.tailData = ts.datas[ts.nums-1]

	ts.check()
}

func (ts *BucketStruct) check() {
	if ts.nums < ts.caps {
		return
	}

	pnums := 0
	if ts.prev != nil {
		if ts.prev.nums < ts.prev.maxs { //分给前面的bucket
			pnums = ts.prev.maxs - ts.prev.nums
			ts.prev.addTails(ts.datas[0:pnums])
		}
	}
	//nnums := 0
	if (ts.nums - pnums) > ts.maxs { //强制分给后面的bucket
		//nnums = (ts.nums - pnums) - ts.maxs
		ts.loadNext()
		ts.next.addHeads(ts.datas[ts.maxs+pnums : ts.nums])
	}

	for index := 0; index < ts.maxs; index++ {
		ts.datas[index] = ts.datas[index+pnums]
	}

	ts.nums = ts.maxs
	ts.headData = ts.datas[0]
	ts.tailData = ts.datas[ts.nums-1]
}

func (ts *BucketStruct) addTails(bdatas []basedataInterface) {
	length := len(bdatas)
	i := 0

	for index := ts.nums; index < (ts.nums + length); index++ {
		ts.datas[index] = bdatas[i]
		i++
	}
	ts.nums += length
	ts.tailData = ts.datas[ts.nums-1]
}
func (ts *BucketStruct) addHeads(bdatas []basedataInterface) {
	ts.datas = append(bdatas, ts.datas...) //有待优化

	ts.nums += len(bdatas)
	ts.headData = ts.datas[0]
	ts.tailData = ts.datas[ts.nums-1]

	ts.check()
}

func (ts *BucketStruct) loadNext() {
	if ts.next != nil {
		return
	}

	new := NewBucketStruct()
	new.next = nil
	new.prev = ts
	ts.next = new
}

func (ts *BucketStruct) Print() {
	for index := 0; index < ts.nums; index++ {
		ts.datas[index].Print()
	}
}
