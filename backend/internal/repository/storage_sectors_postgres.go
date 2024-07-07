package repository

import (
	"app/backend/internal/models"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type StorageSectorsPostgres struct {
	db *sqlx.DB
}

func NewStorageSectorsPostgres(db *sqlx.DB) *StorageSectorsPostgres {
	return &StorageSectorsPostgres{db: db}
}

// AddSector adds a new sector to the database and returns the ID of the inserted sector.
//
// Parameters:
// - point: A Sector struct representing the coordinates of the sector.
//
// Returns:
// - A int64 representing the ID of the inserted sector.
// - An error if there was an issue with adding the sector.
func (s StorageSectorsPostgres) AddSector(point models.Sector) (int64, error) {
	// Create a query to insert a new sector into the table.
	query := fmt.Sprintf("INSERT INTO %s (x_max, x_min, y_max, y_min) VALUES ($1, $2, $3, $4) RETURNING sector_id", tableSectors)

	// Create a variable to hold the ID of the inserted sector.
	var id int64

	// Scan the result into the ID variable.
	if err := s.db.QueryRow(query, point.X_max, point.X_min, point.Y_max, point.Y_min).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetSectorForCoordinate retrieves the sector ID for a given coordinate.
//
// This function executes a SQL query to find the sector ID for a given coordinate.
// The coordinate is provided as a Billboard struct with latitude and longitude.
//
// Parameters:
// - coordinate: A Billboard struct containing the longitude and latitude of the coordinate.
//
// Returns:
// - The ID of the sector that contains the coordinate, or 0 if no sector is found.
// - An error if there was an issue with executing the query.
func (s StorageSectorsPostgres) GetSectorForCoordinate(coordinate models.Billboard) (int, error) {
	// Create a query to select the sector ID from the sectors table
	// where the coordinate falls within the sector's coordinates.
	query := fmt.Sprintf(`
		SELECT sector_id 
		FROM %s 
		WHERE x_min <= $1 AND x_max >= $1 AND y_min <= $2 AND y_max >= $2
	`, tableSectors)

	// Execute the query and get the result set.
	lon, err := strconv.ParseFloat(coordinate.Lon, 64)
	if err != nil {
		return 0, err
	}

	lat, err := strconv.ParseFloat(coordinate.Lat, 64)
	if err != nil {
		return 0, err
	}

	rows, err := s.db.Query(query, lat, lon)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// If there is a row returned, scan the sector ID into the variable.
	var id int
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}

	// If no sector is found, return 0.
	return 0, nil
}

// GetSizeStorageSectors retrieves the number of sectors in the database.
//
// Executes a SQL query that counts the number of rows in the sectors table.
// The result is returned as an integer.
//
// Returns:
// - The number of sectors in the database.
// - An error if there was any issue with executing the query.
func (s StorageSectorsPostgres) GetSizeStorageSectors() (int, error) {
	// Create a query to count the number of rows in the sectors table.
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableSectors)

	// Execute the query and get the result set.
	var count int
	// Scan the result into the count variable.
	if err := s.db.QueryRow(query).Scan(&count); err != nil {
		// Return the error if there was an issue with executing the query.
		return 0, err
	}

	// Return the count of sectors and any error that occurred.
	return count, nil
}
