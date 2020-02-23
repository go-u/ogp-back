package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server/secrets/cloudsql"
	"strings"
)

var Connection *sql.DB

func InitializeDB(PROJECT_ID string) {
	db_cred := getDBCred(PROJECT_ID)
	Connection = createConnection(db_cred)
}

func getDBCred(PROJECT_ID string) string {
	DB_SOURCE := ""
	if strings.Contains(PROJECT_ID, "prd") {
		DB_SOURCE = secrets_cloudsql.DB_Prd
	} else if strings.Contains(PROJECT_ID, "stg") {
		DB_SOURCE = secrets_cloudsql.DB_Stg
	} else {
		DB_SOURCE = secrets_cloudsql.DB_Local
	}
	if DB_SOURCE == "" {
		log.Fatal("can't find db source")
	}
	return DB_SOURCE
}

func createConnection(DB_SOURCE string) *sql.DB {
	con, err := sql.Open("mysql", DB_SOURCE)
	// defer con.Close() // https://blog.nownabe.com/2017/01/16/570.html#accessing-the-database
	if err != nil {
		log.Fatal("DB Open Error: ", err)
	}
	log.Println("DB initialized")
	return con
}
