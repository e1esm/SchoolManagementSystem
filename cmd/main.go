package main

import (
	"SchoolManagementSystem/internal/app"
	"SchoolManagementSystem/internal/models"
	"SchoolManagementSystem/internal/utils"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jackc/pgx"
)

//TODO: https://github.com/johnfercher/maroto - Library for PDF generation.
//TODO: Locale Eng/Rus

var (
	IsTeacher bool
	Name      string
)

func main() {
	//isSigned := true
	cfg := pgx.ConnConfig{Password: os.Getenv("db_password"), User: os.Getenv("db_username"), Database: os.Getenv("db_name"), Port: 5432}
	db, err := pgx.Connect(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
	go setupDB(db)
	username, authorized := authorize(db)
	if authorized {
		menuLogic(username)
	}

}

func menuLogic(username string) {
	isRunning := true

	for isRunning {
		choice := 0
		if IsTeacher {
			fmt.Println("1.Set a mark.\n2.Add a student to a class.\n3.Generate PDF based on marks of X student.\n4.Generate PDF based on marks of all students.")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				isRunning = false
				break
			case 2:
				isRunning = false
				break
			case 3:
				isRunning = false
				break
			case 4:
				isRunning = false
				break
			}
		} else {
			fmt.Println("1.Get all marks.\n2.Get marks of a certain subject.")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				isRunning = false
				break
			case 2:
				isRunning = false
				break
			}
		}
	}
}

func setupDB(db *pgx.Conn) {
	sqlQuery, err := os.ReadFile("./internal/pkg/db/V1.sql")
	if err != nil {

		log.Fatal(err.Error())
	}
	_, err = db.Exec(string(sqlQuery))
	if err != nil {
		log.Fatal(errors.New("can't perform execution of the V1.sql query"))
	}
}

func authorize(db *pgx.Conn) (string, bool) {
	fmt.Println(utils.Green + "!Welcome to School Management System!" + utils.Reset)
	fmt.Println("Log in/ Sign up")
	var password string
	isSigned := false
	inputReader := bufio.NewReader(os.Stdin)

	for !isSigned {
		fmt.Println("Enter username: ")
		Name, _ = inputReader.ReadString('\n')

		fmt.Println("Enter password: ")
		password, _ = inputReader.ReadString('\n')

		guest := models.NewGuest(Name, password)

		exist := app.AlreadyExists(guest, db)
		if exist {

			var role string
			role, isSigned = app.LogIn(guest, db)
			if role == "teacher" {
				IsTeacher = true
			} else {
				IsTeacher = false
			}

		} else {
			role := app.SignUp(guest, db)
			if role == "teacher" {
				IsTeacher = true
			} else {
				IsTeacher = false
			}
			isSigned = true
		}
	}
	fmt.Println(utils.Yellow + "Successfully logged in." + utils.Reset)
	time.Sleep(time.Second)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return Name, isSigned
}
