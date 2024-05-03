CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(50),
    lastname VARCHAR(50),
    fathers_name VARCHAR(50),
    email VARCHAR(64),
    password VARCHAR(50),
    balance DOUBLE PRECISION DEFAULT 0.0
);
