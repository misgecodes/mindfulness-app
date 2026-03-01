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
	// ⚠️ WARNING: This will drop the entire database
	_, err := db.Exec(`DROP DATABASE IF EXISTS your_db_name;`)
	if err != nil {
		log.Fatal("Failed to drop database:", err)
	}
	log.Println("Database dropped successfully!")

	_, err = db.Exec(`CREATE DATABASE your_db_name;`)
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}
	log.Println("Database created successfully!")

	// Reconnect to the newly created DB
	newDB, err := sql.Open("postgres", "postgres://user:password@localhost:5432/your_db_name?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the new database:", err)
	}
	defer newDB.Close()

	// ✅ Create users table
	_, err = newDB.Exec(`
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			is_verified BOOLEAN DEFAULT FALSE
		);
	`)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
	log.Println("Users table migrated successfully!")
}

func MigrateTopicsTable(db *sql.DB) {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT FALSE
) `)
	if err != nil {
		log.Fatal("Failed to migrate topics table:", err)
	}

}

func MigrateContentsTable(db *sql.DB) {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS contents (
    id SERIAL PRIMARY KEY,
    topic_id INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    uri TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_topic
        FOREIGN KEY(topic_id)
            REFERENCES topics(id)
            ON DELETE CASCADE
)`)
	if err != nil {
		log.Fatal("Failed to migrate contents table:", err)
	}
}

func MigrateRefreshTokensTable(db *sql.DB) {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE
);`)
	if err != nil {
		log.Fatal("Failed to migrate refresh_tokens table:", err)
	} else {
		log.Println("Refresh tokens table migrated successfully!")
	}
}

func MigrateOTPTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE otp_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    code TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE
);`)
	if err != nil {
		log.Fatal("Failed to migrate otp_codes table:", err)
	} else {
		log.Println("OTP codes table migrated successfully!")
	}
}
