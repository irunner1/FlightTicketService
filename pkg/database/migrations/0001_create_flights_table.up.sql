CREATE TABLE IF NOT EXISTS flights (
    id SERIAL PRIMARY KEY,
    airline VARCHAR(30) NOT NULL,
    origin VARCHAR(30) NOT NULL,
    destination VARCHAR(30) NOT NULL,
    departure TIMESTAMP NOT NULL,
    arrival TIMESTAMP NOT NULL,
    price REAL NOT NULL
);