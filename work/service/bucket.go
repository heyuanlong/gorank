package service

type BucketStruct struct {
	prev  *BucketStruct
	next  *BucketStruct
	nums  int                 //当前数量
	max   int                 //合理最大数量
	caps  int                 //最大容量
	datas []basedataInterface //数据

	headData basedataInterface
	tailData basedataInterface
}

func NewBucketStruct() *BucketStruct {
	return &BucketStruct{
		max:   400,
		caps:  500,
		datas: make([]basedataInterface, 0, 500),
	}
}

func (ts *BucketStruct) CanAdd(bdata basedataInterface) int { //-1:前面 0:中间 1:后面
	if ts.nums == 0 {
		return -1
	}

	if bdata.compare(ts.headData) == false {
		return -1
	}
	if bdata.compare(ts.tailData) == true {
		return 0
	}
	return 1
}

func (ts *BucketStruct) Run() {

}
