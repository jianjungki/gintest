package hotel

import (
	"fmt"
	"main/common"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Hotel(c *colly.Collector) {
	//hotel := map[string]string{}
	// Find and visit all links
	////*[@id="app"]/section/div/div[2]/div[1]/div/div[1]/ul
	c.OnHTML(".long-list", func(e *colly.HTMLElement) {

		e.ForEach(".with-decorator-wrap", func(i int, h *colly.HTMLElement) {
			fmt.Printf("title: %s\n", h.ChildText("span.name.font-bold"))
			fmt.Printf("point: %s\n", h.ChildText(".list-card-encourage p.review span"))
			//特点
			h.ForEach(".list-card-tag .card-tag", func(j int, g *colly.HTMLElement) {
				fmt.Printf("special: %s\n", g.Text)
			})

			fmt.Printf("review count: %s\n", h.ChildText(".list-card-comment .describe .count"))
			fmt.Printf("score: %s\n", h.ChildText(".list-card-comment .score"))

			fmt.Printf("price: %s\n", h.ChildText("p.price>span.real-price.font-bold"))
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//2020/11/29
	baseUrl := "https://hotels.ctrip.com/hotels/list"
	baseUrl += "?countryId=1&city=%d&checkin=%s&checkout=%s&starlist=%d&optionId=%d&optionType=City&directSearch=0&crn=1&adult=1&children=0&searchBoxArg=t&travelPurpose=0&ctm_ref=ix_sb_dl&domestic=1"
	err := c.Visit(fmt.Sprintf(common.RenderServer+baseUrl, 43, "2020/11/29", "2020/11/31", 2, 43))
	if err != nil {
		fmt.Printf("ctrip visting error: %v\n", err.Error())
	}
}

func HotelSearch(c *gin.Context) {
	collyObj := colly.NewCollector()
	Hotel(collyObj)
}
