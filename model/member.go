package model

// 會員資料結構
type Member struct {
	Name     string `bson:"name"`
	Account  string `bson:"account"`
	Password string `bson:"password"`
}