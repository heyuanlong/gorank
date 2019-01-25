package base

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"

	kinit "goapi2/initialize"
	kcode "goapi2/work/code"
)

type DataIStruct struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

func init() {

}

func GetParam(c *gin.Context, key string) string {
	v := c.Query(key)
	if v == "" {
		v = c.PostForm(key)
	}
	if v == "" {
		v = c.Param(key)
	}
	return v
}

func ReturnDataI(c *gin.Context, status int, v interface{}, callbackName string) {
	object := DataIStruct{
		Status: status,
		Info:   kcode.GetCodeMsg(status),
		Data:   v,
	}
	ReturnData(c, object, callbackName)
}
func ReturnData(c *gin.Context, v interface{}, callbackName string) {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		kinit.LogError.Println(err)
	}
	if callbackName == "" {
		c.Data(http.StatusOK, "text/plain", jsonStr)
	} else {
		res := []byte(callbackName)
		res = append(res, []byte("(")...)
		res = append(res, jsonStr...)
		res = append(res, []byte(");")...)
		c.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
}

func SendErrorJsonStr(c *gin.Context, code int, callbackName string) {
	jsonStr := GetErrorJsonStr(code)
	if callbackName == "" {
		c.Data(http.StatusOK, "text/plain", jsonStr)
	} else {
		res := []byte(callbackName)
		res = append(res, []byte("(")...)
		res = append(res, jsonStr...)
		res = append(res, []byte(");")...)
		c.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
}

func GetErrorJsonStr(code int) []byte {
	msg := kcode.GetCodeMsg(code)
	chn_msg := kcode.GetCodeChnMsg(code)

	jsonObj := gabs.New()
	jsonObj.Set(code, "status")
	jsonObj.Set(msg, "info")
	jsonObj.Set(chn_msg, "chn_info")
	return jsonObj.Bytes()
}

func UrlPostGetJsonString(url string, paramMap map[string]interface{}, path string) (string, error) {
	param, err := json.Marshal(paramMap)
	if err != nil {
		return "", err
	}
	kinit.LogWarn.Println(string(param))
	request := gorequest.New().Timeout(5 * time.Second)
	_, body, errs := request.Post(url).Type("multipart").Send(string(param)).End()
	if errs != nil {
		kinit.LogWarn.Println("request.Post fail:", errs)
		return "", errs[0]
	}
	kinit.LogInfo.Println(string(body))
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		return "", err
	}

	v, ok := jsonParsed.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}

	return v, nil
}

func GetSignWithSha256(key, secret string, param map[string]interface{}) string {
	paramM := make(map[string]string)
	for k, v := range param {
		switch val := v.(type) {
		case string:
			paramM[k] = val
		case bool:
			paramM[k] = strconv.FormatBool(val)
		case int:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int32:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int64:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case float32:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		case float64:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		default:
			kinit.LogWarn.Println(k, v)
		}
	}
	keys := make([]string, 0, len(paramM))
	for k := range paramM {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	for _, k := range keys {
		fmt.Fprintf(&b, "%s=%s&", k, paramM[k])
	}

	fmt.Fprintf(&b, "key=%s", key)
	paramStr := b.Bytes()
	kinit.LogWarn.Println(string(paramStr))

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(paramStr)
	//bs := mac.Sum(nil)
	sha := hex.EncodeToString(mac.Sum(nil))
	sha = strings.ToUpper(sha)
	kinit.LogWarn.Println(sha)

	return sha
}
