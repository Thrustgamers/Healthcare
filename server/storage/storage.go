package storage

type UserData struct {
	UserID int `json:"UserID,string"`
	Token  int `json:"Token,string"`
}

var SessionManager = make(map[int]UserData)
