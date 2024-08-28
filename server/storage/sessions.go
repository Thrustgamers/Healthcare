package storage

import "github.com/google/uuid"

type UserSession struct {
	UserID     uint      `json:"UserID,string"`
	Token      uuid.UUID `json:"Token,string"`
	EmployeeId int       `json:"EmployeeId,string"`
	Admin      bool      `json:"Admin,omitempty" default:"false"`
}

var SessionManager = make(map[uint]UserSession)

func DoesSessionExist(EmployeeId int) bool {
	for i := range SessionManager {
		if SessionManager[i].EmployeeId == EmployeeId {
			return true
		}
	}
	return false
}
