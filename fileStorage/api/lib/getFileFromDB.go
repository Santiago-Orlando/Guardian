package lib

import (
	"Guardian/fileStorage/api/database"
)

func GetFileFromDB(userID string, filename string) string {

	db := database.GetConnection()

	var serverName string

	query := "SELECT server_name FROM files WHERE user_id = $1 and name = $2"

	err := db.QueryRow(query, userID, filename).Scan(&serverName)
	if err != nil {
		ErrorHandler(err, "database")
	}

	return serverName
}