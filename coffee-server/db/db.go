package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	 
)

//  types --> these are just type --> we need actual instance with data --> so those are instance
//  and if some func uses it --> becomes method to be accesible via that initiated instance
//  cause after all its instance actually which implemented that method to use that type data
type DB struct {
	DB *sql.DB
}

var DBConnection = &DB{} // empty instance of DB type struct
const (	
	maxOpenDbConnections = 10
	maxIdleDbConnections = 5
	maxDbLifetime = maxIdleDbConnections * time.Minute
)

// @Connection to database
func Connect2Postgress(dbConnString string) (*DB,error) {
	// opening sql connection
	db,err := sql.Open("pgx",dbConnString)
	if err != nil {
		return nil,err
	}

	// configuring database properties
	db.SetMaxOpenConns(maxOpenDbConnections)
	db.SetMaxIdleConns(maxIdleDbConnections)
	db.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(db)
	if err != nil {
		return nil,err
	}
	//* assigning db to fill in instance --> so we get that data from this func return
	DBConnection.DB = db  //stores db
	return DBConnection,nil
}

func testDB(db *sql.DB)error {
	err := db.Ping()
	if err != nil {
		return err
	}
	
	fmt.Println("** Pinged database successfully ✅✅ **")
	return nil
}