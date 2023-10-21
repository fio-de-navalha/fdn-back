CREATE TABLE IF NOT EXISTS product (
    id UUID NOT NULL PRIMARY KEY,
    salon_id UUID REFERENCES salon(id),
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    available BOOLEAN  NOT NULL,
    image_id VARCHAR,
    image_url VARCHAR
);
