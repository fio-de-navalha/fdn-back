CREATE TABLE IF NOT EXISTS contact (
    id UUID NOT NULL PRIMARY KEY,
    barber_id UUID REFERENCES barber(id),
    contact VARCHAR NOT NULL
);
