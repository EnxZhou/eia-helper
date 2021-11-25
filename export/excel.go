package export

import (
	"eia-helper/model"
	"eia-helper/tools/config"
	"eia-helper/tools/coord"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func WriteXlsx(fileName string, date model.PoiInfo) {
	cfg := config.SearchConfig
	xlsx := excelize.NewFile()

	centerLoc := coord.Loc{}
	centerLoc.Lon, centerLoc.Lat = coord.GCJ02toWGS84(
		coord.LocationStringToFloat(cfg.Location))
	for i := 0; i < len(date.Pois); i++ {
		direct := coord.Direction(cfg.Location, date.Pois[i].Location, 10)

		//GCJ02toWGS84坐标转换为84坐标，GCJ02坐标系：即火星坐标系，WGS84坐标系经加密后的坐标系。Google Maps，高德在用。
		longitude, latitude := coord.GCJ02toWGS84(coord.LocationStringToFloat(date.Pois[i].Location))
		//将米转换成千米
		currentLoc := coord.Loc{
			Lon: longitude,
			Lat: latitude,
		}
		// fmt.Printf("centerLoc:%v\t currentLoc:%v\n", currentLoc, currentLoc)
		distance_Float := coord.LocDistance(centerLoc, currentLoc)
		distance_Float = distance_Float / 1000
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i+1), strconv.Itoa(i+1))
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i+1), date.Pois[i].Name)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(i+1), direct)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(i+1), distance_Float)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(i+1), longitude)
		xlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(i+1), latitude)
		xlsx.SetCellValue("Sheet1", "I"+strconv.Itoa(i+1), date.Pois[i].Tel)
	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(1)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./" + fileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
