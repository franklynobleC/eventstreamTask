package model

type TargetUser struct {
	Id    string
	Email string
	Data  *Data
}

type Data struct {
	Infomessage string
	Email       string
}
