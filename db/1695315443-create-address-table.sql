CREATE TABLE IF NOT EXISTS address (
    id UUID NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barber(id),
    address VARCHAR NOT NULL
);
