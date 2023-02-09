package app

import (
	"SchoolManagementSystem/internal/models"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
)

func AlreadyExists(guest models.Guest, db *pgx.Conn) bool {
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

func LogIn(guest models.Guest, db *pgx.Conn) (string, bool) {
	isEntered := true
	sqlQuery := "SELECT role from users WHERE password = $1;"
	row := db.QueryRow(sqlQuery, guest.Password)
	var role string
	switch err := row.Scan(&role); err {
	case sql.ErrNoRows:
		log.Println("Password is incorrect")
		isEntered = false
	case nil:
		isEntered = true
	}
	if role == "teacher" {
		return "teacher", isEntered
	} else {
		return "student", isEntered
	}
}

func SignUp(guest models.Guest, db *pgx.Conn) string {

	var option int
	fmt.Println("Choose your role at school:\n1.Student\n2.Teacher")
	fmt.Fscan(os.Stdin, &option)
	guest.Password = hashPassword(guest.Password)
	switch option {
	case 1:
		guest.Role = "student"
		err := addNewStudentToDatabase(guest, db)
		if err != nil {
			log.Fatal("Couldn't add student's credentials in the database.")
		}
		return guest.Role
	case 2:
		guest.Role = "teacher"
		subject := chooseSubject()
		teacher := models.TeacherGenerator(guest, subject)
		err := addNewTeacherToDatabase(teacher, guest.Password, subject, guest.Role, db)
		if err != nil {
			log.Fatal("Couldn't add new teacher to the database.")
		}
		return guest.Role
	}
	return ""
}

func chooseSubject() models.Subject {
	fmt.Println("1.Maths\n2.Foreign language\n3.Informatics\n4.Native language\n5.Drawing\n6.Music\n7.Literature\n8.Law science")
	option := 0
	for !(option <= 8 && option >= 1) {
		fmt.Scan(&option)
	}
	switch option {
	case 1:
		return models.Maths
	case 2:
		return models.ForeignLanguage
	case 3:
		return models.Informatics
	case 4:
		return models.NativeLanguage
	case 5:
		return models.Drawing
	case 6:
		return models.Music
	case 7:
		return models.Literature
	case 8:
		return models.LawScience
	default:
		return ""
	}
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
