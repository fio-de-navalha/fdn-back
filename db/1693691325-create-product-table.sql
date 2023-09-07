CREATE TABLE IF NOT EXISTS products (
    id UUID NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barbers(id),
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    available BOOLEAN  NOT NULL
);
