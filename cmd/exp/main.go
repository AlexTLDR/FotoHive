package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (c PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "rick",
		Password: "picklerick",
		Database: "GalacticFederation",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")

	// Create a table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
			);
			
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
			);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully!")

	// Insert some data
	// name := "Morty"
	// email := "morty@pickle.org"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;`, name, email)
	// //row.Err()
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User inserted successfully, with id:", id, "!")

	// Querying a single record
	id := 1
	row := db.QueryRow(`
		SELECT name, email 
		FROM users 
		WHERE id = $1;`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err == sql.ErrNoRows {
		fmt.Println("No user found with id:", id)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("User:", name, "with email:", email)

	userID := 1
	// for i := 1; i < 5; i++ {
	// 	amount := 100 * i
	// 	description := fmt.Sprintf("Order #%d", i)
	// 	_, err = db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1, $2, $3);`, userID, amount, description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Order inserted successfully!")
	// }

	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	rows, err := db.Query(`
			SELECT id, amount, description
			FROM orders
			WHERE user_id=$1;`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		order.UserID = userID
		err = rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Orders:", orders)
}
