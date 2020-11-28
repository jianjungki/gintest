package common

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type CommonCard struct {
	ReviewNum     string   `json:"review_num"`    //评论数
	Price         int      `json:"price"`         //价格
	Title         string   `json:"title"`         //标题
	EstimatedTime int      `json:"estimate_time"` //分钟
	Desc          string   `json:"desc"`          //描述
	Image         string   `json:"image"`         //图片
	TravelNum     int      `json:"travel_num"`    //游玩人数
	Location      string   `json:"location"`      //大概位置
	SellPoint     []string `json:"sell_point"`    //卖点

	Selected int `json:"selector"`
	TripID   int `json:"trip_id"` //行程id
}
type TransferObj struct {
	StartCity string `json:"start"`
	DestCity  string `json:"end"`
	RunTime   string `json:"run_time"`
	//0 火车 1 汽车 2 公交 3 机票
	TransferType int `json:"type"`
}

type TravelObj struct {
	CommonCard
	TransferObj

	CardType int `json:"card_type"`
}

const (
	Trasfer int = iota
	Eat
	Hotel
	Play
)

type TravelAdd struct {
	TravelNum   int    `json:"person"`
	Destination string `json:"destination"`
	TravelTime  string `json:"travel_time"`
}

var RenderServer = ""

//CommonResp
type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func CommJOSN(c *gin.Context, statusCode int, resp interface{}) {

	respObj := CommonResp{
		Code:    0,
		Message: "成功",
		Data:    resp,
	}
	if reflect.TypeOf(resp).Kind() == reflect.Slice {
		respObj = CommonResp{
			Code:    0,
			Message: "成功",
			Data: map[string]interface{}{
				"list": resp,
			},
		}

	}

	c.JSON(statusCode, respObj)

}

func FaildJOSN(c *gin.Context, statusCode int, resp interface{}) {
	respObj := CommonResp{
		Code:    1,
		Message: "失败",
		Data:    resp,
	}
	c.JSON(statusCode, respObj)
}

type CommonReq struct {
	TripID    int `json:"trip_id"`
	Category  int `json:"category"`
	Price     int `json:"price"`
	Person    int `json:"person"`
	HotelRate int `json:"hotel_rate"`
}
