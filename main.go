package main

import (
 "github.com/gin-gonic/gin"
 "main/transfer"
 "main/hotel"
 "main/play"
 "main/eat"
 "main/travel"
)


func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("index.html")
  r.GET("/homepage", travel.TravelHomePage)

  r.POST("/travel/add", travel.TravelCardSubmit)

	r.GET("/travel/list", func(c *gin.Context) {
		c.JSON(200, "")
	})

  r.POST("/travel/save", func(c *gin.Context) {
		c.JSON(200, "")
	})

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
