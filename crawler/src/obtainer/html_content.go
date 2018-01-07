package obtainer

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HtmlContent struct {
	url      string
	content  string
	document *goquery.Document
}

func NewHtmlContent() *HtmlContent {
	return &HtmlContent{"", "", nil}
}

func (self *HtmlContent) Load(url, content string) bool {
	reader := strings.NewReader(content)
	utf_8_reader, err := DecodeHTMLBody_Utf8(reader, "")
	if nil != err {
		return false
	}

	doc, err := goquery.NewDocumentFromReader(utf_8_reader)
	if nil != err {
		return false
	}
	self.url = url
	self.document = doc
	self.content = content

	return nil != doc
}

func (self *HtmlContent) Title() string {
	return ""
}

func (self *HtmlContent) SubHrefUrls() []string {
	if nil == self.document {
		return []string{}
	}

	urls := make([]string, 0)
	link := self.document.Find("a")
	host, err := url.Parse(self.url)
	if nil == err {
		link.Each(func(i int, content *goquery.Selection) {
			href, _ := content.Attr("href")
			u, err := url.Parse(href)
			if nil != err {
				return
			}
			real := host.ResolveReference(u).String()
			real = strings.TrimSpace(real)
			real_lower := strings.ToLower(real)

			if strings.HasPrefix(real_lower, "https://") ||
				strings.HasPrefix(real_lower, "http://") {
				index := strings.IndexAny(real_lower, "#")
				if index > 0 {
					urls = append(urls, real_lower[0:index])
				} else {
					urls = append(urls, real_lower)
				}
			}
		})
	}
	return urls
}

func (self *HtmlContent) ContentText() string {
	src := self.content

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}
