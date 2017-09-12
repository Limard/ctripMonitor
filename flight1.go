package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GuoneiFlight struct {
	Fis []struct {
		Dcc  string  `json:"dcc"`  // 起飞城市
		Dbid int     `json:"dbid"` // 起飞城市的机场编号？
		Acc  string  `json:"acc"`  // 降落城市
		Abid int     `json:"abid"` // 降落城市的机场编号？
		Alc  string  `json:"alc"`  // 航空公司
		Fn   string  `json:"fn"`   // 航班名
		Dt   string  `json:"dt"`   // 起飞时间
		At   string  `json:"at"`   // 降落时间
		Tax  float64 `json:"tax"`  // tax??
		Scs  []struct {
			Chip struct {
				Price float64 // 价格??
			}
		}
	}

	Als map[string]string `json:"als"` // 航空公司简称-列表
	Ibc interface{}       // 保险相关
}

func GetGuonei() {
	// http://flights.ctrip.com/domesticsearch/search/SearchFirstRouteFlights?DCity1=nkg&ACity1=tao&SearchType=S&DDate1=2016-8-14&IsNearAirportRecommond=0

	// http://flights.ctrip.com/domesticsearch/search/SearchFirstRouteFlights?
	// DCity1=nkg &	<< 南京
	// ACity1=tao & << 青岛
	// SearchType=S &
	// DDate1=2016-8-14 &
	// IsNearAirportRecommond=0

	// BJS - beijing
	// BBK - mangu
	// 国内航线接口
	// http://flights.ctrip.com/domesticsearch/search/SearchFirstRouteFlights?DCity1=bjs&ACity1=bbk&SearchType=S&DDate1=2017-9-15&DDate2=2017-9-22IsNearAirportRecommond=0

	DCity1 := "nkg"
	ACity1 := "tao"
	DDate1 := "2017-9-13"
	ACity2 := "nkg"
	DDate2 := "2017-9-15"
	//searchUri := fmt.Sprintf(`http://flights.ctrip.com/domesticsearch/search/SearchFirstRouteFlights?DCity1=%s&ACity1=%s&SearchType=S&DDate1=%s&ACity2=%s&DDate2=%sIsNearAirportRecommond=0`,
	//	DCity1, ACity1, DDate1, ACity2, DDate2)
	searchUri := fmt.Sprintf(`http://flights.ctrip.com/domesticsearch/search/SearchFirstRouteFlights?DCity1=%s&ACity1=%s&SearchType=S&DDate1=%s&ACity2=%s&DDate2=%sIsNearAirportRecommond=0`,
		DCity1, ACity1, DDate1, ACity2, DDate2)
	fmt.Println(searchUri)

	rh, err := http.Get(`http://www.ctrip.com`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rh.Body.Close()
	fmt.Println(rh.Header)

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", searchUri, nil)
	reqest.Header.Add(`Accept`, `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`)
	reqest.Header.Add(`Accept-Language`, `zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4,ja;q=0.2`)
	reqest.Header.Add(`Cache-Control`, `max-age=0`)
	reqest.Header.Add(`Connection`, `keep-alive`)
	reqest.Header.Add(`Cookie`, `_abtest_userid=18056319-0cb8-4563-91d4-bcdfd8ee81f6; _bfa=1.1504861514644.22i34m.1.1504861514644.1504861514644.1.1; FlightIntl=Search=%5B%22Beijing%7C%E5%8C%97%E4%BA%AC(BJS)%7C1%7CBJS%7C480%22%2C%22Tokyo%7C%E4%B8%9C%E4%BA%AC(TYO)%7C228%7CTYO%7C540%22%2C%222017-10-13%22%2C%222017-10-19%22%5D; ASP.NET_SessionSvc=MTAuMTUuMTI4LjMxfDkwOTB8b3V5YW5nfGRlZmF1bHR8MTQ3MDkwNzQ3MDQwOA; page_time=1504861526226; _RF1=114.241.83.194; _RSG=kIa5HEa7VK9QYogOhrI398; _RGUID=951f1e43-9116-44ae-8243-4b5ad8ab289a; MKT_Pagesource=PC; _jzqco=%7C%7C%7C%7C%7C1.1033681901.1504861527295.1504861527295.1504861527296.1504861527295.1504861527296.0.0.0.1.1; __zpspc=9.1.1504861527.1504861527.1%234%7C%7C%7C%7C%7C%23; _bfi=p1%3D104003%26p2%3D0%26v1%3D1%26v2%3D0`)
	//reqest.Header.Add(`Cookie`, `_abtest_userid=18056319-0cb8-4563-91d4-bcdfd8ee81f6`)
	reqest.Header.Add(`Host`, `flights.ctrip.com`)
	reqest.Header.Add(`Upgrade-Insecure-Requests`, `1`)
	reqest.Header.Add(`User-Agent`, `Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36`)
	fmt.Println(reqest.Header)

	r, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	//gzipReader, err :=  gzip.NewReader(r.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer gzipReader.Close()

	//b, err := ioutil.ReadAll(gzipReader)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("%v\n", string(b))
	var f GuoneiFlight
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("size: %d\n", len(f.Fis))
	//fmt.Printf("%v\n", f.Fis)
	for _, fi := range f.Fis {
		fmt.Printf("%s: %s -> %s\n", fi.Fn, fi.Dt, fi.At)
		for _, sc := range fi.Scs {
			fmt.Printf("price: %v(%v, %v)\n", sc.Chip.Price+fi.Tax, sc.Chip.Price, fi.Tax)
		}
	}

	fmt.Println("finished")
}
