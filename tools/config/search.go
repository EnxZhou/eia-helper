package config

import "github.com/spf13/viper"

type Search struct {
	Keywords string
	Types    string
	Location string
	Radius   string
	SortRule string
	Region     string
	PageSize     string
	PageNum   string
	ShowFields string
}

func InitSearch(cfg *viper.Viper) *Search {
	return &Search{
		Keywords: cfg.GetString("keywords"),
		Types:    cfg.GetString("types"),
		Location: cfg.GetString("location"),
		Radius:   cfg.GetString("radius"),
		SortRule: cfg.GetString("sortrule"),
		Region:     cfg.GetString("region"),
		PageSize:   cfg.GetString("page_size"),
		PageNum:     cfg.GetString("page_num"),
		ShowFields:     cfg.GetString("show_fields"),
	}
}

var SearchConfig = new(Search)
