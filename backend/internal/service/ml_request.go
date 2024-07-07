package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RecomendationOutput struct {
	SectorId int                `json:"sector_id"`
	Value    int                `json:"value" db:"value"`
	Points   []models.Billboard `json:"points"`
}

type MLRequestOutput struct {
	Gender              string `json:"gender"`
	AgeFrom             int    `json:"age_from"`
	AgeTo               int    `json:"age_to"`
	Income              string `json:"income"`
	CountPointOnSegment []int  `json:"count_point_on_segment"`
}

type MLRequestInput struct {
	Value int `json:"value" binding:"required"`
}

type RecomendationInput struct {
	Filters Filters `json:"filters" binding:"required"`
}

type Filters struct {
	AgeFrom  int                `json:"ageFrom" binding:"required"`
	AgeTo    int                `json:"ageTo" binding:"required"`
	Gender   string             `json:"gender" binding:"required"`
	Income   IncomeInputHandler `json:"income" binding:"required"`
	Quantity int                `json:"quantity" binding:"required"`
}

type IncomeInputHandler struct {
	A bool `json:"a" binding:"required"`
	B bool `json:"b" binding:"required"`
	C bool `json:"c" binding:"required"`
}

type MLRequestService struct {
	reposSectors   repository.StorageSectors
	reposBillboard repository.StorageBillboard
}

func NewMLRequestService(reposSectors repository.StorageSectors, reposBillboard repository.StorageBillboard) *MLRequestService {
	return &MLRequestService{
		reposSectors:   reposSectors,
		reposBillboard: reposBillboard,
	}
}

// GetRecomendation retrieves sector recommendations based on the input filters.
//
// The function takes in a RecomendationInput struct as a parameter and returns a slice of RecomendationOutput structs and an error.
// It iterates over each sector, retrieves the billboards within the sector, and makes an HTTP POST request to the "denis" endpoint.
// The response from the POST request is unmarshaled into an MLRequestInput struct, and the sector ID, the value from the response,
// and the retrieved billboards are added to the RecomendationOutput struct.
func (s MLRequestService) GetRecomendation(input RecomendationInput) ([]RecomendationOutput, error) {
	var data []RecomendationOutput

	// Get the total number of sectors
	countSectors, err := s.reposSectors.GetSizeStorageSectors()
	if err != nil {
		return data, err
	}

	// Iterate over each sector
	for i := 1; i <= countSectors; i++ {
		// Get the billboards within the sector
		countBilboard, err := s.reposBillboard.GetBillboardsBySectorId(i, input.Filters.Quantity)
		if err != nil {
			return nil, err
		}

		// Create a slice to hold the count of billboards for each segment
		countPointOnSegment := make([]int, 400)
		countPointOnSegment[i-1] = len(countBilboard)

		// Create a string representing the income values
		var income string
		if input.Filters.Income.A {
			income += "a"
		}
		if input.Filters.Income.B {
			income += "b"
		}
		if input.Filters.Income.C {
			income += "c"
		}

		// Create a JSON object with the sector information
		dataJson := MLRequestOutput{
			Gender:              input.Filters.Gender,
			AgeFrom:             input.Filters.AgeFrom,
			AgeTo:               input.Filters.AgeTo,
			Income:              income,
			CountPointOnSegment: countPointOnSegment,
		}

		// Marshal the JSON object into bytes
		jsonData, err := json.Marshal(dataJson)
		if err != nil {
			return nil, err
		}

		// Make an HTTP POST request to the "denis" endpoint with the JSON data
		resp, err := http.Post("denis", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Unmarshal the response body into an MLRequestInput struct
		var value MLRequestInput
		err = json.Unmarshal(body, &value)
		if err != nil {
			return nil, err
		}

		// Get the billboards within the sector again
		bilboards, err := s.reposBillboard.GetBillboardsBySectorId(i, input.Filters.Quantity)
		if err != nil {
			return nil, err
		}

		// Add the sector ID, the value from the response, and the retrieved billboards to the RecomendationOutput struct
		data = append(data, RecomendationOutput{
			SectorId: i,
			Value:    value.Value,
			Points:   bilboards,
		})

	}

	return data, nil
}
