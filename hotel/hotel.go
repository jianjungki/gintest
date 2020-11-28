package hotel

import (
	"fmt"
	"main/common"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

var hotelList []common.TravelObj

func Hotel(c *colly.Collector, params common.CommonReq) {
	//hotel := map[string]string{}
	// Find and visit all links
	////*[@id="app"]/section/div/div[2]/div[1]/div/div[1]/ul
	c.OnHTML(".long-list", func(e *colly.HTMLElement) {

		e.ForEach(".with-decorator-wrap", func(i int, h *colly.HTMLElement) {
			cardItem := common.TravelObj{
				CommonCard: common.CommonCard{
					Title: h.ChildText("span.name.font-bold"),
					//Price:  h.ChildText("p.price>span.real-price.font-bold"),
					LiveDays:  params.LiveDays,
					RoomNum:   params.RoomNum,
					Desc:      common.DescText,
					ReviewNum: h.ChildText(".list-card-comment .describe .count"),
					Image:     "https://dimg04.c-ctrip.com/images/020691200082co240B8EA_R_300_225_R5_Q70_D.jpg",
					TripID:    params.TripID,
					//Score:     h.ChildText(".list-card-comment .score"),
					//Price: h.ChildText("p.price>span.real-price.font-bold"),
				},
				CardType: common.Hotel,
			}

			h.ForEach(".list-card-tag .card-tag", func(j int, g *colly.HTMLElement) {
				cardItem.SellPoint = append(cardItem.SellPoint, g.Text)
			})
			//第三个
			if len(hotelList) == 3 {
				cardItem.Selected = 1
			}
			hotelList = append(hotelList, cardItem)
			/*
				fmt.Printf("title: %s\n", h.ChildText("span.name.font-bold"))
				fmt.Printf("point: %s\n", h.ChildText(".list-card-encourage p.review span"))
				//特点


				fmt.Printf("review count: %s\n", h.ChildText(".list-card-comment .describe .count"))
				fmt.Printf("score: %s\n", h.ChildText(".list-card-comment .score"))

				fmt.Printf("price: %s\n", h.ChildText("p.price>span.real-price.font-bold"))
			*/
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//2020/11/29
	baseUrl := "https://hotels.ctrip.com/hotels/list"
	baseUrl += "?countryId=1&city=%d&checkin=%s&checkout=%s&starlist=%d&optionId=%d&optionType=City&directSearch=0&optionName=%s&crn=%d&adult=%d&children=0&searchBoxArg=t&travelPurpose=0&ctm_ref=ix_sb_dl&domestic=1"
	err := c.Visit(fmt.Sprintf(common.RenderServer+baseUrl, 43, "2020/11/29", "2020/11/31", params.HotelRate, 43, params.AroundPos, params.RoomNum, params.Person))
	if err != nil {
		fmt.Printf("ctrip visting error: %v\n", err.Error())
	}
}

func HotelSearch(c *gin.Context) {
	commReq := common.CommonReq{}
	if err := c.BindJSON(&commReq); err == nil {
		collyObj := colly.NewCollector()
		Hotel(collyObj, commReq)

		common.AddTravelRecord(commReq.TripID, hotelList[3])
		common.CommJOSN(c, 200, hotelList)
	} else {
		common.FaildJOSN(c, 200, "")
	}

}
