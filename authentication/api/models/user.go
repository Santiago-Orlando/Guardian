package models

type User struct {

	ID				string 				`bson:"_id,omitempty" json:"id,omitempty"`
	Username		string 				`json:"username"`
	Email			string				`json:"email"`
	Password		string				`json:"password"`
	RecoveryToken	string				`bson:"recoveryToken" json:"recoveryToken"`

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