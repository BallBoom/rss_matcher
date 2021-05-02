package search

// 默认实现的匹配器

type defaultMatcher struct{}

// 初始化时 注册默认匹配器
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 实现matcher接口方法
func (d defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
