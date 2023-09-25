CREATE TABLE IF NOT EXISTS salon_member (
    id UUID NOT NULL PRIMARY KEY,
    salon_id UUID REFERENCES salon(id),
    professional_id UUID REFERENCES professional(id),
    role VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL
)