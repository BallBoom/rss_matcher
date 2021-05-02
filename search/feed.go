package search

import (
	"encoding/json"
	"os"
)

// 用于读取json数据文件

// josn文件路径
const dataFile = "data/data.json"

// feed结构体，映射json数据
// {
// 	"site":"",
// 	"link":"",
// 	"type":""
// }

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 解析json文件 将json转换成对应的结构体指针切片
func RetrieveFeeds() ([]*Feed, error) {
	df, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer df.Close()

	// 定义一个feed指针切片
	var feed []*Feed

	// 构建json 反序列化 解码
	err = json.NewDecoder(df).Decode(&feed)
	return feed, err
}
