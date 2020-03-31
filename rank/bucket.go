package rank

import (
	"errors"
	"sort"
)

const (
	PRE_POS  = -1
	MID_POS  = 0
	NEXT_POS = 1
)

type datasStruct []BasedataInterface

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
	prev     *BucketStruct
	next     *BucketStruct
	nums     int               //当前数量
	goodNums int               //合理最大数量
	maxs     int               //最大容量
	datas    datasStruct       //数据
	headData BasedataInterface //头部数据
	tailData BasedataInterface //尾部数据
}

func NewBucketStruct() *BucketStruct {
	return &BucketStruct{
		goodNums: 400,
		maxs:     500,
		datas:    make(datasStruct, 501),
	}
}
func (ts *BucketStruct) GetNums() int {
	return ts.nums
}
func (ts *BucketStruct) CanAdd(bdata BasedataInterface, isFind bool) int { //-1:前面 0:中间 1:后面

	if ts.nums == 0 {
		if isFind == false {
			//fmt.Print(MID_POS, "pm    ")
			return MID_POS
		}
		//fmt.Print(NEXT_POS, "pm    ")
		return NEXT_POS
	}

	//fmt.Print(ts.headData.GetValue(), "h ", ts.tailData.GetValue(), "v ")
	if bdata.Comparep(ts.headData) == true {
		//fmt.Print(PRE_POS, "p    ")
		return PRE_POS
	}
	if bdata.Compare(ts.tailData) == true {
		//fmt.Print(MID_POS, "p    ")
		return MID_POS
	}
	//fmt.Print(NEXT_POS, "p    ")
	return NEXT_POS
}

func (ts *BucketStruct) Add(bdata BasedataInterface) {

	pos := sort.Search(ts.nums, func(i int) bool { return bdata.Compare(ts.datas[i]) })
	//fmt.Println("pos:", pos, "value:", bdata.GetValue())
	for index := ts.nums; index >= pos; index-- {
		ts.datas[index+1] = ts.datas[index] //数据向后移动
	}
	ts.datas[pos] = bdata

	ts.nums += 1
	ts.headData = ts.datas[0]
	ts.tailData = ts.datas[ts.nums-1]

	ts.check()
}
func (ts *BucketStruct) Find(bdata BasedataInterface) (int, error) {
	pos := sort.Search(ts.nums, func(i int) bool { return bdata.Compare(ts.datas[i]) })
	for i := pos; i < ts.nums; i++ {
		//fmt.Println("GetKey:", ts.datas[i].GetKey(), "GetValue:", ts.datas[i].GetValue())
		if ts.datas[i].GetKey() == bdata.GetKey() {
			return i, nil
		}
	}
	return 0, errors.New("not find")
}
func (ts *BucketStruct) UpdateInThisBucket(pos int, newd BasedataInterface) error {
	if newd.Comparep(ts.headData) == false && newd.Compare(ts.tailData) == true { //说明新的值也是落在这个bucket里
		ts.datas[pos] = newd
		sort.Sort(ts.datas[:ts.nums])
		ts.headData = ts.datas[0]
		ts.tailData = ts.datas[ts.nums-1]
		return nil
	}

	return errors.New("new data not in this bucket")
}

func (ts *BucketStruct) Del(bdata BasedataInterface) error {
	pos := sort.Search(ts.nums, func(i int) bool { return bdata.Compare(ts.datas[i]) })
	for i := pos; i < ts.nums; i++ {
		////fmt.Println("GetKey:", ts.datas[i].GetKey(), "GetValue:", ts.datas[i].GetValue())
		if ts.datas[i].GetKey() == bdata.GetKey() {
			for j := i; j < ts.nums-1; j++ {
				ts.datas[j] = ts.datas[j+1] //数据向前移动
			}
			ts.nums -= 1
			ts.tailData = ts.datas[ts.nums-1]
			return nil
		}
	}
	return errors.New("not find")
}

func (ts *BucketStruct) check() {
	if ts.nums < ts.maxs { //未到达最大数据
		return
	}
	//fmt.Println("go li---")

	pnums := 0
	if ts.prev != nil {
		if ts.prev.nums < ts.prev.goodNums { //分给前面的bucket
			pnums = ts.prev.goodNums - ts.prev.nums
			ts.prev.addTails(ts.datas[0:pnums])
		}
	}
	if (ts.nums - pnums) > ts.goodNums { //强制分给后面的bucket
		ts.loadNext()
		ts.next.addHeads(ts.datas[ts.goodNums+pnums : ts.nums])
	}

	for index := 0; index < ts.goodNums; index++ {
		ts.datas[index] = ts.datas[index+pnums] //从pnums位置的数据移动前面
	}

	ts.nums = ts.goodNums
	ts.headData = ts.datas[0]
	ts.tailData = ts.datas[ts.nums-1]
}

func (ts *BucketStruct) addTails(bdatas []BasedataInterface) {
	length := len(bdatas)
	i := 0

	for index := ts.nums; index < (ts.nums + length); index++ {
		ts.datas[index] = bdatas[i]
		i++
	}
	ts.nums += length
	ts.tailData = ts.datas[ts.nums-1]
}
func (ts *BucketStruct) addHeads(bdatas []BasedataInterface) {
	tmp := make([]BasedataInterface, len(bdatas)) //有待优化
	copy(tmp, bdatas)                             //刚开始没有用copy，坑了自己一把
	ts.datas = append(tmp, ts.datas...)

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
func (ts *BucketStruct) Datas() (int, datasStruct) {
	return ts.nums, ts.datas[:ts.nums]
}

func (ts *BucketStruct) Write() {
	for index := 0; index < ts.nums; index++ {
		ts.datas[index].Write()
	}
}
