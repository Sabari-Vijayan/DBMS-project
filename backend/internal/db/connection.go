package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect()(*pgxpool.Pool, error) {
  //Get database URL from environment
	databaseURL:=os.Getenv("DATABASE_URL")
	if databaseURL == ""{
		databaseURL="postgres://dbms_user:Sabari123@localhost:5432/dbms_project?sslmode=disable"
	}

  //Create connection Pool
	pool,err:=pgxpool.New(context.Background(),databaseURL)
	if err!=nil {
		return nil,err
	}

	//TEST THE connection
	if err:=pool.Ping(context.Background());err!=nil{
		return nil,err
	}

	return pool,nil

}
