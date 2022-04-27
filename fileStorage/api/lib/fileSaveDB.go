package lib

import (
	"Guardian/fileStorage/api/database"
	"database/sql"
	"path/filepath"
	"strings"
)

func FileSaveDB(name string, hash string, serverName string, userID string) error {

	db := database.GetConnection()

	serverName = strings.Replace(serverName, filepath.Ext(serverName), ".gz", -1)

	query := "INSERT INTO files (name, hash, server_name, user_id) VALUES ($1, $2, $3, $4) "
	
	err := db.QueryRow(query, name, hash, serverName, userID).Err()
	if err == sql.ErrNoRows {
		return nil
	}
	
	return err
}