CREATE TABLE IF NOT EXISTS agents (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO agents VALUES (1, 'common name');
