package database

import (
	"database/sql"
	"fmt"
	log "log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	// Only load .env locally, not on Render
	if os.Getenv("GIN_MODE") != "release" {
		_ = godotenv.Load()
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlSetup := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)

	db, err := sql.Open("postgres", psqlSetup)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Ping ensures the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	DB = db
	log.Println("Successfully connected to database!")
}

func MigrateUsersTable(db *sql.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            password TEXT NOT NULL
        )
    `)
	if err != nil {
		log.Fatal("Failed to migrate users table:", err)
	}
}

func MigrateTopicsTable(db *sql.DB) {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT FALSE,
) `)
	if err != nil {
		log.Fatal("Failed to migrate topics table:", err)
	}

}

func MigrateContentsTable(db *sql.DB) {
	_, err := db.Exec(` CREATE TABLE contents (
    id SERIAL PRIMARY KEY,
    topic_id INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    uri TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_topic
        FOREIGN KEY(topic_id)
            REFERENCES topics(id)
            ON DELETE CASCADE
`)
	if err != nil {
		log.Fatal("Failed to migrate contents table:", err)
	}
}
