package common

import "github.com/gin-gonic/gin"

type CommonCard struct  {
    ReviewNum string `json:"review_num"` //评论数
    Price string `json:"price"` //价格
    Title string `json:"title"` //标题
    EstimatedTime string `json:"estimate_time"`
    Desc string `json:"desc"` //描述
    Image string `json:"image"` //图片
    TravelNum int `json:"travel_num"` //游玩人数
    Location string `json:"location"` //大概位置
    SellPoint string `json:"sell_point"` //卖点
}

var RenderServer = ""

//CommonResp
type CommonResp struct{
  Code int `json:"code"`
  Message string `json:"msg"`
  Data interface{} `json:"data"`
}

func CommJOSN(c *gin.Context, statusCode int,resp interface{}){
    respObj := CommonResp{
      Code: 0,
      Message: "成功",
      Data: resp,
    }
    c.JSON(statusCode, respObj)
}

func FaildJOSN(c *gin.Context, statusCode int,resp interface{}){
    respObj := CommonResp{
      Code: 1,
      Message: "失败",
      Data: resp,
    }
    c.JSON(statusCode, respObj)
}