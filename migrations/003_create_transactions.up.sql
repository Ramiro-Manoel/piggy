CREATE TABLE transactions (
    id           TEXT PRIMARY KEY,
    description  TEXT NOT NULL,
	amount       BIGINT NOT NULL DEFAULT 0,
	date         TIMESTAMPTZ NOT NULL,
	category_id  TEXT REFERENCES categories(id),
	account_id   TEXT REFERENCES accounts(id) NOT NULL
);