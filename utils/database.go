package utils

import "database/sql"

func GetIDCount(db *sql.DB, bidID string) (int, error) {

	query := "SELECT COUNT(*) FROM winning_bids WHERE id = ?"
	var count int
	err := db.QueryRow(query, bidID).Scan(&count)

	return count, err
}