package play

import (
	"fmt"
	"main/common"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

var playList []common.TravelObj

func Play(c *colly.Collector) {

	//hotel := map[string]string{}
	// Find and visit all links
	////*[@id="app"]/section/div/div[2]/div[1]/div/div[1]/ul

	c.OnHTML("#container", func(e *colly.HTMLElement) {

		e.ForEach(".row.row-top5 div.item.clearfix div.info div.middle", func(i int, h *colly.HTMLElement) {
			playList = append(playList, common.TravelObj{
				CommonCard: common.CommonCard{
					SellPoint: []string{h.ChildText("p")},
					Title:     h.ChildText("h3 a:nth-child(2)"),
					ReviewNum: h.ChildText("h3 a:nth-child(3) .rev-total"),
				}})
			//fmt.Printf("point: %s\n", h.ChildText("p"))
			//fmt.Printf("title: %s\n", h.ChildText("h3 a:nth-child(2)"))
			//fmt.Printf("reviews: %s\n", h.ChildText("h3 a:nth-child(3) .rev-total"))
			//fmt.Printf("desc: %s", h.ChildText("div.links:nth-child(3)"))
			//fmt.Printf("point: %s\n", h.ChildText("div.info>div.middle p"))
			//fmt.Printf("point: %s\n", h.ChildText("div.info>div.middle p"))
		})

		//景点
		//
		//e.ForEach("div.row.row-allScenic:nth-child(5) div.wrapper div.bd ul.scenic-list.clearfix", func(i int, h *colly.HTMLElement) {
		//fmt.Printf("point: %s\n", h.ChildAttr("li:nth-child(1) > a:nth-child(1)", "href"))
		//fmt.Printf("img: %s\n", h.ChildText("li a:nth-child(1)>h3:nth-child(2)"))
		//fmt.Printf("desc: %s", h.ChildText("div.links:nth-child(3)"))
		//fmt.Printf("point: %s\n", h.ChildText("div.info>div.middle p"))
		//fmt.Printf("point: %s\n", h.ChildText("div.info>div.middle p"))
		//})

		//fmt.Printf("title: %s\n", e.ChildText(".row .item .clearfix a[0]"))
		//fmt.Printf("point: %s\n", e.ChildText(".list-card-encourage p.review span"))
	})
	/*
		c.OnResponse(func(resp *colly.Response) {
			fmt.Printf("resp %s", string(resp.Body))
		})
	*/

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//2020/11/29
	baseUrl := "https://www.mafengwo.cn/jd/10198/gonglve.html"
	err := c.Visit(common.RenderServer + baseUrl)
	if err != nil {
		fmt.Printf("mafengwo visting error: %v\n", err.Error())
	}

}

func PlaySearch(c *gin.Context) {
	collyObj := colly.NewCollector()
	Play(collyObj)
	common.CommJOSN(c, 200, playList)
}
