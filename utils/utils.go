package utils

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"

	"golune/conf"
	"time"
)

//ID转换成string
func Id2Str(id bson.ObjectId) string {
	idstr := id.Hex()
	return idstr
}

//把模版解析成字符串
func ParseTmplateToStr(tplname string) string {
	b, err := ioutil.ReadFile(tplname)
	if err != nil {
		fmt.Println(err)
	}
	s := string(b)
	return s
}

//把英文分类替换成中文
func ReplaceCate(cate string) string {
	catename := conf.Cate[cate]
	return catename
}

//把时间戳转换成时间
func UnixToStr(unix int64) string {
	return time.Unix(unix, 0).Format("2006-01-02 15:04:05")
}
