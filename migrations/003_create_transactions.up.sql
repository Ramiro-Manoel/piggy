CREATE TABLE transactions (
    id           TEXT PRIMARY KEY,
    description  TEXT NOT NULL,
	amount       NUMERIC(12, 2) NOT NULL,
	date         TIMESTAMPTZ NOT NULL,
	category_id  TEXT REFERENCES categories(id),
	account_id   TEXT REFERENCES accounts(id) NOT NULL
);