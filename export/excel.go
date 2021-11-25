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
	centerLoc.Lat, centerLoc.Lon = coord.GCJ02toWGS84(
		coord.LocationStringToFloat(cfg.Location))
	for index := 0; index < len(date.Pois); index++ {
		direct := coord.Direction(cfg.Location, date.Pois[index].Location, 10)
		//log.Println(direct)

		//GCJ02toWGS84坐标转换为84坐标，GCJ02坐标系：即火星坐标系，WGS84坐标系经加密后的坐标系。Google Maps，高德在用。
		longitude, latitude := coord.GCJ02toWGS84(coord.LocationStringToFloat(date.Pois[index].Location))
		//将米转换成千米
		currentLoc := coord.Loc{
			Lat: longitude,
			Lon: latitude,
		}
		distance_Float := coord.LocDistance(centerLoc, currentLoc)
		distance_Float = distance_Float / 1000
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(index+1), strconv.Itoa(index+1))
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(index+1), date.Pois[index].Name)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(index+1), direct)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(index+1), distance_Float)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(index+1), longitude)
		xlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(index+1), latitude)
		xlsx.SetCellValue("Sheet1", "I"+strconv.Itoa(index+1), date.Pois[index].Tel)
	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(1)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./" + fileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
