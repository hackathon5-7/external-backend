package repository

import (
	"app/backend/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type StorageBillboardPostgres struct {
	db *sqlx.DB
}

func NewStorageBillboardPostgres(db *sqlx.DB) *StorageBillboardPostgres {
	return &StorageBillboardPostgres{db: db}
}

// AddBillboard adds a billboard to the database.
//
// point: A Billboard struct containing the latitude, longitude, azimuth, and sector ID of the billboard.
// Returns: An int64 representing the ID of the billboard and an error if there was any issue.
// If the insertion is successful, 0 is returned.
func (r StorageBillboardPostgres) AddBillboard(point models.Billboard) (int64, error) {
	// Create a query to insert the billboard into the table and return its ID.
	// The query inserts the latitude, longitude, azimuth, and sector ID of the billboard into the billboards table.
	// The ID of the inserted billboard is returned.
	query := fmt.Sprintf("INSERT INTO %s (lat, lon, azimuth, sector_id) VALUES ($1, $2, $3, $4) RETURNING billboard_id", tableBillboards)

	// Execute the query and get the result row.
	row := r.db.QueryRow(query, point.Lat, point.Lon, point.Azimuth, point.SectorId)

	// Create a variable to hold the ID of the inserted billboard.
	var id int64

	// Scan the result row into the ID variable.
	// If there is an error, return 0 and the error.

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	// Return the ID of the inserted billboard and no error.
	return id, nil
}

// GetBillboardById returns a Billboard struct from the database that matches the given ID.
//
// id: An int64 representing the ID of the billboard.
// Returns: A Billboard struct and an error if there was any issue.
func (r StorageBillboardPostgres) GetBillboardById(id int64) (models.Billboard, error) {
	// Create a query to select the billboard from the table that matches the given ID.
	query := fmt.Sprintf("SELECT lat, lon, azimuth FROM %s WHERE billboard_id = $1", tableBillboards)
	row := r.db.QueryRow(query, id)

	// Create a variable to hold the billboard.
	var point models.Billboard
	// Scan the result into the billboard variable.
	if err := row.Scan(&point.Lat, &point.Lon, &point.Azimuth); err != nil {
		return models.Billboard{}, err
	}

	return point, nil
}

// DeleteBillboardById deletes a billboard from the database by its ID.
//
// id: An int64 representing the ID of the billboard.
// Returns: An error if there was any issue with deleting the billboard.
func (r StorageBillboardPostgres) DeleteBillboardById(id int64) error {
	// Create a query to delete the billboard with the given ID.
	query := fmt.Sprintf("DELETE FROM %s WHERE billboard_id = $1", tableBillboards)

	// Execute the query and ignore the result since we only care about the error.
	// The query will delete the row from the table that has the same ID as the given id.
	_, err := r.db.Exec(query, id)

	// Return any error that occurred during the deletion.
	return err
}

// GetAllBillboards retrieves all billboards from the database.
//
// Returns a slice of Billboard structs and an error if there was any issue.
func (r StorageBillboardPostgres) GetAllBillboards() ([]models.Billboard, error) {
	// Create a query to select all billboards from the table.
	query := fmt.Sprintf("SELECT * FROM %s", tableBillboards)

	// Execute the query and get the result set.
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the retrieved billboards.
	var billboards []models.Billboard

	// Loop through the result set and scan each row into a Billboard struct.
	for rows.Next() {
		var billboard models.Billboard
		if err := rows.Scan(&billboard.BillboardId, &billboard.Lat, &billboard.Lon, &billboard.Azimuth); err != nil {
			return nil, err
		}
		billboards = append(billboards, billboard)
	}

	// Return the retrieved billboards and any error that occurred.
	return billboards, nil
}

// GetSizeStorage retrieves the number of billboards in the database.
//
// It executes a SQL query that counts the number of rows in the billboards table.
// The result is returned as an integer.
//
// Returns:
// - The number of billboards in the database.
// - An error if there was any issue with executing the query.
func (r StorageBillboardPostgres) GetSizeStorageBillboards() (int, error) {
	// Create a query to count the number of rows in the billboards table.
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableBillboards)

	// Execute the query and get the result set.
	row := r.db.QueryRow(query)

	// Create a variable to hold the number of billboards.
	var count int

	// Scan the result into the count variable.
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	// Return the number of billboards and any error that occurred.
	return count, nil
}

// GetAllBillboardsBySectorId retrieves all billboards from the database that are within the specified sector.
//
// sectorId: An int representing the ID of the sector.
// Returns: A slice of Billboard structs and an error if there was any issue.
func (r StorageBillboardPostgres) GetBillboardsBySectorId(sectorId int, limit int) ([]models.Billboard, error) {
	// Create a query to select all billboards from the table that are within the specified sector.
	query := fmt.Sprintf("SELECT * FROM %s WHERE sector_id = $1 LIMIT $2", tableBillboards)

	// Execute the query and get the result set.
	rows, err := r.db.Query(query, sectorId, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the retrieved billboards.
	var billboards []models.Billboard

	// Loop through the result set and scan each row into a Billboard struct.
	for rows.Next() {
		var billboard models.Billboard
		if err := rows.Scan(&billboard.BillboardId, &billboard.SectorId, &billboard.Lat, &billboard.Lon, &billboard.Azimuth); err != nil {
			return nil, err
		}
		billboards = append(billboards, billboard)
	}

	// Return the retrieved billboards and any error that occurred.
	return billboards, nil
}
