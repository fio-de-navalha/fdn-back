CREATE TABLE IF NOT EXISTS security_question (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(255) NOT NULL
);
