package models

type Password struct {
	ID          string `json:"id" bson:"_id"`
	UserID      string `json:"userid"`
	Application string `json:"application"`
	Password    string `json:"password"`
}
