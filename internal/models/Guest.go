package models

type Guest struct {
	Username string
	Password string
	Role     string
}

func NewRoledGuest(username string, password string, role string) Guest {
	return Guest{username, password, role}
}

func NewGuest(username string, password string) Guest {
	return Guest{username, password, ""}
}
