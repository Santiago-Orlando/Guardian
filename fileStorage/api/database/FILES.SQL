
CREATE TABLE IF NOT EXISTS files (
    id          serial       NOT NULL,
    name        VARCHAR(150) NOT NULL,
    hash        VARCHAR(256) NOT NULL,
    server_name VARCHAR(150) NOT NULL,
    user_id     VARCHAR(150) NOT NULL,
    created_at  TIMESTAMP DEFAULT now(),
    updated_at  TIMESTAMP,
    CONSTRAINT pk_files PRIMARY KEY(id)
);