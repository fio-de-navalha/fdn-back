CREATE TABLE IF NOT EXISTS services (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barbers(id),
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    duration_in_min INTEGER NOT NULL,
    available BOOLEAN  NOT NULL
);
