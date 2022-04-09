package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jackc/pgx/v4"

	"github.com/georgysavva/scany/pgxscan"
)

func PostgresSimpleQuery() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:example@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var position string
	var id int64
	err = conn.QueryRow(context.Background(), "select * from users where id=$1", 2).Scan(&id, &name, &position)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, position)
}

func PostgresScanToObject() {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, "postgres://postgres:example@localhost:5432/postgres")

	type User struct {
		Name string
		Rank string
		Id   int
	}

	defer db.Close()

	var users []User
	pgxscan.Select(ctx, db, &users, "select * from users") //lib to map to slice
	fmt.Println(users)

	var user User
	pgxscan.Get(ctx, db, &user, "select * from users where id=$1", 2) //get & map 1 job
	fmt.Println(user)
}

func PostgresJoinQuery() {
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, "postgres://postgres:example@localhost:5432/postgres")

	query := `SELECT users.name, users.rank, c.brand, c.colour, c.license_plate FROM users JOIN cars c on users.id = c.user_id WHERE rank = $1`

	type UserCar struct {
		Name         string
		Rank         string
		Brand        string
		Colour       string
		LicensePlate string
	} //для join запроса нам нужна плоская структура, хранящая в себе данные

	var carsBounds []UserCar

	pgxscan.Select(ctx, db, &carsBounds, query, "CEO")

	fmt.Println(carsBounds)
}

func PostgresAddCar(userId int, colour string, brand string, licensePlate string) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:example@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(ctx, "INSERT INTO cars(user_id, colour, brand, license_plate) VALUES ($1,$2,$3,$4)", userId, colour, brand, licensePlate) // для записи в базу нам надо смаппить конкретные поля в значения
	fmt.Println(err)
}

func main() {
	// PostgresSimpleQuery()
	// PostgresScanToObject()
	// PostgresJoinQuery()
	PostgresAddCar(9, "orange", "Lexus", "OEM")
}
