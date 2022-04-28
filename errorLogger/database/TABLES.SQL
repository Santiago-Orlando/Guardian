CREATE TABLE database (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE system (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE email (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE web (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);

CREATE TABLE authentication (
    id          SERIAL          NOT NULL,
    err         TEXT            NOT NULL
);