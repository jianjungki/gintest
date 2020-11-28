package transfer

import (
	"fmt"
	"io/ioutil"
	"main/common"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"encoding/json"
)

var appCode = "c4387c3c3422485fb072a5ead254d226"

const TIME_LAYOUT = "2006-01-02 15:04:05"

func TransferSearch(c *gin.Context) {
	//transferType, _ := strconv.Atoi(c.get("type"))
	req := TransferReq{}
	resp := make([]common.TravelObj, 0)
	if err := c.BindJSON(&req); err == nil {
		resp = append(resp, TrainSearch(c, req))
		resp = append(resp, FlightSearch(c, req))

		common.CommJOSN(c, 200, resp)
	} else {
		common.FaildJOSN(c, 200, "")
	}

}

//curl -i -k -X ANY 'https://jisugjdtmf.market.alicloudapi.com/transit/station2s?city=%E6%9D%AD%E5%B7%9E&end=%E6%9D%AD%E5%B7%9E%E6%B1%BD%E8%BD%A6%E5%8C%97%E7%AB%99&endcity=endcity&start=%E8%A5%BF%E6%BA%AA%E7%AB%9E%E8%88%9F%E8%8B%91&type=transit'  -H 'Authorization:APPCODE 你自己的AppCode'
func LocalTransSearch(c *gin.Context, req TransferReq) common.TravelObj {
	baseUrl := "https://jisugjdtmf.market.alicloudapi.com/transit/station2s?city=深圳&end=深圳火车站&start=宝安机场&type=transit"
	transobj := common.TravelObj{}
	client := &http.Client{}
	reqObj, err := http.NewRequest("GET", fmt.Sprintf(baseUrl, req.ArriveCty, req.StartCty), nil)
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

	data := LocalTransData{}
	err = json.Unmarshal(rawBody, &data)
	if err != nil {
		fmt.Printf("json unmarshal error: %v", err.Error())
	}
	if len(data.Result) > 0 {
		resultObj := data.Result[0]
		transobj.RunTime = resultObj.Totalduration
	}

	return transobj
}

type TransferReq struct {
	TransferType int    `json:"type"`
	StartCty     string `json:"leave"`
	ArriveCty    string `json:"arrive"`
	StartDate    string `json:"sart_date"`
	EndDate      string `json:"end_date"`
}

//curl -i -k -X ANY 'https://jisutrain.market.alicloudapi.com/train/station2s?date=2019-11-21&end=%E5%8C%97%E4%BA%AC&ishigh=0&start=%E6%9D%AD%E5%B7%9E'  -H 'Authorization:APPCODE 你自己的AppCode'
func TrainSearch(c *gin.Context, req TransferReq) common.TravelObj {
	baseUrl := "https://jisutrain.market.alicloudapi.com/train/station2s?date=%s&end=%s&ishigh=1&start=%s"

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

	runTimeMinute := ""
	price := 0
	result := data.Result
	if len(data.Result.List) > 0 {
		trainItem := data.Result.List[0]
		runTimeMinute = trainItem.Costtime
		price = int(trainItem.Priceed)
	}

	transObj := common.TravelObj{
		TransferObj: common.TransferObj{
			StartCity:    result.Start,
			DestCity:     result.End,
			RunTime:      runTimeMinute,
			TransferType: 0,
		},
		CommonCard: common.CommonCard{
			Price:     price,
			TravelNum: 2,
		},
		CardType: 0,
	}
	return transObj
	//common.CommJOSN(c, 200, transObj)

}

//http://plane.market.alicloudapi.com/ai_market/ai_airplane/get_airplane_list?END_CITY=%E6%8A%B5%E8%BE%BE%E5%9F%8E%E5%B8%82&END_DATE=%E8%BF%94%E7%A8%8B%E6%97%A5%E6%9C%9F&START_CITY=%E5%87%BA%E5%8F%91%E5%9F%8E%E5%B8%82&START_DATE=%E5%87%BA%E5%8F%91%E6%97%A5%E6%9C%9F'  -H 'Authorization:APPCODE 你自己的AppCode'
//secret n41OKG54j0MYHd7ni6AAZ4OOEhCaCGQg
//code c4387c3c3422485fb072a5ead254d226
func FlightSearch(c *gin.Context, req TransferReq) common.TravelObj {

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
	err = json.Unmarshal(rawBody, &data)
	if err != nil {
		fmt.Printf("unmarshal get error: %v", err.Error())
	}

	runTimeMinute := 0
	if len(data.Flights) > 0 {
		flightItem := data.Flights[0]
		endDateTime := flightItem.EndDate + " " + flightItem.EndTime + ":00"
		startDateTime := flightItem.StartDate + " " + flightItem.StartTime + ":00"

		fmt.Printf("endDateTime: %s", endDateTime)
		fmt.Printf("startDateTime: %s", startDateTime)
		endTime, _ := time.Parse(TIME_LAYOUT, endDateTime)
		startTime, _ := time.Parse(TIME_LAYOUT, startDateTime)

		runTime := endTime.Sub(startTime)
		runTimeMinute = int(runTime.Hours()*24 + runTime.Minutes())
	}

	transObj := common.TravelObj{
		TransferObj: common.TransferObj{
			StartCity:    data.StartCity,
			DestCity:     data.EndCity,
			RunTime:      fmt.Sprintf("%d", runTimeMinute),
			TransferType: 3,
		},
		CommonCard: common.CommonCard{
			Price:     rand.Intn(1000),
			TravelNum: 2,
		},
		CardType: 0,
	}
	return transObj
}
