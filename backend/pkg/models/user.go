package models

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	MasterKey string     `json:"-"`
	SavedPwds []Password `json:"-"`
}
