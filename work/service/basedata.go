package service

type basedataInterface interface {
	GetKey() int
	compare(basedataInterface) bool
}
