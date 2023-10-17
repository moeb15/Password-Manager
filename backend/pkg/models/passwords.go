package models

type Password struct {
	ID          string `json:"id" bson:"_id"`
	UserID      string `json:"userid" bson:"userid"`
	Application string `json:"application" bson:"application"`
	Password    string `json:"password" bson:"password"`
}
