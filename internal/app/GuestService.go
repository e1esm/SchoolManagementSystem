package app

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
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
	row := db.QueryRow("SELECT name from users where name = $1", guest.Username)

	switch err := row.Scan(&name); err {
	case pgx.ErrNoRows:
		doesExist = false
		break
	case nil:
		break
	}

	return doesExist
}

func LogIn(guest Guest, db *pgx.Conn) bool {
	isEntered := true
	sqlQuery := "SELECT 1 from users WHERE password = $1;"
	row := db.QueryRow(sqlQuery, guest.Password)
	var amount int
	switch err := row.Scan(&amount); err {
	case sql.ErrNoRows:
		log.Println("Password is incorrect")
		isEntered = false
	case nil:
		isEntered = true
	}
	return isEntered
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
		//log.Fatal("Couldn't add your credentials in the database.")
		log.Fatal(err.Error())
	}
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func NewRoledGuest(username string, password string, role string) Guest {
	return Guest{username, password, role}
}

func NewGuest(username string, password string) Guest {
	return Guest{username, password, ""}
}

func addNewUserToDatabase(guest Guest, db *pgx.Conn) error {
	log.Println(guest.Password)
	insertStmt := "INSERT INTO users (name, role, password) VALUES ($1, $2, $3)"

	_, err := db.Exec(insertStmt, guest.Username, guest.role, guest.Password)
	return err
}
