package transfer

var flightStation = map[string]string{
	"阿勒泰":      "AAT",
	"安康":       "AKA",
	"阿克苏":      "AKU",
	"安庆":       "AQC",
	"包头":       "BAV",
	"北海":       "BHY",
	"北京":       "BJS",
	"保山":       "BSD",
	"广州":       "CAN",
	"常德":       "CGD",
	"郑州":       "CGO",
	"长春":       "CGQ",
	"朝阳":       "CHG",
	"酒泉":       "CHW",
	"赤峰":       "CIF",
	"长治":       "CIH",
	"重庆":       "CKG",
	"长海":       "CNI",
	"长沙":       "CSX",
	"成都":       "CTU",
	"常州":       "CZX",
	"大同":       "DAT",
	"达县":       "DAX",
	"丹东":       "DDG",
	"大连":       "DLC",
	"敦煌":       "DNH",
	"大庸":       "DYG",
	"恩施":       "ENH",
	"延安":       "ENY",
	"福州":       "FOC",
	"阜阳":       "FUG",
	"富蕴":       "FYN",
	"广汉":       "GHN",
	"格尔木":      "GOQ",
	"海口":       "HAK",
	"黑河":       "HEK",
	"呼和浩特":     "HET",
	"合肥":       "HFE",
	"杭州":       "HGH",
	"黄花机场(长沙)": "HHA",
	"海拉尔":      "HLD",
	"乌兰浩特":     "HLH",
	"哈密":       "HMI",
	"衡阳":       "HNY",
	"哈尔滨":      "HRB",
	"和田":       "HTN",
	"汉中":       "HZG",
	"银川":       "INC",
	"且末":       "IQM",
	"庆阳":       "IQN",
	"景德镇":      "JDZ",
	"嘉峪关":      "JGN",
	"九江":       "JIU",
	"晋江":       "JJN",
	"佳木斯":      "JMU",
	"库车":       "KCA",
	"喀什":       "KHG",
	"南昌":       "KHN",
	"昆明":       "KMG",
	"吉安":       "KNC",
	"赣州":       "KOW",
	"库尔勒":      "KRL",
	"克拉玛依":     "KRY",
	"贵阳":       "KWE",
	"桂林":       "KWL",
	"兰州":       "LHW",
	"梁平":       "LIA",
	"庐山":       "LUZ",
	"拉萨":       "LXA",
	"林西":       "LXI",
	"洛阳":       "LYA",
	"连云港":      "LYG",
	"临沂":       "LYI",
	"兰州东":      "LZD",
	"柳州":       "LZH",
	"牡丹江":      "MOG",
	"梅县":       "MXZ",
	"南充":       "NAO",
	"齐齐哈尔":     "NDG",
	"宁波":       "NGB",
	"南京":       "NKG",
	"南宁":       "NNG",
	"南阳":       "NNY",
	"首都机场(北京)": "PEK",
	"上海浦东":     "PVG",
	"上海":       "SHA",
	"沈阳":       "SHE",
	"山海关":      "SHP",
	"沙市":       "SHS",
	"西安":       "SIA",
	"汕头":       "SWA",
	"思茅":       "SYM",
	"三亚":       "SYX",
	"深圳":       "SZX",
	"青岛":       "TAO",
	"铜仁":       "TEN",
	"辽通":       "TGO",
	"济南":       "TNA",
	"天津":       "TSN",
	"屯溪":       "TXN",
	"太原":       "TYN",
	"乌鲁木齐":     "URC",
	"榆林":       "UYN",
	"武汉":       "WUH",
	"万县":       "WXN",
	"兴城":       "XEN",
	"襄樊":       "XFN",
	"西昌":       "XIC",
	"锡林浩特":     "XIL",
	"兴宁":       "XIN",
	"咸阳机场(西安)": "XIY",
	"厦门":       "XMN",
	"西宁":       "XNN",
	"徐州":       "XUZ",
	"宜昌":       "YIH",
	"伊宁":       "YIN",
	"依兰":       "YLN",
	"延吉":       "YNJ",
	"烟台":       "YNT",
	"昭通":       "ZAT",
	"中川机场(兰州)": "ZGC",
	"湛江":       "ZHA",
	"珠海":       "ZUH",
}

type LocalTransit struct {
	Msg    string `json:"msg"`
	Result []struct {
		Arrivetime string `json:"arrivetime"`
		Steps      []struct {
			Distance string `json:"distance"`
			Duration string `json:"duration"`
			Endname  string `json:"endname"`
			Endpoi   struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"endpoi"`
			Sname    string `json:"sname"`
			Startpoi struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"startpoi"`
			Steptext string `json:"steptext"`
			Type     string `json:"type"`
			Vehicle  struct {
				Endname    string `json:"endname"`
				Endtime    string `json:"endtime"`
				Name       string `json:"name"`
				Startname  string `json:"startname"`
				Starttime  string `json:"starttime"`
				Stopnum    string `json:"stopnum"`
				Totalprice string `json:"totalprice"`
				Type       string `json:"type"`
				Zoneprice  string `json:"zoneprice"`
			} `json:"vehicle"`
		} `json:"steps"`
		Tiptype           string   `json:"tiptype"`
		Totaldistance     string   `json:"totaldistance"`
		Totalduration     string   `json:"totalduration"`
		Totalprice        string   `json:"totalprice"`
		Totalstopnum      string   `json:"totalstopnum"`
		Totalwalkdistance string   `json:"totalwalkdistance"`
		Vehicles          []string `json:"vehicles"`
	} `json:"result"`
	Status int64 `json:"status"`
}

type FligtData struct {
	EndCity       string       `json:"END_CITY"`
	EndDate       string       `json:"END_DATE"`
	Flights       []FlightItem `json:"FLIGHTS"`
	FlightsStatus string       `json:"FLIGHTS_STATUS"`
	StartCity     string       `json:"START_CITY"`
	StartDate     string       `json:"START_DATE"`
}

type FlightItem struct {
	EndAirportCh      string `json:"END_AIRPORT_CH"`
	EndAirportEn      string `json:"END_AIRPORT_EN"`
	EndDate           string `json:"END_DATE"`
	EndTerminalEn     string `json:"END_TERMINAL_EN"`
	EndTime           string `json:"END_TIME"`
	FlightAirwaysCh   string `json:"FLIGHT_AIRWAYS_CH"`
	FlightAirwaysEn   string `json:"FLIGHT_AIRWAYS_EN"`
	FlightAirIsMorrow string `json:"FLIGHT_AIR_IS_MORROW"`
	FlightAirTime     string `json:"FLIGHT_AIR_TIME"`
	FlightID          string `json:"FLIGHT_ID"`
	FlightLowestPrice string `json:"FLIGHT_LOWEST_PRICE"`
	FlightPlaneCn     string `json:"FLIGHT_PLANE_CN"`
	FlightPlaneStyle  string `json:"FLIGHT_PLANE_STYLE"`
	FlightPlaneType   string `json:"FLIGHT_PLANE_TYPE"`
	StartAirportCh    string `json:"START_AIRPORT_CH"`
	StartAirportEn    string `json:"START_AIRPORT_EN"`
	StartDate         string `json:"START_DATE"`
	StartTerminalEn   string `json:"START_TERMINAL_EN"`
	StartTime         string `json:"START_TIME"`
}

type TrainData struct {
	Msg    string      `json:"msg"`
	Result TrainResult `json:"result"`
	Status int64       `json:"status"`
}

type TrainResult struct {
	Date  string      `json:"date"`
	End   string      `json:"end"`
	List  []TrainList `json:"list"`
	Start string      `json:"start"`
}

type TrainList struct {
	Arrivaltime   string  `json:"arrivaltime"`
	Costtime      string  `json:"costtime"`
	Day           int64   `json:"day"`
	Departuretime string  `json:"departuretime"`
	Distance      int64   `json:"distance"`
	Endstation    string  `json:"endstation"`
	Isend         int64   `json:"isend"`
	Pricedw       string  `json:"pricedw"`
	Pricedw1      string  `json:"pricedw1"`
	Priceed       float64 `json:"priceed"`
	Pricegr1      string  `json:"pricegr1"`
	Pricegr2      string  `json:"pricegr2"`
	Priceqt       string  `json:"priceqt"`
	Pricerw1      string  `json:"pricerw1"`
	Pricerw2      string  `json:"pricerw2"`
	Pricerz       string  `json:"pricerz"`
	Pricesw       int64   `json:"pricesw"`
	Pricetd       string  `json:"pricetd"`
	Pricewz       string  `json:"pricewz"`
	Priceyd       int64   `json:"priceyd"`
	Priceyw1      string  `json:"priceyw1"`
	Priceyw2      string  `json:"priceyw2"`
	Priceyw3      string  `json:"priceyw3"`
	Priceyz       string  `json:"priceyz"`
	Sequenceno    int64   `json:"sequenceno"`
	Station       string  `json:"station"`
	Trainno       string  `json:"trainno"`
	Trainno12306  string  `json:"trainno12306"`
	Type          string  `json:"type"`
	Typename      string  `json:"typename"`
}

type LocalTransData struct {
	Msg    string             `json:"msg"`
	Result []LocalTransResult `json:"result"`
	Status int64              `json:"status"`
}

type LocalTransPOI struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LocalTransStation struct {
	InterventType int64 `json:"intervent_type"`
	Underground   int64 `json:"underground"`
}

type LocalTransStep struct {
	CanRide           int64               `json:"can_ride"`
	Distance          int64               `json:"distance"`
	Duration          int64               `json:"duration"`
	Ename             string              `json:"ename"`
	Endname           string              `json:"endname"`
	Endpoi            LocalTransPOI       `json:"endpoi"`
	IsDepot           int64               `json:"is_depot"`
	SpathBegin        []int64             `json:"spath_begin"`
	SpathEnd          []int64             `json:"spath_end"`
	Startpoi          LocalTransPOI       `json:"startpoi"`
	Station           []LocalTransStation `json:"station"`
	Steptext          string              `json:"steptext"`
	SwalkEndLeadpoint []int64             `json:"swalk_end_leadpoint"`
	Tip               int64               `json:"tip"`
	TipBackground     string              `json:"tip_background"`
	TipText           string              `json:"tip_text"`
	TransType         int64               `json:"trans_type"`
	Type              int64               `json:"type"`
	Vehicle           []interface{}       `json:"vehicle"`
	WalkType          int64               `json:"walk_type"`
}

type LocalTransResult struct {
	Arrivetime        string           `json:"arrivetime"`
	Steps             []LocalTransStep `json:"steps"`
	Tiptype           int64            `json:"tiptype"`
	Totaldistance     string           `json:"totaldistance"`
	Totalduration     string           `json:"totalduration"`
	Totalprice        int64            `json:"totalprice"`
	Totalstopnum      int64            `json:"totalstopnum"`
	Totalwalkdistance string           `json:"totalwalkdistance"`
	Vehicles          []string         `json:"vehicles"`
}
