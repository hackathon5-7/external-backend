package models

type Billboard struct {
	BillboardId string `json:"billboard_id" db:"billboard_id"`
	SectorId    int    `json:"sector_id" db:"sector_id"`
	Lat         string `json:"lat" db:"lat"`
	Lon         string `json:"lon" db:"lon"`
	Azimuth     int    `json:"azimuth" db:"azimuth"`
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
}

type RequestBillboard struct {
	RequestsBillboardsId int    `json:"-" db:"requests_billboards_id"`
	RequestId            int    `json:"request_id" db:"request_id"`
	BillboardId          int    `json:"billboard_id" db:"billboard_id"`
	Value                string `json:"value" db:"value"`
}

type Sector struct {
	SectorsId int    `json:"-" db:"sector_id"`
	X_max     string `json:"x_max" db:"x_max"`
	X_min     string `json:"x_min" db:"x_min"`
	Y_max     string `json:"y_max" db:"y_max"`
	Y_min     string `json:"y_min" db:"y_min"`
}
