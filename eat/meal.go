package eat

import (
	"fmt"
	"main/common"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

var mealList []common.TravelObj

func Meal(c *colly.Collector) {
	// Find and visit all links
	////*[@id="app"]/section/div/div[2]/div[1]/div/div[1]/ul

	c.OnHTML("#shop-all-list ul", func(e *colly.HTMLElement) {

		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			fmt.Println("image", h.ChildAttr("div.pic a img", "src"))
			fmt.Println("title", h.ChildText("div.txt div.tit a>h4"))

			sellPoints := make([]string, 0)
			h.ForEach("div.txt a.recommend-click", func(j int, g *colly.HTMLElement) {
				sellPoints = append(sellPoints, g.Text)
			})

			fmt.Println("price", h.ChildText("div.comment a.mean-price>b:nth-child(1)"))

			re, _ := regexp.Compile("(/d+)")
			priceStr := re.FindString(h.ChildText("div.comment a.mean-price>b:nth-child(1)"))
			priceInt, _ := strconv.Atoi(priceStr)

			travelObj := common.TravelObj{
				CommonCard: common.CommonCard{
					Desc:      common.DescText,
					ReviewNum: h.ChildText("div.comment>a.review-num"),
					SellPoint: sellPoints,
					Image:     h.ChildAttr("div.pic a img", "src"),
					Title:     h.ChildText("div.txt div.tit a>h4"),
					Price:     priceInt,
				},
				CardType: common.Eat,
			}
			//fmt.Println("response", h.ChildAttr("div.pic a img", "src"))
			//fmt.Println("response", h.ChildAttr("div.pic a img", "src"))
			mealList = append(mealList, travelObj)
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//
	err := c.Visit("http://36555474-1911004356767178.test.functioncompute.com/render/http://www.dianping.com/shenzhen/ch10")
	if err != nil {
		fmt.Printf("meituan visting error: %v\n", err.Error())
	}

	//fmt.Printf("meals: %v", meal)
}

func MealSearch(c *gin.Context) {
	commReq := common.CommonReq{}
	if err := c.BindJSON(&commReq); err == nil {
		collyObj := colly.NewCollector()
		Meal(collyObj)
		common.AddTravelRecord(commReq.TripID, mealList[3])
		common.CommJOSN(c, 200, mealList)
	} else {
		common.FaildJOSN(c, 200, "")
	}
}
