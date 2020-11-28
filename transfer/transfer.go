package transfer

import (
	"fmt"
	"io/ioutil"
	"main/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"strconv"
)

var appCode = "c4387c3c3422485fb072a5ead254d226"

const TIME_LAYOUT = "2006-01-02 15:04:05"

func TransferSearch(c *gin.Context) {
	//transferType, _ := strconv.Atoi(c.get("type"))
	req := TransferReq{}
	if err := c.Bind(&req); err != nil {
		switch req.TransferType {
		case 0:
			TrainSearch(c, req)
		case 1:
			LocalTransSearch(c)
		case 2:
			LocalTransSearch(c)
		case 3:
			FlightSearch(c, req)
		}
	} else {
		common.FaildJOSN(c, 200, "")
	}

}

func LocalTransSearch(c *gin.Context) {
	c.HTML(200, "wait for develop", nil)
	//return
}

type TransferReq struct {
	TransferType int    `json:"type"`
	StartCty     string `json:"leave"`
	ArriveCty    string `json:"arrive"`
	StartDate    string `json:"sart_date"`
	EndDate      string `json:"end_date"`
}

//curl -i -k -X ANY 'https://jisutrain.market.alicloudapi.com/train/station2s?date=2019-11-21&end=%E5%8C%97%E4%BA%AC&ishigh=0&start=%E6%9D%AD%E5%B7%9E'  -H 'Authorization:APPCODE 你自己的AppCode'
func TrainSearch(c *gin.Context, req TransferReq) {
	baseUrl := "https://jisutrain.market.alicloudapi.com/train/station2s?date=%s&end=%s&ishigh=1&start=%s"
	/*
		leaveCity := c.PostForm("leave")
		arriveCity := c.PostForm("arrive")
		start_date := c.PostForm("start_date")
	*/

	fmt.Printf("req url: %s\n", fmt.Sprintf(baseUrl, req.StartDate, req.ArriveCty, req.StartCty))

	client := &http.Client{}
	reqObj, err := http.NewRequest("GET", fmt.Sprintf(baseUrl, req.StartDate, req.ArriveCty, req.StartCty), nil)
	if err != nil {
		fmt.Printf("get error: %v", err.Error())
	}
	reqObj.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", appCode))

	response, err := client.Do(reqObj)
	if err != nil {
		fmt.Printf("request url get error: %v", err.Error())
	}
	defer response.Body.Close()

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("readall get error: %v", err.Error())
	}

	data := TrainData{}
	err = json.Unmarshal(rawBody, &data)
	if err != nil {
		fmt.Printf("json unmarshal error: %v", err.Error())
	}

	runTimeMinute := 0
	price := 0
	result := data.Result
	if len(data.Result.List) > 0 {
		trainItem := data.Result.List[0]
		runTimeMinute, _ = strconv.Atoi(trainItem.Costtime)
		price = int(trainItem.Priceed)
	}

	transObj := common.TravelObj{
		TransferObj: common.TransferObj{
			StartCity: result.Start,
			DestCity:  result.End,
			//Price:        price,
			RunTime:      runTimeMinute,
			TransferType: 0,
			//TravelNum:    2,
		},
		CommonCard: common.CommonCard{
			Price:     price,
			TravelNum: 2,
		},
		CardType: 0,
	}
	common.CommJOSN(c, 200, transObj)

}

//http://plane.market.alicloudapi.com/ai_market/ai_airplane/get_airplane_list?END_CITY=%E6%8A%B5%E8%BE%BE%E5%9F%8E%E5%B8%82&END_DATE=%E8%BF%94%E7%A8%8B%E6%97%A5%E6%9C%9F&START_CITY=%E5%87%BA%E5%8F%91%E5%9F%8E%E5%B8%82&START_DATE=%E5%87%BA%E5%8F%91%E6%97%A5%E6%9C%9F'  -H 'Authorization:APPCODE 你自己的AppCode'
//secret n41OKG54j0MYHd7ni6AAZ4OOEhCaCGQg
//code c4387c3c3422485fb072a5ead254d226
func FlightSearch(c *gin.Context, req TransferReq) {

	baseUrl := "http://plane.market.alicloudapi.com/ai_market/ai_airplane/get_airplane_list?END_CITY=%s&START_CITY=%s&START_DATE=%s"

	/*
		leaveCity := c.PostForm("leave")
		arriveCity := c.PostForm("arrive")
		start_date := c.PostForm("start_date")
		end_date := c.PostForm("end_date")
	*/
	fmt.Printf("req url: %s\n", fmt.Sprintf(baseUrl, req.ArriveCty, req.StartCty, req.StartDate))

	client := &http.Client{}
	reqObj, err := http.NewRequest("GET", fmt.Sprintf(baseUrl, req.ArriveCty, req.StartCty, req.StartDate), nil)
	if err != nil {
		fmt.Printf("get error: %v", err.Error())
	}
	reqObj.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", appCode))

	response, err := client.Do(reqObj)
	if err != nil {
		fmt.Printf("request url get error: %v", err.Error())
	}
	defer response.Body.Close()
	fmt.Printf("response %v", response.Body)

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("readall get error: %v", err.Error())
	}

	data := FligtData{}
	json.Unmarshal(rawBody, &data)

	runTimeMinute := 0
	if len(data.Flights) > 0 {
		flightItem := data.Flights[0]
		endDateTime := flightItem.EndDate + " " + flightItem.EndTime + ":00"
		startDateTime := flightItem.StartDate + " " + flightItem.StartTime + ":00"

		endTime, _ := time.Parse(TIME_LAYOUT, endDateTime)
		startTime, _ := time.Parse(TIME_LAYOUT, startDateTime)
		runTime := endTime.Sub(startTime)
		runTimeMinute = int(runTime.Hours()*24 + runTime.Minutes())
	}

	transObj := common.TravelObj{
		TransferObj: common.TransferObj{
			StartCity: data.StartCity,
			DestCity:  data.EndCity,
			//Price:        0,
			RunTime:      runTimeMinute,
			TransferType: 3,
			//TravelNum:    2,
		},
		CommonCard: common.CommonCard{
			Price:     0,
			TravelNum: 2,
		},
		CardType: 0,
	}

	common.CommJOSN(c, 200, transObj)

}
