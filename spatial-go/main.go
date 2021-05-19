package main

import (
	"fmt"

	"github.com/spatial-go/geoos/encoding/wkt"
)

func main() {
	var polygonWktStr = "POLYGON ((30 10, 10 20, 20 40, 40 40, 30 10))"
	getGeom, err := wkt.UnmarshalString(polygonWktStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getGeom)
}

// [[[30 10] [10 20] [20 40] [40 40] [30 10]]]
