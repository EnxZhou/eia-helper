package amap

import (
	"eia-helper/model"
	"eia-helper/tools/config"
	"eia-helper/tools/coord"
	"eia-helper/tools/request"
	"fmt"
	"log"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

func url() string {
	key := config.AmapKey
	cfg := config.SearchConfig
	url := "https://restapi.amap.com/v3/place/around?"
	url += "key=" + key
	url += "&keywords=" + cfg.Keywords
	url += "&types=" + cfg.Types
	url += "&location=" + cfg.Location
	url += "&radius=" + cfg.Radius
	url += "&sortrule=" + cfg.SortRule
	url += "&page_size=" + cfg.PageSize
	url += "&page_num=" + cfg.PageNum
	url += "&show_fields=" + cfg.ShowFields
	return url
}

func centerLoc() (float64, float64) {
	cfg := config.SearchConfig
	location := cfg.Location
	return coord.LocationStringToFloat(location)
}

func GetPoiList() (res model.PoiInfo, err error) {
	remoteUrl := url()
	var resByte []byte
	resByte, err = request.Get(remoteUrl)
	if err != nil {
		err = errors.Wrap(err, "request fail")
		return
	}
	info := jsonString(resByte, "info")
	if info != "OK" {
		err = errors.Errorf("api response fail, info: %s", info)
		return
	}
	count := jsonString(resByte, "count")
	log.Printf("get count: %s result", count)
	poiList := make([]model.Poi, 0)
	_, err = jsonparser.ArrayEach(resByte,
		func(value []byte,
			dataType jsonparser.ValueType,
			offset int,
			err error) {
			tmpBus := model.Business{
				Tel: jsonString(value, "business", "tel"),
			}
			tmp := model.Poi{
				Name:     jsonString(value, "name"),
				Id:       jsonString(value, "id"),
				Location: jsonString(value, "location"),
				Type:     jsonString(value, "type"),
				Typecode: jsonString(value, "typecode"),
				Pname:    jsonString(value, "pname"),
				Cityname: jsonString(value, "cityname"),
				Adname:   jsonString(value, "adname"),
				Address:  jsonString(value, "address"),
				Pcode:    jsonString(value, "pcode"),
				Adcode:   jsonString(value, "adcode"),
				Citycode: jsonString(value, "citycode"),
				Business: tmpBus,
			}
			poiList = append(poiList, tmp)
		}, "pois")
	if err != nil {
		err = errors.Wrap(err, "json marshal fail")
		return
	}
	res.Info = info
	res.Count = count
	res.Pois = poiList
	res.CenterCorX, res.CenterCorY = centerLoc()
	fmt.Println(res)
	return
}

func jsonString(data []byte, key ...string) string {
	tmp, err := jsonparser.GetString(data, key...)
	if err != nil {
		log.Printf("json get %s fail\n", key)
	}
	return tmp
}
