package eat

import (
	"fmt"
	"main/common"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

var mealList []common.TravelObj

func Meal(c *colly.Collector) {

	meal := map[string]string{}
	// Find and visit all links
	////*[@id="app"]/section/div/div[2]/div[1]/div/div[1]/ul

	c.OnHTML(".more.clear", func(e *colly.HTMLElement) {
		//fmt.Println("response", e.Text)
		e.ForEach("li a", func(i int, h *colly.HTMLElement) {
			meal[h.Text] = h.Attr("href")
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(common.RenderServer + "https://sz.meituan.com/meishi/")
	if err != nil {
		fmt.Printf("meituan visting error: %v\n", err.Error())
	}

	fmt.Printf("meals: %v", meal)
}

func MealSearch(c *gin.Context) {
	collyObj := colly.NewCollector()
	Meal(collyObj)
	common.CommJOSN(c, 200, mealList)
}
