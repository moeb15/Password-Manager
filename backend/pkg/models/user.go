package models

type User struct {
	ID        string `json:"id" bson:"_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	MasterKey string `json:"masterkey"`
}
