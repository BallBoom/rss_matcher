package search

import (
	"log"
	"sync"
)

// 执行搜索的主控逻辑

//注册用于搜索的匹配器
var matchers = make(map[string]Matcher)

func Register(tp string, matcher Matcher) {
	matchers[tp] = matcher
}

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	log.Println("feeds =", &feeds)
	if err != nil {
		log.Fatal(err)
	}

	// 创建无缓冲通道 接收匹配结果
	results := make(chan *Result)

	// wait 以便处理所有数据源
	var waitGroup sync.WaitGroup

	// 添加处理所有数据源的groutine 数量
	waitGroup.Add(len(feeds))

	// 给每个数据原启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			// 如果不存在此类型匹配器  则给一个默认的，以便不会出现 阻塞现象
			matcher = matchers["default"]
		}

		// 启动一个goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			// 执行匹配的方法
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)

	}

	// 启动一个goroutine来监控是否所有的搜索工作都做完
	go func() {
		// 等待所有的 goroutine跑完
		waitGroup.Wait()

		// 关闭通道
		close(results)
	}()

	Display(results)

}
