package parser

import "github.com/s51ds/ctydat/adif"

type CtyDta struct {
	CountryName   string //Country Name
	PrimaryPrefix string
	AliasPrefix   string             //Primary or Alias DXCC Prefix without optional * indicator
	Continent     adif.ContinentEnum //2-letter Continent abbreviation
	CqZone        adif.CqzoneEnum    //CQ Zone
	ItuZone       adif.ItuzoneEnum   //ITU Zone
	LatLon        LatLonDeg          //Latitude in degrees, + for North; Longitude in degrees, + for West
	TimeOffset    string             //Local time offset from GMT
}
