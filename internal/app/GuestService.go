package app

import (
	"crypto/sha512"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
)

type Guest struct {
	Username string
	Password string
	role     string
}

func AlreadyExists(guest Guest, db *pgx.Conn) bool {
	doesExist := true
	var name string
	row := db.QueryRow("SELECT name from users where name = ?", guest.Username)

	switch err := row.Scan(&name); err {
	case sql.ErrNoRows:
		doesExist = false
		break
	default:
		break
	}

	return doesExist
}

func LogIn(guest Guest, db *pgx.Conn) {

}

func SignUp(guest Guest, db *pgx.Conn) {
	var option int
	fmt.Println("Choose your role at school:\n1.Student\n2.Teacher")
	fmt.Fscan(os.Stdin, &option)
	switch option {
	case 1:
		guest.role = "student"
	case 2:
		guest.role = "teacher"
	}
	guest.Password = hashPassword(guest.Password)
	err := addNewUserToDatabase(guest, db)
	if err != nil {
		log.Fatal("Couldn't add your credentials in the database.")
	}
}

func hashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return string(hash.Sum(nil))
}

func NewRoledGuest(username string, password string, role string) Guest {
	return Guest{username, password, role}
}

func NewGuest(username string, password string) Guest {
	return Guest{username, password, ""}
}

func addNewUserToDatabase(guest Guest, db *pgx.Conn) error {
	insertStmt := "INSERT INTO users (name, role, password) VALUES (?, ?, ?)"

	_, err := db.Exec(insertStmt, guest.Username, guest.role, guest.Password)
	return err
}
