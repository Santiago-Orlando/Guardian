package services

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Guardian/fileStorage/api/database"
	"Guardian/fileStorage/api/lib"
	m "Guardian/fileStorage/api/models"

	"github.com/go-playground/validator"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)

	filename := m.RequestFile{}

	err := json.NewDecoder(r.Body).Decode(&filename)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	validate := validator.New()
	err = validate.Struct(filename)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}

	db := database.GetConnection()

	file := m.DatabaseFile{}

	query := "SELECT hash, server_name FROM files WHERE user_id = $1 AND name = $2"

	err = db.QueryRow(query, userID, filename.Filename).Scan(&file.Sha, &file.NewFileName)
	if err != nil {
		if err == sql.ErrNoRows {
			lib.ErrorHandler(err, "web")
			w.WriteHeader(400)
			return
		}
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	query = "SELECT COUNT(*) FROM files WHERE hash = $1"

	res, err := db.Query(query, file.Sha)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	var coincidences int8
	res.Scan(&coincidences)


	if coincidences > 1 {

		query = "DELETE FROM files WHERE user_id = $1 AND hash = $2"

		db.QueryRow(query, userID, file.Sha)

		return
	}

	query = "DELETE FROM files WHERE user_id = $1 AND hash = $2"

	db.QueryRow(query, userID, file.Sha)

	err = lib.DeleteFile("/files/" + file.NewFileName)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(202)
}
