package entity

type UserCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	cred UserCred
}
