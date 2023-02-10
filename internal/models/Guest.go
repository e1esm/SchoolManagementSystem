package models

type Guest struct {
	Username string
	Password string
	Role     Role
}

type Role string

const (
	TEACHER Role = "teacher"
	STUDENT Role = "student"
)

func NewRoledGuest(username string, password string, role string) Guest {
	return Guest{username, password, Role(role)}
}

func NewGuest(username string, password string) Guest {
	return Guest{username, password, ""}
}
