package data

import "database/sql"

type DBClient interface {

}

type DBClientImpl struct {
	DBURL string
	DB    *sql.DB
}

func NewDBClient(url string) (client DBClient, err error) {

	db, err := sql.Open("mysql", url)

	client = &DBClientImpl{
		DBURL: url,
		DB:    db,
	}

	return client, err
}