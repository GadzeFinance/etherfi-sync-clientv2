package utils

import (
	"database/sql"
)

func GetIDCount(db *sql.DB, bidID string) (int, error) {

	query := "SELECT COUNT(*) FROM winning_bids WHERE id = ?"
	var count int
	err := db.QueryRow(query, bidID).Scan(&count)

	return count, err
}

func CreateTable(db *sql.DB) error {
	// Create the table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS winning_bids (
		id STRING PRIMARY KEY,
		pubkey TEXT,
		password TEXT,
		nodeAddress TEXT,
		syncStatus TEXT DEFAULT 'PENDING',
		keystore TEXT
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

/*
func UpdateRowStatus(db *sql.DB, bidId string, status string) (error) {
	query := `
	UPDATE winning_bids
	SET syncStatus = $2
	WHERE id = $1;`

	_, err := db.Exec(query, bidId, status)
	if err != nil {
		return err
	}

	return nil
}

func GetRowsByStatus(db *sql.DB, status string) ([]schemas.DisplayBid, error) {
	rows, err := db.Query("SELECT id, pubkey FROM winning_bids WHERE syncStatus=$1", status)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var bids []schemas.DisplayBid
	for rows.Next() {
    var bid schemas.DisplayBid
    err = rows.Scan(&bid.Id, &bid.Pubkey)
    if err != nil {
      return nil, err
    }
    bids = append(bids, bid)
  }
	return bids, nil
}

func GetBid(db *sql.DB, id string) (schemas.TableBid, error) {

	var data schemas.TableBid
	row := db.QueryRow("SELECT * FROM winning_bids WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Pubkey, &data.Password, &data.NodeAddress, &data.SyncStatus, &data.Keystore)
	if err != nil {
		return schemas.TableBid{}, err
	}

	return data, nil
}
*/
