package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
)

type AsyncSearchHandlerSOAIIRelpy struct  {
	RoundTripFlightList []struct {
		FareList []struct{
			Price int
			Tax int
			OilFee int
			ClassName string
			isOnlineCheckInTip bool
			TicketLack string
			Tips []struct {
				Content string
			}
			IsShowHKCarShipFree bool
		}
		FlightInfoKeys []string
		CanCompute int
		IsFree bool
		ContainsVirtualFlight bool
	}
	Eligibility []string
	//SearchStatus
}

func main() {
	GetGuonei()
	//getGuoji()
}

func getGuoji() {
	// http://flights.ctrip.com/international/AjaxRequest/SearchFlights/AsyncSearchHandlerSOAII.ashx
	SearchMode := `Search`
	condition := `{"FlightWay":"D","SegmentList":[{"DCityCode":"BJS","ACityCode":"PNH","DCity":"Beijing|北京(BJS)|1|BJS|480","ACity":"Phnom Penh|金边(PNH)|303|PNH|420","DepartDate":"2017-9-19"},{"DCityCode":"PNH","ACityCode":"BJS","DCity":"Phnom Penh|金边(PNH)|303|PNH|480","ACity":"Beijing|北京(BJS)|1|BJS|420","DepartDate":"2017-9-21"}],"TransferCityID":0,"Quantity":1,"ClassGrade":"Y_S","TransNo":"c03cefe271a94f17af28d36c5ff3cb2b","SearchRandomKey":"","RecommendedFlightSwitch":1,"EngineFlightSwitch":1,"SearchKey":"05803CE80A253DC6BA940854EDC935501402DD68745CC4825CF2645227F1C4B5D51B0DD157BD07FB03A5D8E87E02EC7B4260BCD0B0CE7F8F","MultiPriceUnitSwitch":1,"TransferCitySwitch":false,"EngineScoreABTest":"B","SearchStrategySwitch":1,"MaxSearchCount":3,"TicketRemarkSwitch":1,"RowNum":"1500","TicketRemarkChannels":["GDS-WS","ZY-WS"],"AddSearchLogOneByOne":true,"TFAirlineQTE":"AA","IsWifiPackage":0,"SegmentVerifySwitch":false,"ComparePriceByAttributeSwitch":true,"IsOpenCFNoDirectRecommendYS":false,"IsDomesticIntlPUVersionSwitch":true,"DisplayBaggageSizeSwitch":true,"IsOpen24Refund":true,"IsOpenTransPU":true,"IsOpenVirtualFlight":false}`
	//condition := `{"FlightWay":"D","SegmentList":[{"DCityCode":"BJS","ACityCode":"BKK","DCity":"Beijing|北京(BJS)|1|BJS|480","ACity":"Bangkok|曼谷(BKK)|359|BKK|420","DepartDate":"2017-9-15"},{"DCityCode":"BKK","ACityCode":"BJS","DCity":"Bangkok|曼谷(BKK)|359|BKK|480","ACity":"Beijing|北京(BJS)|1|BJS|420","DepartDate":"2017-9-22"}],"TransferCityID":0,"Quantity":1,"ClassGrade":"Y_S","TransNo":"8d1470420568417da5b52ed2d0ac2712","SearchRandomKey":"","RecommendedFlightSwitch":1,"EngineFlightSwitch":1,"SearchKey":"EAC529612B423DE79856A19CE0E914C6D33B7CAFB24FD845C3C1A4FB98C19DA02456467E21BEDBA7BD160C1874FECBFA3E438E587D303D6C","MultiPriceUnitSwitch":1,"TransferCitySwitch":false,"EngineScoreABTest":"B","SearchStrategySwitch":1,"MaxSearchCount":3,"TicketRemarkSwitch":1,"RowNum":"1500","TicketRemarkChannels":["GDS-WS","ZY-WS"],"AddSearchLogOneByOne":true,"TFAirlineQTE":"AA","IsWifiPackage":0,"SegmentVerifySwitch":false,"ComparePriceByAttributeSwitch":true,"IsOpenCFNoDirectRecommendYS":false,"IsDomesticIntlPUVersionSwitch":true,"DisplayBaggageSizeSwitch":true,"IsOpen24Refund":true,"IsOpenTransPU":true,"IsOpenVirtualFlight":false}`
	DisplayMode := "RoundTripGroup"
	SearchToken := "1"
	r, e := http.PostForm(`http://flights.ctrip.com/international/AjaxRequest/SearchFlights/AsyncSearchHandlerSOAII.ashx`,
		url.Values{"SearchMode": {SearchMode}, "condition": {condition}, "DisplayMode": {DisplayMode }, "SearchToken": {SearchToken}})
	if e != nil {
		log.Print(e)
	}
	defer r.Body.Close()

	b, e := ioutil.ReadAll(r.Body)
	//log.Print(len(b))
	//log.Print(string(b[0:1000]))
	var reply AsyncSearchHandlerSOAIIRelpy
	e =  json.Unmarshal(b, &reply )
	if e != nil {
		log.Print(e)
	}

	for _, value := range reply.RoundTripFlightList {
		log.Println(value.FareList[0].Price)
	}
}


func lowPrice() {
	// http://flights.ctrip.com/international/AjaxRequest/AsyncResult/LowPriceGeneric.ashx
	//SearchMode := `Search`
	//condition := `{"FlightWay":"D","SegmentList":[{"DCityCode":"BJS","ACityCode":"PNH","DCity":"Beijing|北京(BJS)|1|BJS|480","ACity":"Phnom Penh|金边(PNH)|303|PNH|420","DepartDate":"2017-9-19"},{"DCityCode":"PNH","ACityCode":"BJS","DCity":"Phnom Penh|金边(PNH)|303|PNH|480","ACity":"Beijing|北京(BJS)|1|BJS|420","DepartDate":"2017-9-21"}],"TransferCityID":0,"Quantity":1,"ClassGrade":"Y_S","TransNo":"c03cefe271a94f17af28d36c5ff3cb2b","SearchRandomKey":"","RecommendedFlightSwitch":1,"EngineFlightSwitch":1,"SearchKey":"05803CE80A253DC6BA940854EDC935501402DD68745CC4825CF2645227F1C4B5D51B0DD157BD07FB03A5D8E87E02EC7B4260BCD0B0CE7F8F","MultiPriceUnitSwitch":1,"TransferCitySwitch":false,"EngineScoreABTest":"B","SearchStrategySwitch":1,"MaxSearchCount":3,"TicketRemarkSwitch":1,"RowNum":"1500","TicketRemarkChannels":["GDS-WS","ZY-WS"],"AddSearchLogOneByOne":true,"TFAirlineQTE":"AA","IsWifiPackage":0,"SegmentVerifySwitch":false,"ComparePriceByAttributeSwitch":true,"IsOpenCFNoDirectRecommendYS":false,"IsDomesticIntlPUVersionSwitch":true,"DisplayBaggageSizeSwitch":true,"IsOpen24Refund":true,"IsOpenTransPU":true,"IsOpenVirtualFlight":false}`

}