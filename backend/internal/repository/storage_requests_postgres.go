package repository

import (
	"app/backend/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type StorageRequestsToTheModelPostgres struct {
	db *sqlx.DB
}

func NewStorageRequestsToTheModelPostgres(db *sqlx.DB) *StorageRequestsToTheModelPostgres {
	return &StorageRequestsToTheModelPostgres{
		db: db,
	}
}

// AddRequest adds a request to the database.
//
// Parameters:
// - request: A Request struct containing the age range, gender, income, name of the billboard, user ID, and billboard ID.
//
// Returns:
// - An int64 representing the ID of the request.
// - An error if there was any issue.
func (r StorageRequestsToTheModelPostgres) AddRequest(request models.Request) (int64, error) {
	// Create a query to insert the request into the table and return its ID.
	query := fmt.Sprintf(`
		INSERT INTO %s (
			age_from, age_to, gender, income_a, income_b, income_c, name_billboard
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		) RETURNING request_id`, tableRequest)

	// Execute the query and scan the result into the ID variable.
	var id int64
	rows, err := r.db.Query(query, request.AgeFrom, request.AgeTo, request.Gender, request.IncomeA, request.IncomeB, request.IncomeC, request.NameBillboard)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// If there is a row returned, scan the ID into the variable.
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

// DeleteRequest deletes a request from the database by its ID.
//
// Parameters:
// - RequestId: An int64 representing the ID of the request.
//
// Returns:
// - An error if there was any issue with deleting the request.
func (r StorageRequestsToTheModelPostgres) DeleteRequest(RequestId int64) error {
	// Create a query to delete the request with the given ID.
	query := fmt.Sprintf("DELETE FROM %s WHERE request_id = $1", tableRequest)

	// Execute the query and ignore the result since we only care about the error.
	// The query will delete the row from the table that has the same ID as the given RequestId.
	_, err := r.db.Exec(query, RequestId)

	// Return any error that occurred during the deletion.
	return err
}
