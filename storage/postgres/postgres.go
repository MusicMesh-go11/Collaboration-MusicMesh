//package postgres
//
//import (
//	"database/sql"
//	"log"
//	"os"
//
//	_ "github.com/lib/pq"
//)
//
//func Connect() (*sql.DB, error) {
//	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
//	if err != nil {
//		log.Printf("Failed to connect to database: %v", err)
//		return nil, err
//	}
//	return db, nil
//}

package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "BEKJONS"
	dbname   = "collaboration_m"
)

func Conn() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	return db, err
}
