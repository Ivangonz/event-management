CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100)
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    description TEXT,
    date TIMESTAMP,
    user_id INTEGER REFERENCES users(id)
);
