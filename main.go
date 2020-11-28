package main

import (
	"main/eat"
	"main/hotel"
	"main/play"
	"main/transfer"
	"main/travel"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("index.html")
	r.GET("/homepage", travel.TravelHomePage)

	r.POST("/travel/add", travel.TravelCardSubmit)

	r.GET("/travel/list", travel.TravelCardList)

	r.GET("/travel/del/:id", func(c *gin.Context) {
		c.JSON(200, "{}")
	})

	r.POST("/travel/update", func(c *gin.Context) {
		c.JSON(200, "{}")
	})

	r.POST("/transfer/query", transfer.TransferSearch)

	r.POST("/play/query", play.PlaySearch)

	r.POST("/hotel/query", hotel.HotelSearch)

	r.POST("/eat/query", eat.MealSearch)

	r.Run(":9000")
}
