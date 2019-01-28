package mfomo

import (
	"net/http"
	"strconv"

	kinit "gorank/initialize"
	kcode "gorank/work/code"
	kbase "gorank/work/control/base"

	"github.com/gin-gonic/gin"
)

type Rank struct {
}

func NewRank() *Rank {
	return &Rank{}
}
func (ts *Rank) Load() []kbase.RouteWrapStruct {
	m := make([]kbase.RouteWrapStruct, 0)
	m = append(m, kbase.Wrap("GET", "/mfomo/rank1", ts.Rank1, 0))
	m = append(m, kbase.Wrap("GET", "/mfomo/rank2", ts.Rank2, 0))

	return m
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/rank1?rank_id=1&min_coins=0.5&chain_id=1
type subRank2Struct struct {
	RankId   string  `json:"rank_id"`   // 转币地址
	MinCoins float64 `json:"min_coins"` //最少投入eth数量
	ChainId  int     `json:"chain_id"`  //分红占比

}

func (ts *Rank) Rank1(c *gin.Context) {
	callbackName := kbase.GetParam(c, "callback")
	rank_id := kbase.GetParam(c, "rank_id")

	min_coins, err := strconv.ParseFloat(kbase.GetParam(c, "min_coins"), 0)
	if err != nil {
		kinit.LogWarn.Println(min_coins, err)
		kbase.SendErrorJsonStr(c, kcode.PARAM_WRONG, callbackName)
		return
	}
	chain_id, err := strconv.Atoi(kbase.GetParam(c, "chain_id"))
	if err != nil {
		kbase.SendErrorJsonStr(c, kcode.PARAM_WRONG, callbackName)
		return
	}

	subObject := subRank2Struct{
		RankId:   rank_id,
		MinCoins: min_coins,
		ChainId:  chain_id,
	}
	kbase.ReturnDataI(c, kcode.SUCCESS_STATUS, subObject, callbackName)
}

//-----------------------------------------------------------------------------------

// http://127.0.0.1:8080/mfomo/rank2
func (ts *Rank) Rank2(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("Rank2"))
}

//-----------------------------------------------------------------------------------
