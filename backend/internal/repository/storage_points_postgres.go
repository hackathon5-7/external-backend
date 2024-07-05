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
// point: A Billboard struct containing the latitude, longitude, and azimuth of the billboard.
// Returns: An int64 representing the ID of the billboard and an error if there was any issue.
func (r StorageBillboardPostgres) AddBillboard(point models.Billboard) (int64, error) {
	// Create a query to insert the billboard into the table and return its ID.
	query := fmt.Sprintf("INSERT INTO %s (lat, lon, azimuth) VALUES ($1, $2, $3) RETURNING billboard_id", tableBillboards)
	row := r.db.QueryRow(query, point.Lat, point.Lon, point.Azimuth)

	// Create a variable to hold the ID of the inserted billboard.
	var id int64
	// Scan the result into the ID variable.
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

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
