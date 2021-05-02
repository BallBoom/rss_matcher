package search

import (
	"fmt"
	"log"
)

// 用于支持不同匹配器的接口

// 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// 新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(match Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 调用传进去的匹配器的search方法，返回搜索结果
	searResults, err := match.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将搜索的结果 写入到只写result通道中
	for _, result := range searResults {
		results <- result
	}
}


// 显示函数，将搜索的结果显示读取输出
func Display(results chan *Result){
	// 通道会一直阻塞，直到有结果写入
	// 直到通道被关闭，for循环才会结束
	for result := range results{
		fmt.Printf("%s: \n%s\n\n",result.Field, result.Content)
	}
}