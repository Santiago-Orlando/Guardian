CREATE TABLE IF NOT EXISTS database (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE IF NOT EXISTS system (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE IF NOT EXISTS email (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE IF NOT EXISTS web (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE IF NOT EXISTS authentication (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);