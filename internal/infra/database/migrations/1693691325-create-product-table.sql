CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    barber_id VARCHAR(255) NOT NULL,
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    available BOOLEAN  NOT NULL,
    CONSTRAINT fk_barber_x_product FOREIGN KEY (barber_id) REFERENCES barbers (id)
);
