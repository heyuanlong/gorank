package rank

type BasedataInterface interface {
	GetKey() int                     // key
	Comparep(BasedataInterface) bool // 比较 > 或者 <
	Compare(BasedataInterface) bool  // 比较 >= 或者 <=
	Equal(BasedataInterface) bool    // 比较 =
	SetValue(BasedataInterface)      // 设置 比较值
	Write()                          // 打印
}
