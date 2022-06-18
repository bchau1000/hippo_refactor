package models

type User struct {
	Uid         string `json:"uid"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	IdToken     string `json:"idToken"`
}
