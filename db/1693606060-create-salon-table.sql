CREATE TABLE IF NOT EXISTS salon (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    owner_id UUID REFERENCES professional(id),
    created_at TIMESTAMP NOT NULL
)