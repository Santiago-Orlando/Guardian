package lib

import (
	"Guardian/fileStorage/api/database"
	m "Guardian/fileStorage/api/models"
	"database/sql"
	"fmt"
	"net/http"
)

func DuplicatePreventor(sha string, userID string) ( m.DatabaseFile, error ) {

	db := database.GetConnection()

	file := m.DatabaseFile{}

	query := "SELECT hash, server_name, user_id FROM files WHERE hash = $1"
	
	err := db.QueryRow(query, sha).Scan(&file.Sha, &file.NewFileName, &file.UserID)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return file, err
	}

	if file.UserID == userID && file.Sha == sha {
		return file, http.ErrBodyNotAllowed
	}

	return file, nil
}