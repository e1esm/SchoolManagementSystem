package models

import "github.com/jackc/pgx"

type SchoolAttendant interface {
	Leave(db *pgx.Conn)
}
