package service

type basedataInterface interface {
	GetKey() int
	Compare(basedataInterface) bool
	Print()
	GetValue() int
}
