package factory

import (
	"abstract-factory/app/inter"
	"abstract-factory/app/service/brand"
)

var sportMap map[string]inter.ISportsFactory

const (
	Brand_Adidas = "adidas"
	Brand_Nike = "nike"
)

func init(){
	sportMap = make(map[string]inter.ISportsFactory,0)
	RegSports(Brand_Nike, &brand.Nike{})
	RegSports(Brand_Adidas, &brand.Adidas{})
}

func SportsManager(brand string) inter.ISportsFactory{
	return sportMap[brand]
}

func RegSports(brand string, sport inter.ISportsFactory){
	sportMap[brand] = sport
}