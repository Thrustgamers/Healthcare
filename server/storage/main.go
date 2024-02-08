package storage

type UserData struct {
	UserID int  `json:"UserID,string"`
	Token  int  `json:"Token,string"`
	Admin  bool `json:"Admin,omitempty" default:"false"`
}

var SessionManager = make(map[int]UserData)
