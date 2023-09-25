CREATE TABLE IF NOT EXISTS contact (
    id UUID NOT NULL PRIMARY KEY,
    salon_id UUID REFERENCES salon(id),
    contact VARCHAR NOT NULL
);
