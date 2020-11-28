package travel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/common"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TravelTrip struct {
	ID         int                `json:"trip_id"`
	CommonCard []common.TravelObj `json:"cards"`
}

type Trip struct {
	Title string `json:"title"`
	ID    int    `json:"trip_id"`
	Color string `json:"color"`
}

type HomepageObj struct {
	NowLocation    string `json:"location"`
	RecommendTrips []Trip `json:"recommend_trips"`
	TitleTips      string `json:"title_tips"`
	Destination    string `json:"destination"`
	Notice         string `json:"notice"`
}

var DBClient *sql.DB

func init() {
	DBClient = db.GetDB()
}

func TravelHomePage(c *gin.Context) {
	ipObj := IPToLoc(c)

	location := ipObj.City
	if len(ipObj.City) == 0 {
		location = ipObj.Country
	}
	page := HomepageObj{
		NowLocation: location,
		RecommendTrips: []Trip{
			{Title: "浪漫大梅沙三天两夜", Color: "orange"},
			{Title: "KLOOK精选情人岛行程", Color: "blue"},
			{Title: "惠州双月湾", Color: "green"},
			{Title: "罗湖美食之旅", Color: "green"},
		},
		TitleTips:   "打造无与伦比的行程•探索未知的你",
		Destination: "深圳",
		Notice:      "* 关于规划行程中展现的价格与时间为估算值，仅做参考",
	}
	common.CommJOSN(c, 200, page)
}

type IPGeoData struct {
	Area      string `json:"area"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	IP        string `json:"ip"`
	Isp       string `json:"isp"`
	LongIP    string `json:"long_ip"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
}

type IPGeoInfo struct {
	Data  IPGeoData `json:"data"`
	LogID string    `json:"log_id"`
	Msg   string    `json:"msg"`
	Ret   int64     `json:"ret"`
}

var appCode = "c4387c3c3422485fb072a5ead254d226"

// curl -i -k --get --include 'https://api01.aliyun.venuscn.com/ip?ip=218.18.228.178'  -H 'Authorization:APPCODE 你自己的AppCode'
func IPToLoc(c *gin.Context) IPGeoData {
	baseUrl := "https://api01.aliyun.venuscn.com/ip?ip=%s"
	client := &http.Client{}
	ipInfo := c.ClientIP()
	fmt.Printf("ip info: %v", ipInfo)
	req, err := http.NewRequest("GET", fmt.Sprintf(baseUrl, c.ClientIP()), nil)
	if err != nil {
		fmt.Printf("get error: %v", err.Error())
	}
	req.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", appCode))

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("request url get error: %v", err.Error())
	}
	defer response.Body.Close()

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("readall get error: %v", err.Error())
	}

	data := IPGeoInfo{}
	json.Unmarshal(rawBody, &data)
	return data.Data
	//c.JSON(200, data)
}

const uniqueIDKey = "uniqueID"
const uniqueIDTripKey = "uniqueID:%d:trip"

func TravelCardSubmit(c *gin.Context) {
	//dbObj := db.ReplDB{}
	var travelParam common.TravelAdd
	jsonerr := c.BindJSON(&travelParam)

	smt, err := DBClient.Prepare(`insert into trip(nums, destination, travel_date) values(?,?,?)`)
	if err != nil {
		fmt.Printf("db client preare error: %v", err.Error())
	}

	args := []interface{}{}
	if jsonerr == nil {
		args = []interface{}{travelParam.TravelNum, travelParam.Destination, travelParam.TravelTime}
	}

	ret, err := smt.Exec(args...)
	if err != nil {
		fmt.Printf("db client exec error: %v", err.Error())
	} else {
		fmt.Printf("insert rows %v", ret)
	}
	insertID, _ := ret.LastInsertId()

	tripObj := TravelTrip{
		ID: int(insertID),
	}
	common.CommJOSN(c, 200, tripObj)
}

func TravelCardUpdate(c *gin.Context) {

}

func TravelCardDel(c *gin.Context) {

}

func TravelCardList(c *gin.Context) {

}
