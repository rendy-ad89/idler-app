package util

func ValidateUserRequest(username, password string) bool {
	if len(username) < 5 || len(username) > 20 {
		return false
	}
	if len(password) < 5 || len(password) > 20 {
		return false
	}
	return true
}
