package app

import (
	"SchoolManagementSystem/internal/models"
	"SchoolManagementSystem/internal/utils"
	"crypto/sha256"
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

func LogIn(guest models.Guest, db *pgx.Conn) (models.SchoolAttendant, bool) {
	isEntered := true
	sqlQuery := "SELECT id, name, role, teacher_of, is_left from users WHERE password = $1;"

	hash := sha256.New()
	hash.Write([]byte(guest.Password))
	guest.Password = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	row := db.QueryRow(sqlQuery, guest.Password)

	var id int
	var name, role, teacher_of string
	var is_left bool
	switch err := row.Scan(&id, &name, &role, &teacher_of, &is_left); err {
	case pgx.ErrNoRows:
		log.Println(utils.Red + "Password is incorrect" + utils.Reset)
		isEntered = false
		return nil, false
	case nil:
		isEntered = true
	}
	var schoolAttendant models.SchoolAttendant
	if is_left {
		log.Fatal("This user has quit our school.")
	} else {
		if role == string(models.TEACHER) {
			schoolAttendant = models.TeacherFullyGenerator(models.Guest{Username: name, Role: models.TEACHER}, models.Subject(teacher_of), id)
		} else {
			schoolAttendant = models.StudentGenerator(id, name)
		}
	}
	return schoolAttendant, isEntered
}

func SignUp(guest models.Guest, db *pgx.Conn) (models.SchoolAttendant, bool) {

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
		return models.Student{Name: guest.Username}, true
	case 2:
		guest.Role = "teacher"
		subject := chooseSubject()
		teacher := models.TeacherGenerator(guest, subject)
		err := addNewTeacherToDatabase(teacher, guest.Password, subject, string(guest.Role), db)
		if err != nil {
			log.Fatal("Couldn't add new teacher to the database.")
		}
		return models.TeacherGenerator(guest, subject), true
	}
	return nil, false
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
