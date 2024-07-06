package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
	"encoding/json"
	"os"
)

type TargetAudience struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	AgeFrom int    `json:"ageFrom"`
	AgeTo   int    `json:"ageTo"`
	Income  string `json:"income"`
}

type DataObject struct {
	Hash           string             `json:"hash"`
	TargetAudience TargetAudience     `json:"targetAudience"`
	Points         []models.Billboard `json:"points"`
	Value          float64            `json:"value"`
}

type DatasetParserJSONService struct {
	reposBillboard repository.StorageBillboard
	reposSectors   repository.StorageSectors
}

func NewDatasetParserJSONService(reposBillboard repository.StorageBillboard, reposSectors repository.StorageSectors) *DatasetParserJSONService {
	return &DatasetParserJSONService{
		reposBillboard: reposBillboard,
		reposSectors:   reposSectors,
	}
}

// GetBillboards reads a JSON file from the specified path and parses it
// to extract the billboard points. It then adds each point to the database.
//
// pathFile: The path to the JSON file.
// Returns: An error if any occurred while parsing the JSON file or adding the billboard points to the database.
func (s DatasetParserJSONService) GetBillboards(pathFile string) error {
	// Read the JSON file from the specified path.
	bytes, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	// Create a slice to store the parsed data.
	var data []DataObject

	// Parse the JSON bytes into the data slice.
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	// Create a map to store the unique points.
	var pointsHash = make(map[models.Billboard]bool)

	// Iterate over each data object and add its points to the hash map.
	for _, elem := range data {
		// Iterate over each point in the data object.
		for _, point := range elem.Points {
			// Add the point to the hash map.
			pointsHash[point] = true
		}
	}

	// Iterate over the hash map and add each point to the slice.
	for key := range pointsHash {
		// Get the sector for the coordinate and assign it to the point.
		key.SectorId, err = s.reposSectors.GetSectorForCoordinate(key)
		if err != nil {
			return err
		}

		// Add the billboard point to the database.
		_, err := s.reposBillboard.AddBillboard(key)
		if err != nil {
			return err
		}
	}

	return nil
}
