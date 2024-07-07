package service

import (
	"app/backend/internal/models"
	"app/backend/internal/repository"
	"encoding/json"
	"os"
)

type SectorParserJSONService struct {
	repos repository.StorageSectors
}

func NewSectorParserJSONService(repos repository.StorageSectors) *SectorParserJSONService {
	return &SectorParserJSONService{
		repos: repos,
	}
}

// GetSectors reads a JSON file from the specified path and parses it
// to extract the sectors. It then adds each sector to the database.
//
// pathFile: The path to the JSON file.
// Returns: An error if any occurred while parsing the JSON file or adding the sectors to the database.
func (s SectorParserJSONService) GetSectors(pathFile string) error {
	// Read the JSON file from the specified path.
	bytes, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	// Create a slice to store the parsed data.
	var data []models.Sector

	// Parse the JSON bytes into the data slice.
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	// Iterate over each sector and add it to the database.
	for _, elem := range data {
		// Add the sector to the database.
		_, err := s.repos.AddSector(elem)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSectorsArray reads a JSON file from the specified path and parses it
// to extract the sectors. It returns the parsed data as a slice of models.Sector.
//
// pathFile: The path to the JSON file.
// Returns:
// - A slice of models.Sector containing the parsed data.
// - An error if any occurred while parsing the JSON file.
func (s SectorParserJSONService) GetSectorsArray(pathFile string) ([]models.Sector, error) {
	// Read the JSON file from the specified path.
	bytes, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}

	// Create a slice to store the parsed data.
	var data []models.Sector

	// Parse the JSON bytes into the data slice.
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
