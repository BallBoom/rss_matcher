package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"main/search"
	"net/http"
	"regexp"
)

// 搜索rss源的匹配器

type (
	rssDocment struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		Guid        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
)

// rss匹配器，实现 Matcher 接口
type rssMatcher struct{}

// 将匹配器注册到程序
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// 实现matcher接口
func (r rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	//todo
	var results []*search.Result

	log.Printf("search feed type[%s] site[%s] for uri[%s]\n",
		feed.Type, feed.Name, feed.URI)

	// 执行搜索结果
	document, err := r.retrieve(feed)
	// log.Fatalln("document is = ", document)
	if err != nil {
		return nil, err
	}

	for _, item := range document.Channel.Item {
		// 检查标题部门是否含有搜索项
		matched, err := regexp.MatchString(searchTerm, item.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: item.Title,
			})
		}

		matched, err = regexp.MatchString(searchTerm, item.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: item.Description,
			})
		}

	}
	return results, err
}

// rss匹配器执行http 请求搜索，返回数据及解码
func (r rssMatcher) retrieve(feed *search.Feed) (*rssDocment, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// 从网络获取数据
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocment

	// 将rss连接获取到的xml数据  解码到document中
	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err
}
