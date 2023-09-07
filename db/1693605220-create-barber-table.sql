CREATE TABLE IF NOT EXISTS barber (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
)