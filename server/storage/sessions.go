package storage

import "github.com/google/uuid"

type UserData struct {
	UserID     int       `json:"UserID,string"`
	Token      uuid.UUID `json:"Token,string"`
	EmployeeId int       `json:"EmployeeId,string"`
	Admin      bool      `json:"Admin,omitempty" default:"false"`
}

var SessionManager = make(map[int]UserData)

func DoesSessionExist(EmployeeId int) bool {
	for i := range SessionManager {
		if SessionManager[i].EmployeeId == EmployeeId {
			return true
		}
	}
	return false
}
