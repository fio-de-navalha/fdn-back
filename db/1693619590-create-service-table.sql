CREATE TABLE IF NOT EXISTS service (
    id UUID NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barber(id),
    name VARCHAR(45) NOT NULL,
    description VARCHAR,
    price INTEGER NOT NULL,
    duration_in_min INTEGER NOT NULL,
    available BOOLEAN NOT NULL,
    image_id VARCHAR,
    image_url VARCHAR
);
