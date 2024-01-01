package main

import (
	"fmt"

	"github.com/aceagles/Go-Practice/csvParse/plantMap"
)

func main() {
	mp, err := plantMap.LoadMap("./db_BLF.csv")
	if err != nil {
		return
	}
	fmt.Println(mp.GetPlantByBase("Inverter"))
	fmt.Println(mp.GetPlantByBase("Inverter").ByPrefix("31_"))
	fmt.Println(mp.GetPlantByBase("Inverter").ByPrefix("38_"))

}
