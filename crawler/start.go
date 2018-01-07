package main

import (
	"fmt"
	"obtainer"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			url := "http://php.net/manual/zh/"
			html := obtainer.NewHtml()
			html.LoadUrl(url, 30)
			html.SetUserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
			//html.SetProxy("http://219.138.58.193:3128")
			html.Get("")
			content, success := html.WaitContent()

			if success {
				html_content := obtainer.NewHtmlContent()
				html_content.Load(url, content)
				//html_content.SubHrefUrls()
				fmt.Println(content)
				fmt.Println(html_content.ContentText())
			} else {
				fmt.Print("fail")
			}
		}()
	}
	/*
		center := dominate.NewCenter()
		center.LoadConf("")

		routune := dominate.NewRoutune()

		routune.Start(func() {
			fmt.Println("start")
		})
		time.Sleep(time.Second * 10000)
	*/
	time.Sleep(time.Second * 10000)
}
