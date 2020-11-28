package common

import (
	"database/sql"
	"fmt"
	"main/db"
	"reflect"

	"github.com/gin-gonic/gin"
)

var DBClient *sql.DB

func init() {
	DBClient = db.GetDB()
}

var DescText = `国内首个集休闲度假、观光旅游、户外运动、科普教育等主题于一体的大型综合性国家生态旅游示范区。
·大侠谷生态乐园和茶溪谷度假公园，集山地郊野和都市主题公园于一体，呈现了一个中西文化交融的世界。
·茵特拉根小镇、海菲德小镇和茶翁古镇，打造出童话一般的茶文化、葡萄酒文化梦幻小镇。
·每天还会有融合了多种艺术手段，以禅茶文化为主题的大型多媒体交响音画晚会《天禅》在大剧院上演。`

type CommonCard struct {
	ID            int      `json:"travel_id"`
	ReviewNum     string   `json:"review_num"`    //评论数
	Price         int      `json:"price"`         //价格
	Title         string   `json:"title"`         //标题
	EstimatedTime int      `json:"estimate_time"` //分钟
	Desc          string   `json:"desc"`          //描述
	Image         string   `json:"image"`         //图片
	TravelNum     int      `json:"travel_num"`    //游玩人数
	Location      string   `json:"location"`      //大概位置
	SellPoint     []string `json:"sell_point"`    //卖点
	Selected      int      `json:"selector"`
	TripID        int      `json:"tripid"` //行程id

	//AroundPos string `json:"around_pos"`
	RoomNum  int `json:"room_num"` //room_num
	LiveDays int `json:"leave_days"`
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
	TripID    int    `json:"tripid"`
	Category  int    `json:"category"`
	Price     int    `json:"price"`
	Person    int    `json:"person"`
	HotelRate int    `json:"hotel_rate"`
	LiveDays  int    `json:"leave_days"`
	AroundPos string `json:"around_pos"`
	RoomNum   int    `json:"room_num"`
}

func AddTravelRecord(tripId int, travelItem TravelObj) {
	if tripId > 0 {

		smt, err := DBClient.Prepare("insert into travel(type, transfer_type, start_city, end_city, price, `desc`, title, location, travel_num, travel_time, image, trip_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Printf("db client preare error: %v", err.Error())
		}

		args := []interface{}{travelItem.CardType, travelItem.TransferType, travelItem.StartCity,
			travelItem.DestCity, travelItem.Price, travelItem.Desc,
			travelItem.Title, travelItem.Location, travelItem.TravelNum,
			travelItem.EstimatedTime, travelItem.Image, tripId}

		ret, err := smt.Exec(args...)
		if err != nil {
			fmt.Printf("db client exec error: %v", err.Error())
		} else {
			fmt.Printf("insert rows %v", ret)
		}
	}
}
