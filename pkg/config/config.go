package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	SQL      = "mysql"
	EndPoint = "root:123@tcp(127.0.0.1:8889)/membersql"
	DbName   = "membersql"
)

const (
	MemberSearch = "http://localhost:8881/memberSearch/"
	MemberAll    = "http://localhost:8881/all"
	Health       = "http://localhost:8881/Health"
	DeleteMember = "http://localhost:8881/deleteMember/"
	AddMember    = "http://localhost:8881/addMember/"
	UpdateMember = "http://localhost:8881/updateMember/"
	MemberHippa  = "http://localhost:8881/Hippa/"
)

func SqlDb() *sql.DB {
	db, err := sql.Open(SQL, EndPoint)

	if err != nil {
		panic(err.Error())
	}
	return db
}
