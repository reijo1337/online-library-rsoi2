package main

type User struct {
	ID       int32
	Login    string `json:"login"`
	Password string `password:"password"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}
