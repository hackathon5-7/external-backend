package models

type Billboard struct {
	BillboardId string `json:"-" db:"billboard_id"`
	Lat         string `json:"lat" db:"lat"`
	Lon         string `json:"lon" db:"lon"`
	Azimuth     string `json:"azimuth" db:"azimuth"`
}

type Request struct {
	RequestId     int    `json:"-" db:"request_id"`
	Gender        string `json:"gender" db:"gender"`
	AgeFrom       int    `json:"age_from" db:"age_from"`
	AgeTo         int    `json:"age_to" db:"age_to"`
	IncomeA       bool   `json:"income_a" db:"income_a"`
	IncomeB       bool   `json:"income_b" db:"income_b"`
	IncomeC       bool   `json:"income_c" db:"income_c"`
	NameBillboard string `json:"name_billboard" db:"name_billboard"`
	UserId        int    `json:"user_id" db:"user_id"`
	BillboardId   int    `json:"billboard_id" db:"billboard_id"`
}
