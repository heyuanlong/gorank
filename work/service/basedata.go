package service

type basedataInterface interface {
	GetKey() int
	Comparep(basedataInterface) bool
	Compare(basedataInterface) bool
	Print()
	GetValue() int
}
