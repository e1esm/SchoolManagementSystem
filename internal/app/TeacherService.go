package app

import (
	"SchoolManagementSystem/internal/models"
	"github.com/jackc/pgx"
)

func addNewTeacherToDatabase(guest models.Teacher, password string, subject models.Subject, role string, db *pgx.Conn) error {
	insertStmt := "INSERT INTO users (name, role, password, teacher_of) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(insertStmt, guest.Name, role, password, subject)

	if err != nil {
		return err
	} else {
		return nil
	}
}
