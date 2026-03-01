package database

const user = `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE
        );
  `

const topics = ` CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
); `

const content = ` CREATE TABLE contents (
    id SERIAL PRIMARY KEY,
    topic_id INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    uri TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_topic
        FOREIGN KEY(topic_id)
            REFERENCES topics(id)
            ON DELETE CASCADE
`
