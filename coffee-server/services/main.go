package services

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3
var db *sql.DB //# type struct that stores db connection

// Models

type Models struct {
	Coffee Coffee  //* as this "Coffee" type belongs to same pkg and exported with Capital algorithm
	JsonResponse JsonResponse
}

// function that invokes db connection and returns instance of Model struct
func NewModel(dbPool *sql.DB) Models {
	db = dbPool 
	return Models{}
}