package main

import (
	"SchoolManagementSystem/internal/app"
	"SchoolManagementSystem/internal/utils"
	"bufio"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
)

//TODO: https://github.com/johnfercher/maroto - Library for PDF generation.
//TODO: Locale Eng/Rus

func main() {
	//isSigned := true
	isRunning := true
	cfg := pgx.ConnConfig{Password: os.Getenv("db_password"), User: os.Getenv("db_username"), Database: os.Getenv("db_name"), Port: 5432}
	db, err := pgx.Connect(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
	go setupDB(db)
	for isRunning {
		fmt.Println(utils.Green + "!Welcome to School Management System!" + utils.Reset)
		fmt.Println("1.Log in/ Sign up")
		authorize(db)
	}

}

func setupDB(db *pgx.Conn) {
	sqlQuery, err := os.ReadFile("./internal/pkg/db/V1.sql")
	if err != nil {
		//log.Fatal(errors.New("can't find required file"))
		log.Fatal(err.Error())
	}
	_, err = db.Exec(string(sqlQuery))
	if err != nil {
		log.Fatal(errors.New("can't perform execution of the V1.sql query"))
	}
}

func authorize(db *pgx.Conn) bool {
	var username string
	var password string
	isSigned := false
	inputReader := bufio.NewReader(os.Stdin)

	for !isSigned {
		fmt.Println("Enter username: ")
		username, _ = inputReader.ReadString('\n')
		fmt.Println("Enter password: ")
		password, _ = inputReader.ReadString('\n')

		guest := app.NewGuest(username, password)

		exist := app.AlreadyExists(guest, db)
		if exist {

			isSigned = app.LogIn(guest, db)

		} else {
			app.SignUp(guest, db)
			isSigned = true
		}
	}
	log.Println("Successfully logged in.")
	return isSigned
}
