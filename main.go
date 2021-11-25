package main

import (
	"eia-helper/amap"
	"eia-helper/export"
	"eia-helper/tools/config"
	"log"
)

func init() {
	config.ConfigSetup("./config/settings.yml")
}

func main() {
	poiInfo, err := amap.GetPoiList()
	if err != nil {
		log.Fatal(err)
	}
	export.WriteXlsx("result", poiInfo)
}
