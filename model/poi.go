package model

type PoiInfo struct {
	Status     string
	Count      string
	Info       string
	Infocode   string
	Pois       []Poi
	CenterCorX float64
	CenterCorY float64
}

type Poi struct {
	Name     string
	Id       string
	Location string
	Type     string
	Typecode string
	Pname    string
	Cityname string
	Adname   string
	Address  string
	Pcode    string
	Adcode   string
	Citycode string
	Business
}

type Business struct{
	Tel string
}