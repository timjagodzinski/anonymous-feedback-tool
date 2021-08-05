CREATE TABLE fruit (
    id SERIAL PRIMARY KEY,
    name text NOT NULL UNIQUE,
    amount integer NOT NULL DEFAULT 0
    price integer NOT NULL
);
