package app

import (
	"SchoolManagementSystem/internal/models"
	"github.com/jackc/pgx"
	"log"
)

func addNewStudentToDatabase(guest models.Guest, db *pgx.Conn) error {
	log.Println(guest.Password)
	insertStmt := "INSERT INTO users (name, role, password) VALUES ($1, $2, $3)"

	_, err := db.Exec(insertStmt, guest.Username, guest.Role, guest.Password)
	return err
}
