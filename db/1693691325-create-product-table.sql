CREATE TABLE IF NOT EXISTS product (
    id UUID NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barber(id),
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    available BOOLEAN  NOT NULL
);
