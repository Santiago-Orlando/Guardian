package models


type RequestFile struct {
	Filename		string		`json:"filename" validate:"required"`
}