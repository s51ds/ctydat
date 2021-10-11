package parser

import "fmt"

type LatLonDeg struct {
	Lat, Lon float64
}

func (a *LatLonDeg) String() string {
	//	return fmt.Sprintf("Lat=%.4f, Lon=%.4f", a.Lat, a.Lon)
	return fmt.Sprintf("{%.6f %.6f}", a.Lat, a.Lon)
}
