package coord

import (
	"math"
	"strconv"
	"strings"
)

func Direction(location_0 string, location string, angle float64) (direct string) {
	//以正北为y正方向，以正东为x正方向
	latitude_0_y, longitude_0_x := LocationStringToFloat(location_0)
	longitude_x, latitude_y := LocationStringToFloat(location)

	y_0_latitude, x_0_longitude := MillierConvertion(latitude_0_y, longitude_0_x)
	y_latitude, x_longitude := MillierConvertion(latitude_y, longitude_x)

	detla_y_latitude := y_latitude - y_0_latitude
	detla_x_longitude := x_longitude - x_0_longitude
	//转换成角度，以y轴负方向（正南方向）为零点，东南为正，西南为负
	direct_angle := math.Atan2(detla_y_latitude, detla_x_longitude) * 180 / math.Pi

	//log.Println(strconv.FormatFloat(direct_angle, 'E', -1, 32))

	switch {
	case direct_angle >= -1*angle && direct_angle < angle:
		direct = "南"
	case direct_angle >= angle && direct_angle < (90-angle):
		direct = "东南"
	case direct_angle >= (90-angle) && direct_angle < (90+angle):
		direct = "东"
	case direct_angle >= (90+angle) && direct_angle < (180-angle):
		direct = "东北"
	case direct_angle >= (180-angle) || direct_angle < -1*(180-angle):
		direct = "北"
	case direct_angle >= -1*(180-angle) && direct_angle < -1*(90+angle):
		direct = "西北"
	case direct_angle >= -1*(90+angle) && direct_angle < -1*(90-angle):
		direct = "西"
	case direct_angle >= -1*(90-angle) && direct_angle < -1*angle:
		direct = "西南"
	}
	return
}

func LocationStringToFloat(location string) (a float64, b float64) {
	locSlice := strings.Split(location, ",")
	a, _ = strconv.ParseFloat(locSlice[0], 64)
	b, _ = strconv.ParseFloat(locSlice[1], 64)
	return
}

//经纬度坐标转换到平面坐标
func MillierConvertion(lat float64, lon float64) (cor_x float64, cor_y float64) {
	var L, H, W, temp, mill, x, y float64
	L = 6381372 * math.Pi * 2 //地球周长
	W = L                     // 平面展开后，x轴等于周长
	H = L / 2                 // y轴约等于周长一半
	mill = 2.3                // 米勒投影中的一个常数，范围大约在正负2.3之间
	temp = math.Pi
	x = lon * temp / 180                           // 将经度从度数转换为弧度
	y = lat * temp / 180                           // 将纬度从度数转换为弧度
	y = 1.25 * math.Log(math.Tan(0.25*temp+0.4*y)) // 米勒投影的转换
	// 弧度转为实际距离
	cor_x = (W / 2) + (W/(2*math.Pi))*x
	cor_y = (H / 2) - (H/(2*mill))*y
	return
}

type Cor struct {
	X float64
	Y float64
}

type Loc struct {
	Lat float64
	Lon float64
}

func LocDistance(one, two Loc) (res float64) {
	var corOne, corTwo Cor
	corOne.X, corOne.Y = MillierConvertion(one.Lat, one.Lon)
	corTwo.X, corTwo.Y = MillierConvertion(two.Lat, two.Lon)
	a := math.Pow(corOne.X-corTwo.X, 2)
	a += math.Pow(corOne.Y-corTwo.Y, 2)
	return math.Sqrt(a)
}
