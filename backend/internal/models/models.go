package models

type Billboard struct {
	BillboardId string `json:"-"`
	Lat         string `json:"lat" db:"lat"`
	Lon         string `json:"lon" db:"lon"`
	Azimuth     string `json:"azimuth" db:"azimuth"`
}

type Request struct {
	RequestId int    `json:"request_id"`
	Gender    string `json:"gender"`
	AgeFrom   int    `json:"age_from"`
	AgeTo     int    `json:"age_to"`
	Income    string `json:"income"`
}
