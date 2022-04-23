package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	ID				primitive.ObjectID 	`bson:"_id,omitempty" json:"id,omitempty"`
	UserName		string 				`json:"userName"`
	Email			string				`json:"email"`
	Password		string				`json:"password"`
	RecoveryToken	string				`json:"recoveryToken,omitempty"`

}

type UserLogin struct {
	Email			string				`json:"email"`
	Password		string				`json:"password"`
}

type Email struct {
	Email			string				`bson:"email" json:"email"`
}

type PasswordRecovery struct {
	Token			string 				`json:"token"`
	NewPassword		string				`json:"newPassword"`
}