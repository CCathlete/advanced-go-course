package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Make sure to db.Close()!
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString())
	if err != nil {
		return nil, fmt.Errorf("error when opening db: %w", err)
	}

	return db, nil
}

func connectionString() string {
	cfg := PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User,
		cfg.Password, cfg.DBName, cfg.SSLMode,
	)
}

// This will be the last thing executed (first defer).
// It gets the last error value of the function and the transaction
// and commits or rolls back according to the error.
func CommitOrRollback(transaction *sql.Tx, err error) {
	switch err {
	case nil:
		// Committing the transaction to the DB.
		// If we have multiple statements in the same transaction
		// we need to commit after all are executed successfully!
		err := transaction.Commit()
		if err != nil {
			log.Println("There was a problem with committing the transaction.")
		}
	default:
		err := transaction.Rollback()
		if err != nil {
			log.Println(
				"There was a problem with rolling the transaction back.")
		}
	}
}

func Ex1(db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	var email string
	rows, err := db.Query(`
	select email from users
	;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email)
		fmt.Println(email)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Ex2(db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	var id int
	var email string

	// Starting a transaction.
	transaction, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	defer CommitOrRollback(transaction, err)

	// Creating the first statement to execute.
	statement, err := transaction.Prepare(`
	select id, email from users
	;`)
	if err != nil {
		log.Println(err)
		return
	}
	defer statement.Close()

	// Executing the statement.
	rows, err := statement.Query()
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	// Processing the result.
	for rows.Next() {
		err := rows.Scan(&id, &email)
		fmt.Println(id, email)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	var wg *sync.WaitGroup
	wg.Add(2)
	db, err := ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	go Ex1(db, wg)
	go Ex2(db, wg)

	wg.Wait()
}
