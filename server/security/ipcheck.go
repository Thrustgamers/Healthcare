package security

type ipCheck struct {
	Attempts int `json:"Attemps"`
}

var IpManager = make(map[string]ipCheck)

func addIpToCheck(ip string) {

	exists := CheckIp(ip)
	if exists {
		incrementAttempts(ip)
	}

	IpManager[ip] = ipCheck{
		Attempts: 1,
	}
}

func incrementAttempts(ip string) bool {

	ipData, ok := IpManager[ip]

	if !ok {
		return false
	}

	ipData.Attempts++

	return true
}

func CheckIp(ip string) bool {

	_, ok := IpManager[ip]

	if !ok {
		addIpToCheck(ip)
	}

	if ok {
		incrementAttempts(ip)
	}

	return ok
}

func RemoveIpFromCheck(ip string) bool {

	_, ok := IpManager[ip]

	if ok {
		delete(IpManager, ip)
	}

	return ok

}
