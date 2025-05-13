package ph3

import (
	"strconv"

	"github.com/uber/h3-go/v4"
)

// https://observablehq.com/@nrabinowitz/h3-radius-lookup
// https://h3geo.org/
// https://h3geo.org/docs/core-library/restable/

// Res	Average edge length (Km/miles) center2center
// 4    26.07175968/16.200235384       28.05963078
// 5    9.854090990/6.1230463725       10.60542743
// 6    3.724532667/2.3143165878       4.008513915

func GetH3Cell(lat, lng float64) (h3.Cell, error) {
	return h3.LatLngToCell(h3.NewLatLng(lat, lng), 5)
}

func GetH3CellString(lat, lng float64) (string, error) {
	h3Cell, err := GetH3Cell(lat, lng)
	if err != nil {
		return "", nil
	}
	return strconv.FormatInt(int64(h3Cell), 16), nil
}

func GetDiskByOriginH3Cell(origin int64) ([]int64, error) {
	cells, err := h3.GridDisk(h3.Cell(origin), 1)
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(cells))
	for i, cell := range cells {
		ret[i] = int64(cell)
	}
	return ret, nil
}

func getDiskByLatLngRadius(lat, lng float64) ([]int64, error) {
	origin, err := GetH3Cell(lat, lng)
	if err != nil {
		return nil, err
	}
	cells, err := h3.GridDisk(origin, 1)
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(cells))
	for i, cell := range cells {
		ret[i] = int64(cell)
	}
	return ret, nil
}
