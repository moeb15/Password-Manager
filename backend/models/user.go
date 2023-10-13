package models

type User struct {
	Username  string `json:"username"`
	Password  string `json:"-"`
	MasterKey string `json:"-"`
}
