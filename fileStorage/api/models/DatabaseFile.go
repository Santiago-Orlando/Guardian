package models


type DatabaseFile struct {
	ID				int32		`json:"id,omitempty"`
	Filename		string		`json:"name,omitempty"`
	Sha				string		`json:"hash"`
	NewFileName		string		`json:"server_name"`
	UserID			string		`json:"user_id"`
}