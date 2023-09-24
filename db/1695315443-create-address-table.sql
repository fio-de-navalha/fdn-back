CREATE TABLE IF NOT EXISTS address (
    id UUID NOT NULL PRIMARY KEY,
    salon_id UUID REFERENCES salon(id),
    address VARCHAR NOT NULL
);
