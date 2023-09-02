CREATE TABLE IF NOT EXISTS services (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    barber_id VARCHAR(255) NOT NULL,
    name VARCHAR(45) NOT NULL,
    price INTEGER NOT NULL,
    duration_in_min INTEGER NOT NULL,
    available BOOLEAN  NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_barber FOREIGN KEY (barber_id) REFERENCES barbers (id)
);
