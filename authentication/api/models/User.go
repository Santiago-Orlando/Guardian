package models


type User struct {

	ID				string 				`bson:"_id,omitempty" json:"id,omitempty"`
	Username		string 				`json:"username" validate:"required"`
	Email			string				`json:"email" validate:"required,email"`
	Password		string				`json:"password" validate:"required,gte=8,lte=60"`
	RecoveryToken	string				`bson:"recoveryToken" json:"recoveryToken"`

}

type UserLogin struct {
	Email			string				`json:"email" validate:"required,email"`
	Password		string				`json:"password" validate:"required,gte=8,lte=33"`
}

type Email struct {
	Email			string				`bson:"email" json:"email" validate:"required,email"`
}

type PasswordRecovery struct {
	Token			string 				`json:"token" validate:"required"`
	NewPassword		string				`json:"newPassword" validate:"required,gte=8,lte=33"`
}